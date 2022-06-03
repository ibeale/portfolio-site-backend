package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
	return []byte(stamp), nil
}

type Post struct {
	gorm.Model
	ID      string    `json:"id" gorm:"primary_key"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}
