package note

import (
	"context"
	"log"

	"example.com/m/internal/model"

	"example.com/m/internal/converter"
	desc "example.com/m/pkg/note_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest, us *model.User) (*desc.GetResponse, error) {
	noteObj, err := i.noteService.Get(ctx, us.Id)
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, title: %s, body: %s, created_at: %v, updated_at: %v\n", noteObj.ID, noteObj.Info.Title, noteObj.Info.Content, noteObj.CreatedAt, noteObj.UpdatedAt)

	return &desc.GetResponse{
		Note: converter.ToNoteFromService(noteObj),
	}, nil
}
