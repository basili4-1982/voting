package storage

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"voting/pkg/poll"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func TestDbRep_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		dto poll.PollDTO
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   poll.PollDTO
	}{
		{
			name: "sddsg",
			fields: fields{
				func() *sql.DB {
					psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
						"password=%s dbname=%s sslmode=disable",
						"localhost", 5432, "postgres", "example", "local")
					db, err := sql.Open("postgres", psqlInfo)
					if err != nil {
						panic(err)
					}

					err = db.Ping()
					if err != nil {
						panic(err)
					}
					return db
				}(),
			},
			args: args{
				dto: poll.PollDTO{
					ID:    uuid.MustParse("352cad1b-fe0d-4398-82ae-796adcca43c0"),
					Title: "Вопрос",
					Questions: []poll.QuestionDTO{
						{
							ID:    uuid.New(),
							Title: "Ответ 1",
						},
						{
							ID:    uuid.New(),
							Title: "Ответ 2",
						},
						{
							ID:    uuid.New(),
							Title: "Ответ 3",
						},
					},
				},
			},
			want: poll.PollDTO{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDbRep(tt.fields.db)
			got, err := d.Create(tt.args.dto)
			if err != nil {
				t.Errorf("Create() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestDbRep_Get(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    poll.PollDTO
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				db: func() *sql.DB {
					psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
						"password=%s dbname=%s sslmode=disable",
						"localhost", 5432, "postgres", "example", "local")
					db, err := sql.Open("postgres", psqlInfo)
					if err != nil {
						panic(err)
					}

					err = db.Ping()
					if err != nil {
						panic(err)
					}
					return db
				}(),
			},
			args: args{
				id: uuid.MustParse("352cad1b-fe0d-4398-82ae-796adcca43c0"),
			},
			want: poll.PollDTO{
				ID:    uuid.MustParse("352cad1b-fe0d-4398-82ae-796adcca43c0"),
				Title: "Вопрос",
				Questions: []poll.QuestionDTO{
					{
						ID:    uuid.MustParse("71b0b7ca-e968-4f8b-a43b-9b6e591a9458"),
						Title: "Ответ 1",
						Vote:  0,
					},
					{
						ID:    uuid.MustParse("088479ca-28b8-4585-8d8c-86c6174035a6"),
						Title: "Ответ 2",
						Vote:  0,
					},
					{
						ID:    uuid.MustParse("87b02e12-092c-48f9-996c-8eda77ec3d33"),
						Title: "Ответ 3",
						Vote:  0,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DbRep{
				db: tt.fields.db,
			}
			got, err := d.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
