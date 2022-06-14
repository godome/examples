package post

import (
	"github.com/godome/godome/pkg/component/module"
	fiberPlugin "github.com/godome/plugins/pkg/fiber-plugin"

	"github.com/gofiber/fiber/v2"
)

func newPostHandler(m module.Module) fiberPlugin.FiberHandler {
	return fiberPlugin.
		NewFiberHandler().
		AddRoute(func(a *fiber.App) {
			service := m.GetProvider(PostServiceName).(PostService)

			a.Get("/post/:id", func(c *fiber.Ctx) error {
				foundPost, err := service.GetPost(c.Params("id"))
				if err != nil {
					return err
				}
				c.JSON(foundPost)
				return nil
			})

			a.Post("/post", func(c *fiber.Ctx) error {
				post := new(PostEntity)
				if err := c.BodyParser(post); err != nil {
					return c.Status(400).SendString(err.Error())
				}
				post, err := service.CreatePost(post.Name, post.Description)
				if err != nil {
					return c.Status(400).SendString(err.Error())
				}
				c.JSON(post)
				return nil
			})
		})
}
