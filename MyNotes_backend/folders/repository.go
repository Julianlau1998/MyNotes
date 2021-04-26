package folders

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

func (r *Repository) GetFolders() ([]models.Folder, error) {
	var folders []models.Folder

	jsonFile, err := os.Open("json_files/folders.json")
	if err != nil {
		fmt.Println(err)
		return folders, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return folders, err
	}

	err = json.Unmarshal(byteValue, &folders)
	if err != nil {
		fmt.Println(err)
		return folders, err
	}
	return folders, err
}

func (r *Repository) GetById(id string) (models.Folder, error) {
	var folder models.Folder
	var folders []models.Folder

	jsonFile, err := os.Open("json_files/folders.json")
	if err != nil {
		fmt.Println(err)
		return folder, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return folder, err
	}
	err = json.Unmarshal(byteValue, &folders)
	if err != nil {
		fmt.Println(err)
		return folder, err
	}

	for i := 0; i < len(folders); i++ {
		if folders[i].ID == id {
			folder = folders[i]
		}
	}
	return folder, err
}

func (r *Repository) Post(folder *models.Folder) (*models.Folder, error) {
	var folders []models.Folder

	jsonFile, err := os.Open("json_files/folders.json")
	if err != nil {
		fmt.Println(err)
		return folder, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return folder, err
	}
	err = json.Unmarshal(byteValue, &folders)
	if err != nil {
		fmt.Println(err)
		return folder, err
	}

	folders = append(folders, *folder)
	byteCategories, err := json.Marshal(folders)
	if err != nil {
		fmt.Print(err)
		return folder, err
	}

	err = ioutil.WriteFile("json_files/folders.json", byteCategories, 0644)
	if err != nil {
		fmt.Println(err)
		return folder, err
	}
	return folder, nil
}

func (r *Repository) Update(folder *models.Folder, ID string, userID string) (models.Folder, error) {
	var folders []models.Folder

	jsonFile, err := os.Open("json_files/folders.json")
	if err != nil {
		fmt.Println(err)
		return *folder, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return *folder, err
	}
	err = json.Unmarshal(byteValue, &folders)
	if err != nil {
		fmt.Println(err)
		return *folder, err
	}
	var returnedFolder models.Folder
	for i := 0; i < len(folders); i++ {
		if folders[i].ID == ID && folders[i].UserID == userID {
			folders[i].Title = *&folder.Title
			folders[i].Color = *&folder.Color
			folders[i].Type = *&folder.Type
			returnedFolder = folders[i]
		}
	}

	byteCategories, err := json.Marshal(folders)
	if err != nil {
		fmt.Print(err)
		return *folder, err
	}

	err = ioutil.WriteFile("json_files/folders.json", byteCategories, 0644)
	if err != nil {
		fmt.Println(err)
		return *folder, err
	}

	return returnedFolder, err
}

func (r *Repository) Delete(folder models.Folder, id string, userID string) (models.Folder, error) {
	var folders []models.Folder

	jsonFile, err := os.Open("json_files/folders.json")
	if err != nil {
		fmt.Println(err)
		return folder, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return folder, err
	}
	err = json.Unmarshal(byteValue, &folders)
	if err != nil {
		fmt.Println(err)
		return folder, err
	}
	var emptyFolder models.Folder
	for i := 0; i < len(folders); i++ {
		if folders[i].ID == folder.ID && folders[i].UserID == userID && len(folders) >= 2 {
			if i != len(folders)-1 {
				folders = append(folders[:i], folders[i+1:]...)
			} else {
				folders = folders[:len(folders)-1]
				folders = append(folders, emptyFolder)
			}
		} else if folders[i].ID == folder.ID && folders[i].UserID == userID && len(folders) < 2 {
			folders = append(folders[:0], emptyFolder)
		}
		if folders[i].ID == folder.ID && folders[i].UserID != userID {
			folder.ID = "unauthorized"
			return folder, err
		}
	}

	byteCategories, err := json.Marshal(folders)
	if err != nil {
		fmt.Println(err)
		return folder, err
	}

	err = ioutil.WriteFile("json_files/folders.json", byteCategories, 0644)
	if err != nil {
		fmt.Println(err)
		return folder, err
	}

	folder.ID = ""
	return folder, err
}
