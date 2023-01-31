package contentPublicationBACK

type User struct {
	ID       int    `json:"-" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Name     string `json:"name" db:"nickName"`
	Role     string `json:"role" db:"role"`
	Likes    []Like `json:"likes"`
}

type Like struct {
	TitleContent   TitleContent `json:"-"`
	TitleContentID uint         `json:"titleContentId"`
	User           User         `json:"-"`
	UserID         uint         `json:"userId"`
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzM1OTQwNTYsImlhdCI6MTY3MzU1MDg1NiwidXNlcl9pZCI6Mn0.Z-ziOWQGEtFken5e3oic_Lknx1kLGmm91ZlWe1gwZgU
