package debugger

import (
	"html/template"
	"io/fs"
	"os"
	"sync"
	"time"
)

type column struct {
	Name string
}

type row struct {
	TimeId  int64
	RowData []string
}

type data struct {
	Title   string
	Columns []column
	Rows    []row
}

type HtmlReport struct {
	mu                 sync.Mutex
	reportFilePath     string
	templatesFs        fs.ReadFileFS
	delay              time.Duration
	threadIndexes      map[*Thread]int
	data               *data
	makeReportSchedule *time.Timer
}

func NewHtmlReport(reportFilePath string, templatesFs fs.ReadFileFS, delay time.Duration) *HtmlReport {
	h := &HtmlReport{
		reportFilePath: reportFilePath,
		templatesFs:    templatesFs,
		delay:          delay,
		threadIndexes:  map[*Thread]int{},
		data: &data{
			Title:   "Report",
			Columns: []column{},
			Rows:    []row{},
		},
	}
	h.makeReportSchedule = time.AfterFunc(delay, h.MakeReport)
	return h
}

func (h *HtmlReport) AddThread(t *Thread) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.data.Columns = append(h.data.Columns, column{Name: t.name})
	h.threadIndexes[t] = len(h.data.Columns) - 1
}

func (h *HtmlReport) AddLog(te TimeEntry) {
	h.mu.Lock()
	defer h.mu.Unlock()

	var s []string
	ti := h.threadIndexes[te.thread]
	for i := 0; i < ti; i++ {
		s = append(s, "")
	}
	s = append(s, te.logStr)
	if ti < len(h.data.Columns)-1 {
		for i := ti; i < len(h.data.Columns)-1; i++ {
			s = append(s, "")
		}
	}
	h.data.Rows = append(h.data.Rows, row{
		TimeId:  te.tId,
		RowData: s,
	})
}

func (h *HtmlReport) resetTimer() {
	h.makeReportSchedule.Reset(h.delay)
}

func (h *HtmlReport) MakeReport() {
	h.mu.Lock()
	defer h.mu.Unlock()

	t := template.Must(template.ParseFS(h.templatesFs, "*.html"))
	f, err := os.OpenFile(h.reportFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()

	if err = t.Execute(f, h.data); err != nil {
		panic(err)
	}
	h.resetTimer()
}

func (h *HtmlReport) Finish() {
	h.MakeReport()
	h.makeReportSchedule.Stop()
}
