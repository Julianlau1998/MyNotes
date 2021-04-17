package lists

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"notesBackend/models"
	"os"
)

type Repository struct {
}

func NewRepository() Repository {
	return Repository{}
}

func (r *Repository) GetLists() ([]models.List, error) {
	var lists []models.List

	jsonFile, err := os.Open("json_files/lists.json")
	if err != nil {
		fmt.Println(err)
		return lists, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return lists, err
	}

	err = json.Unmarshal(byteValue, &lists)
	if err != nil {
		fmt.Println(err)
		return lists, err
	}
	return lists, err
}

func (r *Repository) GetListById(id string) (models.List, error) {
	var list models.List
	var lists []models.List

	jsonFile, err := os.Open("json_files/lists.json")
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return list, err
	}

	err = json.Unmarshal(byteValue, &lists)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	for i := 0; i < len(lists); i++ {
		if lists[i].ID == id {
			list = lists[i]
		}
	}
	return list, err
}

func (r *Repository) PostList(list *models.List) (*models.List, error) {
	var lists []models.List

	jsonFile, err := os.Open("json_files/lists.json")
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	err = json.Unmarshal(byteValue, &lists)
	if err != nil {
		fmt.Println(err)
		return list, err
	}

	lists = append(lists, *list)
	byteCategories, err := json.Marshal(lists)
	if err != nil {
		fmt.Print(err)
		return list, err
	}

	err = ioutil.WriteFile("json_files/lists.json", byteCategories, 0644)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	return list, nil
}

func (r *Repository) updateList(list *models.List, ID string, userID string) (models.List, error) {
	var lists []models.List

	jsonFile, err := os.Open("json_files/lists.json")
	if err != nil {
		fmt.Println(err)
		return *list, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return *list, err
	}
	err = json.Unmarshal(byteValue, &lists)
	if err != nil {
		fmt.Println(err)
		return *list, err
	}
	var returnedList models.List
	for i := 0; i < len(lists); i++ {
		if lists[i].ID == ID && lists[i].UserID == userID {
			lists[i].Title = *&list.Title
			lists[i].List = *&list.List
			lists[i].DoneItems = *&list.DoneItems
			returnedList = lists[i]
		}
	}

	byteCategories, err := json.Marshal(lists)
	if err != nil {
		fmt.Print(err)
		return *list, err
	}

	err = ioutil.WriteFile("json_files/lists.json", byteCategories, 0644)
	if err != nil {
		fmt.Println(err)
		return *list, err
	}

	return returnedList, err
}

func (r *Repository) DeleteList(list models.List, id string, userID string) (models.List, error) {
	var lists []models.List

	jsonFile, err := os.Open("json_files/lists.json")
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	err = json.Unmarshal(byteValue, &lists)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	var emptyList models.List
	for i := 0; i < len(lists); i++ {
		if lists[i].ID == list.ID && lists[i].UserID == userID && len(lists) >= 2 {
			if i != len(lists)-1 {
				lists = append(lists[:i], lists[i+1:]...)
			} else {
				lists = lists[:len(lists)-1]
				lists = append(lists, emptyList)
			}
		} else if lists[i].ID == list.ID && lists[i].UserID == userID && len(lists) < 2 {
			lists = append(lists[:0], emptyList)
		}
		if lists[i].ID == list.ID && lists[i].UserID != userID {
			list.ID = "unauthorized"
			return list, err
		}
	}

	byteCategories, err := json.Marshal(lists)
	if err != nil {
		fmt.Println(err)
		return list, err
	}

	err = ioutil.WriteFile("json_files/lists.json", byteCategories, 0644)
	if err != nil {
		fmt.Println(err)
		return list, err
	}

	list.ID = ""
	return list, err
}
