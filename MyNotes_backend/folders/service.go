package folders

import (
	"notesBackend/lists"
	"notesBackend/models"
	"notesBackend/notes"

	uuid "github.com/nu7hatch/gouuid"
)

type Service struct {
	folderRepository Repository
	listsService     lists.Service
	noteService      notes.Service
}

func NewService(folderRepository Repository, listsService lists.Service, noteService notes.Service) Service {
	return Service{
		folderRepository: folderRepository,
		listsService:     listsService,
		noteService:      noteService,
	}
}

func (s *Service) GetFolders(userID string) ([]models.Folder, error) {
	folders, err := s.folderRepository.GetFolders()
	var foldersOfUser []models.Folder
	for _, folder := range folders {
		if folder.UserID == userID {
			foldersOfUser = append(foldersOfUser, folder)
		}
	}
	return foldersOfUser, err
}

func (s *Service) GetById(id string, userID string) (models.Folder, error) {
	folder, err := s.folderRepository.GetById(id)

	if folder.UserID == userID {
		return folder, err
	}
	folder.Title = ""
	folder.Color = ""
	folder.Type = ""
	folder.ID = ""
	folder.UserID = ""
	return folder, err
}

func (s *Service) Post(folder *models.Folder, userID string) (*models.Folder, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return folder, err
	}
	folder.ID = id.String()
	folder.UserID = userID

	switch folder.Color {
	case "Red":
		folder.Color = "#C0392B"
	case "Green":
		folder.Color = "#196F3D"
	case "Yellow":
		folder.Color = "#9A7D0A"
	case "Blue":
		folder.Color = "#1F618D"
	case "Gray":
		folder.Color = "#424949"
	}

	return s.folderRepository.Post(folder)
}

func (s *Service) Update(id string, folder *models.Folder, userID string) (models.Folder, error) {
	newFolder, err := s.folderRepository.Update(folder, id, userID)
	return newFolder, err
}

func (s *Service) Delete(ID string, userID string) (models.Folder, error) {
	var folder models.Folder
	folder.ID = ID

	// Delete all lists in folder
	lists, err := s.listsService.GetByFolder(ID, userID)
	if err != nil {
		return folder, err
	}
	for _, list := range lists {
		s.listsService.DeleteList(list.ID, userID)
	}

	//Delete all Notes in folder
	notes, err := s.noteService.GetByFolder(ID, userID)
	if err != nil {
		return folder, err
	}
	for _, note := range notes {
		s.noteService.DeleteNote(note.ID, userID)
	}
	return s.folderRepository.Delete(folder, ID, userID)
}
