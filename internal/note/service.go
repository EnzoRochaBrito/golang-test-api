package note

import (
	repo "api/internal/adapters/postgresql/sqlc"
	"context"
)

type Service interface {
	CreateNote(ctx context.Context, note CreateNote) (repo.Note, error)
	GetNote(ctx context.Context, id string) (repo.Note, error)
	DeleteNote(ctx context.Context, id string) error
}

type svc struct {
	repo *repo.Queries
}

func NewService(repo *repo.Queries) Service {
	return &svc{
		repo: repo,
	}
}

func (r *svc) CreateNote(ctx context.Context, note CreateNote) (repo.Note, error) {

	newNote := repo.CreateNoteParams{
		ID: note.Title,
		Title: note.Title,
		Body: note.Body,
	}

	return r.repo.CreateNote(ctx, newNote)
}

func (r *svc) GetNote(ctx context.Context, id string) (repo.Note, error) {
	return r.repo.GetNote(ctx, id)
}

func (r *svc) DeleteNote(ctx context.Context, id string) error {
	return r.repo.DeleteNote(ctx, id)
}
