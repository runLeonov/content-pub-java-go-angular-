package contentPublicationBACK

type User struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Name     string `json:"name" db:"nickName"`
	Img      string `json:"img64" db:"img64"`
	Role     string `json:"role" db:"role" gorm:"default:'USER'"`
	Likes    []Like `json:"likes"`
}

type Like struct {
	TitleContent   TitleContent `json:"-"`
	TitleContentID uint         `json:"titleContentId"`
	User           User         `json:"-"`
	UserID         uint         `json:"userId"`
}
