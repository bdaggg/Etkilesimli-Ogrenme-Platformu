package controllers

import (
	"EOP/database"
	"EOP/middleware"
	"EOP/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AddQuestions(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": err})
	}
	var Exam model.TeacherQuestion
	log.Println("buraya geldi -1 ----------------")
	c.BodyParser(&Exam)
	Exam.UserID = user.UserID
	log.Println("buraya geldi -2 ----------------")
	err = database.DB.Db.Save(Exam).Error
	if err != nil {
		log.Println("buraya geldi -3 ----------------")
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "database error", "data": err})
	}
	log.Println("buraya geldi -4 ----------------")
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "exam added success", "data": Exam})
}
