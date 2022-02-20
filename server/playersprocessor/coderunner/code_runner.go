package coderunner

import (
	"log"
	"sync"

	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/server/command"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type CodeRunner struct {
	commandCh                chan command.Command
	posCh                    chan wpcomponent.Position
	cpuResourcesPerWorldTick int
	cpuDone                  chan bool
	cpuStart                 chan bool

	mu           sync.Mutex
	cpuResources int
}

type CpuTickFn func(int)
type CodeFn func(chan command.Command, chan wpcomponent.Position, CpuTickFn)

func NewCodeRunner(commandCh chan command.Command, posCh chan wpcomponent.Position, cpuResourcesPerWorldTick int) *CodeRunner {
	return &CodeRunner{
		commandCh:                commandCh,
		posCh:                    posCh,
		cpuResourcesPerWorldTick: cpuResourcesPerWorldTick,
		cpuDone:                  make(chan bool),
		cpuStart:                 make(chan bool),
	}
}

func (cr *CodeRunner) Run(codeFn CodeFn) {
	p := wpcomponent.Position{Pos: fgeom.Point{X: 0, Y: 10}}
	go codeFn(cr.commandCh, cr.posCh, cr.cpuTick)

	posReceived := true
	var cmd command.Command
	for t := 0; t < 100; t++ {
		log.Printf("world tick [%d] started", t)
		if posReceived {
			log.Printf("sending pos")
			cr.posCh <- p
			posReceived = false
		}
		cr.cpuRun(cr.cpuResourcesPerWorldTick)
		log.Printf("waiting cmd")
		select {
		case cmd = <-cr.commandCh:
			log.Printf("cmd received: %v", cmd)
			posReceived = true
		default:
		}
		p = wpcomponent.Position{Pos: p.Pos.Add(fgeom.Point{X: cmd.Move})}
		log.Printf("world tick [%d] ended, pos: %v\n", t, p.Pos)
	}
}

func (cr *CodeRunner) cpuRun(cpuTicks int) {
	log.Printf("sending cpu start")
	cr.cpuStart <- true
	log.Printf("cpu start sent")

	log.Printf("waiting cpu done")
	<-cr.cpuDone

	cr.mu.Lock()
	cr.cpuResources += cpuTicks
	c := cr.cpuResources
	cr.mu.Unlock()

	log.Printf("cpu tick added: %d", c)
}

func (cr *CodeRunner) cpuTick(complexity int) {
	select {
	case <-cr.cpuStart:
		log.Printf("code: cpu start received")
	default:
	}

	cr.mu.Lock()
	if complexity == 0 {
		cr.cpuResources = 0
	} else {
		cr.cpuResources -= complexity
	}
	c := cr.cpuResources
	cr.mu.Unlock()

	log.Printf("code: cpu tick %d, reesources: %d", complexity, c)
	for c <= 0 {
		log.Printf("code: cpu exhausted: %d", c)
		log.Printf("code: sending cpuDone")
		cr.cpuDone <- true
		cr.mu.Lock()
		c = cr.cpuResources
		cr.mu.Unlock()
	}
}
