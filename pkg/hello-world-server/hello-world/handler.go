package helloworld

import (
	"github.com/godome/godome/pkg/component/module"
	fiberPlugin "github.com/godome/plugins/pkg/fiber-plugin"

	"github.com/gofiber/fiber/v2"
)

func newHelloworldHandler(m module.Module) fiberPlugin.FiberHandler {
	return fiberPlugin.
		NewFiberHandler().
		AddRoute(func(a *fiber.App) {
			service := m.GetProvider(HelloworldServiceName).(HelloworldService)

			a.Get("/:name", func(c *fiber.Ctx) error {
				c.Send([]byte(service.SayHello(c.Params("name"))))
				return nil
			})
		})
}
