package note

import (
	"context"
	"example.com/m/internal/converter"
	"example.com/m/internal/model"
	"log"

	desc "example.com/m/pkg/noteAuth_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	// Непосредственное использование полей из req.Info для конвертации или обработки данных
	id, err := i.noteService.Create(ctx, converter.ToNoteInfoFromDesc(req).Username, converter.ToNoteInfoFromDesc(req.Password).Password)

	if err != nil {
		return nil, err
	}

	log.Printf("inserted note with id: %d", id)

	return , nil
}
