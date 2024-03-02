package model

import (
	"gorm.io/gorm"
)

type TeacherQuestion struct {
	gorm.Model
	Question   string `json:"question"`
	Answer1    string `json:"answer_1"`
	Answer2    string `json:"answer_2"`
	Answer3    string `json:"answer_3"`
	Answer4    string `json:"answer_4"`
	TrueAnswer string `json:"true_answer"`
}
