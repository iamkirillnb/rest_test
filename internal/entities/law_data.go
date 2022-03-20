package entities

import "fmt"

type FormData struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Question string `json:"question"`
}

func ValidateFormData(f *FormData) error {
	err_text := "при валидации данных возникла ошибка в "
	if f.Name == "kenny" {
		return fmt.Errorf("%s NAME", err_text)
	}
	return nil
}