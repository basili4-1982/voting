# Сервис опросов

```
POST /createPoll/
```

Cоздает голосование c вариантами ответов

*Запрос*

```
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
```

```GET /api/getResult/?id=0ee6eb47-4eb0-4293-ba47-e1cb9a10a7bb```

Возвращает результат по конкретному голосованию id

```POST /api/poll/```
Позволяет проголосовать за конкретный вариант

*Запрос*
```
{
    "poll_id": "0ee6eb47-4eb0-4293-ba47-e1cb9a10a7bb",
    "question_id": "5ba1098d-77f0-42c9-9610-1b9c3fd411d0"
}
```