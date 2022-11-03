package input

type ControlArrow int

const (
	ControlArrowNone ControlArrow = iota
	ControlArrowDown
	ControlArrowLeft
	ControlArrowRight
	ControlArrowUp
)

type Input interface {
	WhichControlArrowsPressed() ControlArrow
}
