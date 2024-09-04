package converter

import (
	"example.com/m/internal/api/note"
	"example.com/m/internal/model"
	desc "example.com/m/pkg/noteAuth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToNoteFromService(user *model.User) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.Note{
		user: user.Username
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

//func ToNoteInfoFromService(info string) *desc.NoteInfo {
//	return &desc.NoteInfo{Title: info}
//}

func ToNoteInfoFromDesc(info *desc.CreateRequest) *model.User {
	return &model.User{Id: 123123, Username: info.Name, Url: info.Email, Password: "15291529Roma", Role: int(info.Role)}
}
