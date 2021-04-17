package lists

import (
	"notesBackend/models"

	uuid "github.com/nu7hatch/gouuid"
)

type Service struct {
	listRepository Repository
}

func NewService(listRepository Repository) Service {
	return Service{listRepository: listRepository}
}

func (s *Service) GetLists(userID string) ([]models.List, error) {
	lists, err := s.listRepository.GetLists()
	var listsOfUser []models.List
	for _, list := range lists {
		if list.UserID == userID {
			listsOfUser = append(listsOfUser, list)
		}
	}
	return listsOfUser, err
}

func (s *Service) GetListById(id string, userID string) (models.List, error) {
	list, err := s.listRepository.GetListById(id)

	if list.UserID == userID {
		return list, err
	}
	var emptyList []string
	list.Title = ""
	list.List = emptyList
	list.ID = ""
	list.UserID = ""
	return list, err
}

func (s *Service) Post(list *models.List, userID string) (*models.List, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return list, err
	}
	list.ID = id.String()
	list.UserID = userID
	return s.listRepository.PostList(list)
}

func (s *Service) UpdateList(id string, list *models.List, userID string) (models.List, error) {
	newList, err := s.listRepository.updateList(list, id, userID)
	return newList, err
}

func (s *Service) deleteList(ID string, userID string) (models.List, error) {
	var list models.List
	list.ID = ID
	return s.listRepository.DeleteList(list, ID, userID)
}
