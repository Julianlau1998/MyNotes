package notes

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

func (r *Repository) GetNotes() ([]models.Note, error) {
	var notes []models.Note

	jsonFile, err := os.Open("json_files/notes.json")
	if err != nil {
		fmt.Println(err)
		return notes, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return notes, err
	}

	err = json.Unmarshal(byteValue, &notes)
	if err != nil {
		fmt.Println(err)
		return notes, err
	}
	return notes, err
}

func (r *Repository) GetNoteById(id string) (models.Note, error) {
	var note models.Note
	var notes []models.Note

	jsonFile, err := os.Open("json_files/notes.json")
	if err != nil {
		fmt.Println(err)
		return note, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return note, err
	}
	err = json.Unmarshal(byteValue, &notes)
	if err != nil {
		fmt.Println(err)
		return note, err
	}

	for i := 0; i < len(notes); i++ {
		if notes[i].ID == id {
			note = notes[i]
		}
	}
	return note, err
}

func (r *Repository) PostCategory(note *models.Note) (*models.Note, error) {
	var notes []models.Note

	jsonFile, err := os.Open("json_files/notes.json")
	if err != nil {
		fmt.Println(err)
		return note, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return note, err
	}
	err = json.Unmarshal(byteValue, &notes)
	if err != nil {
		fmt.Println(err)
		return note, err
	}

	notes = append(notes, *note)
	byteCategories, err := json.Marshal(notes)
	if err != nil {
		fmt.Print(err)
		return note, err
	}

	err = ioutil.WriteFile("json_files/notes.json", byteCategories, 0644)
	if err != nil {
		fmt.Println(err)
		return note, err
	}
	return note, nil
}

func (r *Repository) updateNote(note *models.Note, ID string, userID string) (models.Note, error) {
	var notes []models.Note

	jsonFile, err := os.Open("json_files/notes.json")
	if err != nil {
		fmt.Println(err)
		return *note, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return *note, err
	}
	err = json.Unmarshal(byteValue, &notes)
	if err != nil {
		fmt.Println(err)
		return *note, err
	}
	var returnedNote models.Note
	for i := 0; i < len(notes); i++ {
		if notes[i].ID == ID && notes[i].UserID == userID {
			notes[i].Title = *&note.Title
			notes[i].Note = *&note.Note
			returnedNote = notes[i]
		}
	}

	byteCategories, err := json.Marshal(notes)
	if err != nil {
		fmt.Print(err)
		return *note, err
	}

	err = ioutil.WriteFile("json_files/notes.json", byteCategories, 0644)
	if err != nil {
		fmt.Println(err)
		return *note, err
	}

	return returnedNote, err
}

func (r *Repository) deleteNote(note models.Note, id string, userID string) (models.Note, error) {
	var notes []models.Note

	jsonFile, err := os.Open("json_files/notes.json")
	if err != nil {
		fmt.Println(err)
		return note, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return note, err
	}
	err = json.Unmarshal(byteValue, &notes)
	if err != nil {
		fmt.Println(err)
		return note, err
	}
	var emptyNote models.Note
	for i := 0; i < len(notes); i++ {
		if notes[i].ID == note.ID && notes[i].UserID == userID && len(notes) >= 2 {
			if i != len(notes)-1 {
				notes = append(notes[:i], notes[i+1:]...)
			} else {
				notes = notes[:len(notes)-1]
				notes = append(notes, emptyNote)
			}
		} else if notes[i].ID == note.ID && notes[i].UserID == userID && len(notes) < 2 {
			notes = append(notes[:0], emptyNote)
		}
		if notes[i].ID == note.ID && notes[i].UserID != userID {
			note.ID = "unauthorized"
			return note, err
		}
	}

	byteCategories, err := json.Marshal(notes)
	if err != nil {
		fmt.Println(err)
		return note, err
	}

	err = ioutil.WriteFile("json_files/notes.json", byteCategories, 0644)
	if err != nil {
		fmt.Println(err)
		return note, err
	}

	note.ID = ""
	return note, err
}
