package poll

type CreateValidator struct {
}

func (v CreateValidator) Validate(dto PollDTO) []string {
	warnings := []string{}

	if len(dto.Title) == 0 {
		warnings = append(warnings, "Необходимо заполнить поле Title")
	}

	if len(dto.Questions) < 2 {
		warnings = append(warnings, "Необходимо добавить не менее 2 вопросов")
	}

	for _, q := range dto.Questions {
		if len(q.Title) == 0 {
			warnings = append(warnings, "Необходимо в вопросе заполнить поле Title")
		}
	}

	return warnings
}

type VoteValidator struct {
}

func (v VoteValidator) Validate(dto VoteDto) []string {
	warnings := []string{}
	if dto.PollID == nil {
		warnings = append(warnings, "Необходимо заполнить поле PollID")
	}

	if dto.QuestionID == nil {
		warnings = append(warnings, "Необходимо заполнить поле QuestionID")
	}

	return warnings
}
