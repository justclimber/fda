package widget

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestLabeledCheckbox_SetState_User(t *testing.T) {
	l := newLabeledCheckbox(t)
	leftMouseButtonClick(labeledCheckboxLabel(l), t)

	assert.Equal(t, CheckboxChecked, l.Checkbox().State())
}

func newLabeledCheckbox(t *testing.T, opts ...LabeledCheckboxOpt) *LabeledCheckbox {
	t.Helper()

	l := NewLabeledCheckbox(append(opts, []LabeledCheckboxOpt{
		LabeledCheckboxOpts.CheckboxOpts(
			CheckboxOpts.ButtonOpts(ButtonOpts.Image(&ButtonImage{
				Idle: newNineSliceEmpty(t),
			})),
			CheckboxOpts.Image(&CheckboxGraphicImage{
				Unchecked: &ButtonImageImage{
					Idle: newImageEmpty(t),
				},
				Checked: &ButtonImageImage{
					Idle: newImageEmpty(t),
				},
			}),
		),
		LabeledCheckboxOpts.LabelOpts(LabelOpts.Text("", loadFont(t), &LabelColor{
			Idle: color.White,
		})),
	}...)...)
	event.ExecuteDeferred()
	render(l, t)
	return l
}

func labeledCheckboxLabel(l *LabeledCheckbox) *Label {
	return l.label
}
