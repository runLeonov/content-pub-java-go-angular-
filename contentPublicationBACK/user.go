package contentPublicationBACK

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"-" db:"id"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
	Name     string `json:"name" db:"nickName" binding:"required" `
	Role     string `json:"role" db:"role" binding:"required"`
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzM1OTQwNTYsImlhdCI6MTY3MzU1MDg1NiwidXNlcl9pZCI6Mn0.Z-ziOWQGEtFken5e3oic_Lknx1kLGmm91ZlWe1gwZgU
