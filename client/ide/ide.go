package ide

import (
	"github.com/justclimber/fda/client/ide/ast"
	"github.com/justclimber/fda/client/ide/program"
)

type IDE struct {
	program         *program.Program
	tabs            []*Tab
	currentTabIndex int
	// version control
}

func NewIDE(program *program.Program, tabs []*Tab, currentTabIndex int) *IDE {
	return &IDE{program: program, tabs: tabs, currentTabIndex: currentTabIndex}
}

func (id *IDE) Render(r ast.Renderer) {
	// current tab . render
	r.DrawHeaderTab()
	offset := 0.
	for i, tab := range id.tabs {
		if i == id.currentTabIndex {
			offset = r.DrawActiveTab(tab.name, offset)
		} else {
			offset = r.DrawInactiveTab(tab.name, offset)
		}
	}
	id.tabs[id.currentTabIndex].Draw(r)

	// render general interface
	// render help
}
