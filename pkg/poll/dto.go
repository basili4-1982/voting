package poll

import (
	"github.com/google/uuid"
)

type QuestionDTO struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Vote  int       `json:"vote"`
}

type PollDTO struct {
	ID        uuid.UUID     `json:"id"`
	Title     string        `json:"title"`
	Questions []QuestionDTO `json:"questions"`
}
