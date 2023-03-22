package contentPublicationBACK

import "time"

type User struct {
	ID           int       `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	Password     string    `json:"password" db:"password"`
	Name         string    `json:"nickName" db:"nickName"`
	Img          string    `json:"img64" db:"img64"`
	Role         string    `json:"role" db:"role" gorm:"default:'USER'"`
	Likes        []Like    `json:"likes"`
	CreationDate time.Time `json:"creationDate"`
	LastName     string    `json:"lastName"`
	FirstName    string    `json:"firstName"`
}

type Like struct {
	TitleContent   TitleContent `json:"-"`
	TitleContentID uint         `json:"titleContentId"`
	User           User         `json:"-"`
	UserID         uint         `json:"userId"`
}
