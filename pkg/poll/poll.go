package poll

import "github.com/google/uuid"

type Poll struct {
	Repository
}

func NewPoll(repository Repository) *Poll {
	return &Poll{Repository: repository}
}

func (p Poll) Get(id uuid.UUID) (PollDTO, error) {
	return p.Repository.Get(id)
}

func (p Poll) Create(dto PollDTO) (PollDTO, error) {
	return p.Repository.Create(dto)
}
