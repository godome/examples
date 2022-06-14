package helloworld

import (
	"github.com/godome/godome/pkg/module"
	fiberHandler "github.com/godome/plugins/pkg/fiber"
	"github.com/gofiber/fiber/v2"
)

func newHelloworldHandler(m module.Module) fiberHandler.FiberHandler {
	h := fiberHandler.NewFiberHandler(m)
	h.AddRoute(func(a *fiber.App) {
		service := m.GetProvider(ServiceType).(HelloworldService)

		a.Get("/:name", func(c *fiber.Ctx) error {
			c.Send([]byte(service.SayHello(c.Params("name"))))
			return nil
		})
	})
	return h
}
