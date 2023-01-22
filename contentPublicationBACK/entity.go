package contentPublicationBACK

import (
	"time"
)

type Title struct {
	ID             uint        `json:"id"`
	TitleName      string      `json:"titleName"`
	OriginalAuthor string      `json:"originalAuthor"`
	CreationDate   time.Time   `json:"creationDate"`
	Description    string      `json:"description"`
	TitleImgBase64 string      `json:"titleImg"`
	Categories     []*Category `json:"categories" gorm:"many2many:titles_categories;"`
	Tags           []*Tag      `json:"tags" gorm:"many2many:titles_tags;"`
	Serials        []*Serial   `json:"serials" gorm:"many2many:titles_serials;"`
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
