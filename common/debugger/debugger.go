package debugger

import (
	"fmt"
	"log"
)

const traceSplitter = " -> "

func NewDebugger(name string, enabled bool) *Debugger {
	return &Debugger{
		name:    name,
		enabled: enabled,
	}
}

type Debugger struct {
	name    string
	prev    *Debugger
	enabled bool
}

func (d *Debugger) CreateNested(name string) *Debugger {
	return &Debugger{
		name:    name,
		prev:    d,
		enabled: d.enabled,
	}
}

func (d *Debugger) Printf(method string, str string, args ...interface{}) {
	if !d.enabled {
		return
	}
	trace := d.gatherTrace()
	log.Printf(fmt.Sprintf("%s: %s. %s", trace, method, str), args...)
}

func (d *Debugger) gatherTrace() string {
	if d.prev != nil {
		return d.prev.gatherTrace() + traceSplitter + d.name
	}
	return d.name
}
