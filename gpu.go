package gcn3

import (
	"fmt"

	"gitlab.com/yaotsu/core"
)

// A Gpu is the unit that one kernel can run on.
//
// A Gpu is a Yaotsu component and it defines the port "ToDriver". Driver is
// a piece of software that conceptually runs in the Cpu. Therefore, all the
// CPU-GPU communication happens on the connection connecting the "ToDriver"
// port.
type Gpu struct {
	*core.BasicComponent

	Driver           core.Component // The DriverComponent
	CommandProcessor core.Component // The CommandProcessor
}

// NewGpu returns a newly created GPU
func NewGpu(name string) *Gpu {
	g := new(Gpu)
	g.BasicComponent = core.NewBasicComponent(name)
	g.AddPort("ToDriver")
	g.AddPort("ToCommandProcessor")
	return g
}

// Handle defines how a GPU handles core.
//
// A GPU should not handle any event by itself.
func (g *Gpu) Handle(e core.Event) error {
	return nil
}

// Receive processes incoming request to the GPU.
//
// The GPU itself does not respond to requests, but it always forward to the
// CommandProcessor.
func (g *Gpu) Receive(req core.Request) *core.Error {
	if req.Source() == g.CommandProcessor { // From the CommandProcessor
		req.SetSource(g)
		req.SetDestination(g.Driver)
		g.GetConnection("ToDriver").Send(req)
		return nil
	} else if req.Source() == g.Driver { // From the Driver
		req.SetSource(g)
		req.SetDestination(g.CommandProcessor)
		g.GetConnection("ToCommandProcessor").Send(req)
		return nil
	}

	return core.NewError(
		fmt.Sprintf("Unrecognized source %s", req.Source().Name()), false, 0)
}
