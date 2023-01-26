package contentPublicationBACK

import (
	"time"
)

type Title struct {
	ID             uint         `json:"id"`
	TitleName      string       `json:"titleName"`
	OriginalAuthor string       `json:"originalAuthor"`
	CreationDate   time.Time    `json:"creationDate"`
	Description    string       `json:"description"`
	TitleImgBase64 string       `json:"titleImg"`
	Content        TitleContent `json:"content" gorm:"foreignKey:ID"`
	Categories     []Category   `json:"categories" gorm:"many2many:titles_categories;"`
	Tags           []Tag        `json:"tags" gorm:"many2many:titles_tags;"`
	Serials        []Serial     `json:"serials" gorm:"many2many:titles_serials;"`
}

type TitleContent struct {
	ID      uint    `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	TitleID uint    `json:"titleId"`
	Likes   uint    `json:"likes"`
	Views   uint    `json:"views"`
	Images  []Image `json:"images"`
}

type Image struct {
	ID             uint         `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	ImgBase64      string       `json:"img64"`
	TitleContentID uint         `json:"-"`
	TitleContent   TitleContent `json:"-" `
}

//gorm:"foreignKey:TitleContentID;references:ID"
type Category struct {
	ID     uint     `json:"id"`
	Genre  string   `json:"genre"`
	Titles []*Title `json:"titles" gorm:"many2many:titles_categories;" `
}

type Tag struct {
	ID      uint     `json:"id"`
	TagName string   `json:"tag"`
	Titles  []*Title `json:"titles" gorm:"many2many:titles_tags;"`
}

type Serial struct {
	ID         uint     `json:"id"`
	SerialName string   `json:"serialName"`
	Titles     []*Title `json:"titles" gorm:"many2many:titles_serials;"`
}
