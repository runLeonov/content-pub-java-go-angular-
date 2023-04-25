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
	Comments     []Comment `json:"comments"`
	CreationDate time.Time `json:"creationDate"`
	LastName     string    `json:"lastName"`
	FirstName    string    `json:"firstName"`
	Banned       bool      `json:"banned" gorm:"default:false"`
}

type Like struct {
	TitleContent   TitleContent `json:"titleContent"`
	TitleContentID uint         `json:"titleContentId"`
	User           User         `json:"-"`
	UserID         uint         `json:"userId"`
}

type Comment struct {
	ID             int          `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	CommentVal     string       `json:"commentVal"`
	CreationDate   time.Time    `json:"creationDate"`
	User           User         `json:"user"`
	UserID         uint         `json:"userId"`
	TitleContent   TitleContent `json:"titleContent"`
	TitleContentID uint         `json:"titleContentId"`
}
