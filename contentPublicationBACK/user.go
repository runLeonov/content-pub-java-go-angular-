package contentPublicationBACK

type User struct {
	ID       int    `json:"-" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"-'" db:"password"`
	Name     string `json:"name" db:"nickName"`
	Role     string `json:"role" db:"role" gorm:"default:'USER'"`
	Likes    []Like `json:"likes"`
}

type Like struct {
	TitleContent   TitleContent `json:"-"`
	TitleContentID uint         `json:"titleContentId"`
	User           User         `json:"-"`
	UserID         uint         `json:"userId"`
}
