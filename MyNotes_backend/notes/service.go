package notes

import (
	"notesBackend/models"

	uuid "github.com/nu7hatch/gouuid"
)

type Service struct {
	noteRepository Repository
}

func NewService(noteRepository Repository) Service {
	return Service{noteRepository: noteRepository}
}

func (s *Service) GetNotes(userID string) ([]models.Note, error) {
	notes, err := s.noteRepository.GetNotes()
	var notesOfUser []models.Note
	for _, note := range notes {
		if note.UserID == userID {
			notesOfUser = append(notesOfUser, note)
		}
	}
	return notesOfUser, err
}

func (s *Service) GetByFolder(folderID string, userID string) ([]models.Note, error) {
	notes, err := s.noteRepository.GetByFolder()
	if err != nil {
		return nil, err
	}
	var notesOfUser []models.Note
	for _, note := range notes {
		if note.UserID == userID && note.FolderID == folderID {
			notesOfUser = append(notesOfUser, note)
		}
	}
	return notesOfUser, nil
}

func (s *Service) GetNoteById(id string, userID string) (models.Note, error) {
	note, err := s.noteRepository.GetNoteById(id)

	if note.UserID == userID {
		return note, err
	}
	note.Title = ""
	note.Note = ""
	note.ID = ""
	note.UserID = ""
	return note, err
}

func (s *Service) Post(note *models.Note, userID string) (*models.Note, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return note, err
	}
	note.ID = id.String()
	note.UserID = userID
	return s.noteRepository.PostCategory(note)
}

func (s *Service) updateNote(id string, note *models.Note, userID string) (models.Note, error) {
	newNote, err := s.noteRepository.updateNote(note, id, userID)
	return newNote, err
}

func (s *Service) DeleteNote(ID string, userID string) (models.Note, error) {
	var note models.Note
	note.ID = ID
	return s.noteRepository.deleteNote(note, ID, userID)
}
