package debugger

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const traceSplitter = " > "

type lazyReporter interface {
	AddThread(thread *Thread)
	AddLog(log TimeEntry)
	Finish()
}

func NewDebuggerWithReportFinish(enabled bool, r lazyReporter) (*Debugger, func()) {
	t := time.Now()
	return &Debugger{
		enabled:   enabled,
		startTime: t,
		r:         r,
	}, r.Finish
}

func (d *Debugger) CreateNested(name string) *Nested {
	t := &Thread{name: name}
	d.AddThread(t)
	return &Nested{
		name:   name,
		root:   d,
		thread: t,
	}
}

type Thread struct {
	name string
}

type TimeEntry struct {
	tId    int64
	thread *Thread
	method string
	logStr string
}

type Debugger struct {
	mu        sync.Mutex
	enabled   bool
	startTime time.Time
	r         lazyReporter
}

func (d *Debugger) AddThread(t *Thread) {
	d.r.AddThread(t)
}

func (d *Debugger) Log(t *Thread, method, logStr string) {
	te := TimeEntry{
		tId:    time.Now().Sub(d.startTime).Microseconds(),
		thread: t,
		method: method,
		logStr: logStr,
	}
	d.r.AddLog(te)
}

type Nested struct {
	name   string
	root   *Debugger
	prev   *Nested
	prefix []string
	thread *Thread
}

func (n *Nested) CreateNested(name string) *Nested {
	return &Nested{
		name:   name,
		prefix: append(n.prefix, name),
		prev:   n,
		root:   n.root,
		thread: n.thread,
	}
}

func (n *Nested) CreateNestedConcurrent(name string) *Nested {
	t := &Thread{name: n.thread.name + traceSplitter + name}
	n.root.AddThread(t)

	return &Nested{
		name:   name,
		prev:   n,
		root:   n.root,
		thread: t,
	}
}

func (n *Nested) LogF(method string, str string, args ...interface{}) {
	if !n.root.enabled {
		return
	}
	var fullMethod = ""
	if len(n.prefix) != 0 {
		fullMethod = strings.Join(append(n.prefix, method), traceSplitter)
	} else {
		fullMethod = method
	}
	logStr := fmt.Sprintf("%s: %s", fullMethod, fmt.Sprintf(str, args...))

	n.root.Log(n.thread, method, logStr)
}
