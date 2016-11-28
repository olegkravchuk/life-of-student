package services

import (
	"github.com/olegkravchuk/life_of_student/models"
	"github.com/olegkravchuk/life_of_student/utils"
)

type ContactService struct {
	contact models.Contact
}

func (service ContactService) CreateContact(contact models.Contact) error {
	err := utils.DB.C("contact").Insert(contact)
	return err
}
