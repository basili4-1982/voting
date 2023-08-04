package poll

import "github.com/google/uuid"

type VoteDto struct {
	PollID     *uuid.UUID `json:"poll_id"`
	QuestionID *uuid.UUID `json:"question_id"`
}
