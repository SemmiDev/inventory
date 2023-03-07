package handler

import (
	"github.com/gofiber/fiber/v2"
	"inventory/entity"
	"inventory/repository"
	"strconv"
)

type ItemHandler struct {
	Repository *repository.ItemRepository
	App        *fiber.App
}

func NewItemHandler(repository *repository.ItemRepository, app *fiber.App) *ItemHandler {
	return &ItemHandler{Repository: repository, App: app}
}

func (h *ItemHandler) SetupRoutes() {
	h.App.Get("/items", h.FindAll)
	h.App.Get("/items/:id", h.FindById)
	h.App.Post("/items", h.Save)
	h.App.Put("/items/:id", h.Update)
	h.App.Delete("/items/:id", h.Delete)
	h.App.Get("/items/status/:status", h.FindByStatus)
}

func (h *ItemHandler) FindAll(c *fiber.Ctx) error {
	items, err := h.Repository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(items)
}

func (h *ItemHandler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	idInUint, err := strconv.ParseUint(id, 10, 64)
	
	item, err := h.Repository.FindById(uint(idInUint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(item)
}

func (h *ItemHandler) Save(c *fiber.Ctx) error {
	item := new(entity.Item)
	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := h.Repository.Save(item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(item)
}

func (h *ItemHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	idInUint, err := strconv.ParseUint(id, 10, 64)

	item, err := h.Repository.FindById(uint(idInUint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := h.Repository.Update(&item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(item)
}

func (h *ItemHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	idInUint, err := strconv.ParseUint(id, 10, 64)

	item, err := h.Repository.FindById(uint(idInUint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := h.Repository.Delete(&item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *ItemHandler) FindByStatus(c *fiber.Ctx) error {
	status := c.Params("status")
	items, err := h.Repository.FindByStatus(entity.Status(status))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(items)
}
