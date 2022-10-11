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
	id.tabs[id.currentTabIndex].Draw(r)

	// render general interface
	// render tab bar
	// render help
}
