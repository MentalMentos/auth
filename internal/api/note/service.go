package note

import (
	"example.com/m/internal/service"
	desc "example.com/m/pkg/noteAuth_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	noteService service.Service
}

func NewImplementation(noteService service.Service) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
