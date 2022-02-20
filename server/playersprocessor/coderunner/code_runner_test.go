package coderunner

import (
	"log"
	"testing"

	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/server/command"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

func TestCodeRunner_Run(t *testing.T) {
	commandCh := make(chan command.Command, 1)
	posCh := make(chan wpcomponent.Position, 1)

	codeFn := func(commandCh chan command.Command, posCh chan wpcomponent.Position, cpuTickFn CpuTickFn) {
		targetPos := fgeom.Point{X: 5, Y: 10}
		cpuTickFn(2)
		var move float64
		cpuTickFn(1)
		for {
			log.Printf("code: waiting pos")
			select {
			case pos := <-posCh:
				log.Printf("code: pos received: %v", pos.Pos)
				cpuTickFn(2)
				if pos.Pos.EqualApprox(targetPos, 0.1) {
					log.Printf("exit!!!!!")
					cpuTickFn(5)
					return
				}
				cpuTickFn(5)

				d := targetPos.Sub(pos.Pos)
				cpuTickFn(5)

				if d.X > 0 {
					move = 1
					cpuTickFn(3)
				} else {
					move = -1
					cpuTickFn(3)
				}
				cpuTickFn(3)
				c := command.Command{Move: move}
				log.Printf("code: cmd to send: %v", c)
				commandCh <- c
				log.Printf("code: cmd sent")
				cpuTickFn(0)
			}
		}
	}
	cpuResourcesPerWorldTick := 50

	cr := NewCodeRunner(commandCh, posCh, cpuResourcesPerWorldTick)
	cr.Run(codeFn)
}
