package storage

import (
	"database/sql"
	"fmt"
	"voting/pkg/poll"

	"github.com/google/uuid"
)

type DbRep struct {
	db *sql.DB
}

func NewDbRep(db *sql.DB) *DbRep {
	return &DbRep{db: db}
}

func (d DbRep) Create(dto poll.PollDTO) (poll.PollDTO, error) {
	const insertPoll = `insert into poll (id, title) VALUES ($1,$2)`
	const insertQuestion = `insert into question(id,poll_id,title) VALUES ($1,$2,$3)`
	tx, err := d.db.Begin()
	if err != nil {
		return poll.PollDTO{}, err
	}
	defer tx.Rollback()
	dto.ID = uuid.New()
	_, err = tx.Exec(insertPoll, dto.ID, dto.Title)
	if err != nil {
		return poll.PollDTO{}, err
	}

	for i, _ := range dto.Questions {
		dto.Questions[i].ID = uuid.New()
		q := dto.Questions[i]
		_, err = tx.Exec(insertQuestion, q.ID, dto.ID, q.Title)
		if err != nil {
			return poll.PollDTO{}, err
		}
	}

	tx.Commit()
	return dto, err
}

func (d DbRep) Get(id uuid.UUID) (poll.PollDTO, error) {
	const selectPollAndQuestion = `select  poll.id    as poll_id,
										   poll.title as poll_title,
										   q.id       as question_id,
										   q.title    as question_title,
										   q.vote     as question_vote
									from poll
											 inner join question q on poll.id = q.poll_id
									where poll.id = $1`

	rows, err := d.db.Query(selectPollAndQuestion, id)
	if err != nil {
		return poll.PollDTO{}, err
	}

	defer rows.Close()

	p := poll.PollDTO{}
	for rows.Next() {
		q := poll.QuestionDTO{}
		err = rows.Scan(&p.ID, &p.Title, &q.ID, &q.Title, &q.Vote)
		if err != nil {
			return poll.PollDTO{}, err
		}
		p.Questions = append(p.Questions, q)
	}

	return p, nil
}

func (d DbRep) Vote(dto poll.VoteDto) (poll.PollDTO, error) {
	const voteQuestion = `update question
					set vote= vote + 1
					where id = $1`

	_, err := d.db.Exec(voteQuestion, dto.QuestionID)
	if err != nil {
		return poll.PollDTO{}, err
	}

	return d.Get(*dto.PollID)
}

func OpenDB(host string, port int, username, password, dbName string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, username, password, dbName,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
