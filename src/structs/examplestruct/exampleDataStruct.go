package examplestruct

import "time"

// MyGuests :
type MyGuests struct {
	ID        int       `json:"id" gorm:"column:id; PRIMARY_KEY"`
	FirstName string    `json:"firstname" gorm:"column:firstname" valid:"required" form:"firstname"`
	Lastname  string    `json:"lastname" gorm:"column:lastname" valid:"required" form:"lastname"`
	Email     string    `json:"email" gorm:"column:email" valid:"required,email" form:"email"`
	RegDate   time.Time `json:"reg_date" gorm:"column:reg_date; default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP""`
}

// MyGuestsIndexResponse :
type MyGuestsIndexResponse struct {
	Status int        `json:"status"`
	Data   []MyGuests `json:"data"`
}

// MyGuestsCreateResponse :
type MyGuestsCreateResponse struct {
	Status int      `json:"status"`
	Data   MyGuests `json:"data"`
}
