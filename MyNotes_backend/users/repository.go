package users

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

func (r *Repository) GetUsers() ([]models.User, error) {
	var users []models.User

	jsonFile, err := os.Open("json_files/users.json")
	if err != nil {
		fmt.Println(err)
		return users, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return users, err
	}

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println(err)
		return users, err
	}
	return users, err
}

func (r *Repository) PostUser(user *models.User) (*models.User, error) {
	var users []models.User

	jsonFile, err := os.Open("json_files/users.json")
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	var exists bool = false
	for _, storedUser := range users {
		if storedUser.Username == user.Username {
			exists = true
		}
	}

	if exists {
		user.Username = "exists"
		user.ID = ""
		user.Password = ""
	} else {
		users = append(users, *user)
		byteCategories, err := json.Marshal(users)
		if err != nil {
			fmt.Print(err)
			return user, err
		}

		err = ioutil.WriteFile("json_files/users.json", byteCategories, 0644)
		if err != nil {
			fmt.Println(err)
			return user, err
		}
	}
	return user, nil
}
