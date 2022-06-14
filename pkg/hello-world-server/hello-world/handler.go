package helloworld

import (
	"github.com/godome/godome/pkg/component/module"
	fiberHandler "github.com/godome/plugins/pkg/fiber"
	"github.com/gofiber/fiber/v2"
)

func newHelloworldHandler(m module.Module) fiberHandler.FiberHandler {
	return fiberHandler.
		NewFiberHandler().
		AddRoute(func(a *fiber.App) {
			service := m.GetProvider(HelloworldServiceName).(HelloworldService)

			a.Get("/:name", func(c *fiber.Ctx) error {
				c.Send([]byte(service.SayHello(c.Params("name"))))
				return nil
			})
		})
}
