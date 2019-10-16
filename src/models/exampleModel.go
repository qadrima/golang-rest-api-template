package models

import (
	"go-template-rest-api/src/database"
	"go-template-rest-api/src/structs/examplestruct"
)

// ExampleModel :
type ExampleModel struct{}

// GetAll func :
func (m ExampleModel) GetAll() (guests []examplestruct.MyGuests, err error) {

	if err = database.MysqlDB().Table("MyGuests").Find(&guests).Error; err != nil {
		return guests, err
	}

	return guests, nil
}

// Create func :
func (m ExampleModel) Create(guest examplestruct.MyGuests) (newGuest examplestruct.MyGuests, err error) {

	if err = database.MysqlDB().Table("MyGuests").Create(&guest).Error; err != nil {
		return newGuest, err
	}

	newGuest = guest
	return newGuest, nil
}
