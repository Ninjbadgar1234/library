package controller

import (
	"example.com/book/src/database"
	"example.com/book/src/model"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct{}

func SetCategoryControllers(app *fiber.App) {
	var c CategoryController
	t := app.Group("/category")
	t.Post("", c.Create)
	t.Get("", c.List)
	t.Get("/:id", c.Get)
	t.Put("/:id", c.Update)
	t.Delete("/:id", c.Delete)
}

func (tc *CategoryController) List(c *fiber.Ctx) error {
	db := database.DBinstanse
	var data []model.Category
	db.Find(&data)
	return c.JSON(fiber.Map{"status": "success", "message": "done", "result": data})
}

func (tc *CategoryController) Create(c *fiber.Ctx) error {
	db := database.DBinstanse
	data := new(model.Category)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error(), "result": fiber.Map{}})
	}

	if errors := ValidateRequest(*data); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "bad request", "result": errors})
	}

	if err := db.Create(&data).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error(), "result": fiber.Map{}})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "done", "result": data})
}

func (tc *CategoryController) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBinstanse
	var data model.Category
	db.Find(&data, id)

	if data.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Мэдээлэл олдсонгүй"})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "done", "result": data})
}

func (tc *CategoryController) Update(c *fiber.Ctx) error {
	db := database.DBinstanse
	id := c.Params("id")
	data := new(model.Category)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error(), "result": fiber.Map{}})
	}

	if errors := ValidateRequest(*data); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "bad request", "result": errors})
	}

	err := db.Model(&data).Where("id=?", id).Updates(data).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error(), "result": fiber.Map{}})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Мэдээлэл засагдлаа", "result": fiber.Map{}})
}

func (tc *CategoryController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBinstanse
	var data model.Category

	if err := db.Delete(&data, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error(), "result": fiber.Map{}})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Мэдээлэл устлаа", "result": fiber.Map{}})
}
