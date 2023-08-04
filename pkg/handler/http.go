package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"voting/pkg/poll"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type HttpErr struct {
	Errors []string
}

func MakeHttpErr(errors ...string) string {
	e := HttpErr{
		Errors: errors,
	}
	d, _ := json.Marshal(e)

	return string(d)
}

func CreatePoll(servicePoll *poll.Poll) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}
		logrus.Trace("CreatePoll: Пришел запрос Тело:", body)
		p := poll.PollDTO{}
		err = json.Unmarshal(body, &p)
		if err != nil {
			logrus.Error("CreatePoll: не смог разобрать тело запроса", err.Error())
			http.Error(w, MakeHttpErr(err.Error()), http.StatusBadRequest)
			return
		}
		v := poll.CreateValidator{}
		warnings := v.Validate(p)
		if len(warnings) > 0 {
			logrus.Error("CreatePoll: не прошел валидацию:", strings.Join(warnings, "/n"))
			http.Error(w, MakeHttpErr(warnings...), http.StatusBadRequest)
			return
		}

		dto, err := servicePoll.Create(p)
		if err != nil {
			logrus.Error("CreatePoll: не смог создать опрос:", err.Error())
			http.Error(w, MakeHttpErr(err.Error()), http.StatusInternalServerError)
		}

		d, err := json.Marshal(dto)
		if err != nil {
			logrus.Error("CreatePoll: не смог разобрать ответ сервиса:", err.Error())
			http.Error(w, MakeHttpErr(err.Error()), http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(d)
		logrus.Trace("CreatePoll: ответил", string(d))
	}
}

func GetResult(servicePoll *poll.Poll) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.URL.Query().Get("id"))
		if err != nil {
			logrus.Error("GetResult: не смог разобрать параметр", err.Error())
			http.Error(w, MakeHttpErr(err.Error()), http.StatusBadRequest)
		}
		p, err := servicePoll.Get(id)
		if err != nil {
			logrus.Error("GetResult: не смог получить опрос", err.Error())
			http.Error(w, MakeHttpErr(err.Error()), http.StatusInternalServerError)
		}
		d, err := json.Marshal(p)
		if err != nil {
			logrus.Error("GetResult: не смог разобрать ответ сервиса:", err.Error())
			http.Error(w, MakeHttpErr(err.Error()), http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(d)
		logrus.Trace("GetResult: ответил", string(d))
	}
}

func Poll(servicePoll *poll.Poll) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}
		voteDto := poll.VoteDto{}
		err = json.Unmarshal(body, &voteDto)
		if err != nil {
			logrus.Error("Poll: не смог разобрать тело запроса", err.Error())
			http.Error(w, MakeHttpErr(err.Error()), http.StatusBadRequest)
			return
		}
		v := poll.VoteValidator{}
		warnings := v.Validate(voteDto)
		if len(warnings) > 0 {
			logrus.Error("Poll: не прошел валидацию:", strings.Join(warnings, "/n"))
			http.Error(w, MakeHttpErr(warnings...), http.StatusBadRequest)
			return
		}
		p, err := servicePoll.Vote(voteDto)
		if err != nil {
			logrus.Error("Poll: не смог отметить голос:", strings.Join(warnings, "/n"))
			http.Error(w, MakeHttpErr(err.Error()), http.StatusInternalServerError)
		}
		d, err := json.Marshal(p)
		if err != nil {
			logrus.Error("Poll: не смог разобрать ответ сервиса:", err.Error())
			http.Error(w, MakeHttpErr(err.Error()), http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(d)
		logrus.Trace("Poll: Ответил:", string(d))
	}
}

func Info(w http.ResponseWriter, r *http.Request) {
	const txt = `
POST /createPoll/ создает голосование c вариантами ответов

Запрос 

{
  "title": "Опрос",
  "questions": [
    {
      "title": "Вопрос 1"
    },
    {
      "title": "Вопрос 2"
    },
    {
      "title": "Вопрос 3"
    }
  ]
}

GET /api/getResult/?id=0ee6eb47-4eb0-4293-ba47-e1cb9a10a7bb -  возвращает результат по конкретному голосованию id  


POST /api/poll/ позволяет проголосовать за конкретный вариант
Запрос

{
  "poll_id": "0ee6eb47-4eb0-4293-ba47-e1cb9a10a7bb",
  "question_id": "5ba1098d-77f0-42c9-9610-1b9c3fd411d0"
}

`
	w.Write([]byte(txt))
}
