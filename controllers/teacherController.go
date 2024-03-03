package controllers

import (
	"EOP/database"
	"EOP/middleware"
	"EOP/model"

	"github.com/gofiber/fiber/v2"
)

func AddQuestions(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": err})
	}
	var Exam model.TeacherQuestion
	c.BodyParser(&Exam)
	Exam.UserID = user.UserID
	err = database.DB.Db.Save(Exam).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "database error", "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "exam added success", "data": Exam})
}

func UpdateQuestions(c *fiber.Ctx) error {
	_, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": err})
	}
	var UpdateQuess model.TeacherQuestion
	//var OldQuess model.TeacherQuestion
	c.BodyParser(&UpdateQuess)
	err = database.DB.Db.Model(model.TeacherQuestion{}).Where("id = ?", UpdateQuess.ID).Updates(model.TeacherQuestion{Question: UpdateQuess.Question, Answer1: UpdateQuess.Answer1,
		Answer2: UpdateQuess.Answer2, Answer3: UpdateQuess.Answer3, Answer4: UpdateQuess.Answer4, TrueAnswer: UpdateQuess.TrueAnswer}).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "data": UpdateQuess})

}

func DeleteQuestions(c *fiber.Ctx) error {
	_, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": err})
	}
	id := c.BodyParser("id")
	var quess model.TeacherQuestion
	err = database.DB.Db.Where("id=?", id).Find(&quess).Error
	if err != nil {
		return err
	}
	err = database.DB.Db.Delete(&quess).Error
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"status": "delete success"})
}
