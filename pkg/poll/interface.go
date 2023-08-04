package poll

import "github.com/google/uuid"

type Repository interface {
	Create(PollDTO) (PollDTO, error)
	Get(uuid.UUID) (PollDTO, error)
	Vote(VoteDto) (PollDTO, error)
}
