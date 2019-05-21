package api

import (
	"github.com/google/uuid"
)

type mapElement struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type configResponse struct {
	DifficultyLevels []mapElement `json:"difficulty_levels"`
	QuestionTypes    []mapElement `json:"question_types"`
	Genders          []mapElement `json:"genders"`
}

func IsValidUUID(u string) error {
	_, err := uuid.Parse(u)
	return err
}

type Response struct {
	Message string `json:"message"`
}
