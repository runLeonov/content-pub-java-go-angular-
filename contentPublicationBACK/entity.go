package contentPublicationBACK

import (
	"time"
)

type Title struct {
	ID             uint         `json:"id"`
	TitleName      string       `json:"titleName"`
	Type           string       `json:"typeName"`
	OriginalAuthor string       `json:"originalAuthor"`
	CreationDate   time.Time    `json:"creationDate"`
	Description    string       `json:"description"`
	TitleImgBase64 string       `json:"titleImg"`
	Content        TitleContent `json:"content" gorm:"foreignKey:TitleID"`
	Categories     []Category   `json:"categories" gorm:"many2many:titles_categories;"`
	Tags           []Tag        `json:"tags" gorm:"many2many:titles_tags;"`
	Serials        []Serial     `json:"serials" gorm:"many2many:titles_serials;"`
}

type TitleContent struct {
	ID         uint    `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Title      *Title  `json:"-"`
	TitleID    uint    `json:"titleId"`
	LikesCount uint    `json:"likesCount" gorm:"default:0"`
	Likes      []Like  `json:"likes"`
	Views      uint    `json:"views"`
	Images     []Image `json:"images"`
}

type Image struct {
	ID             uint         `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	ImgBase64      string       `json:"img64"`
	TitleContentID uint         `json:"-"`
	TitleContent   TitleContent `json:"-" `
}

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

type AllContent struct {
	Types      []StaticType `json:"types"`
	Categories []Category   `json:"categories"`
	Tags       []Tag        `json:"tags"`
	Serials    []Serial     `json:"serials"`
}

type StaticType struct {
	ID       uint   `json:"id"`
	TypeName string `json:"typeName"`
}
