package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type SystemStatus int

const (
	DRAFT = iota
	PUBLISHED
	CLOSED
)

type Blog struct {
	Id           int          `json:"id" gorm:"primaryKey,unique"`
	CreatorId    int          `json:"creatorId"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	SystemStatus SystemStatus `json:"systemStatus"`
	ImageLinks   string       `json:"imageLinks"`
	CreationDate time.Time    `json:"creationDate"`
	BlogStatuses []BlogStatus `json:"blogStatuses" gorm:"foreignKey:BlogId;references:Id"`
	BlogRatings  []BlogRating `json:"blogRatings" gorm:"type:jsonb;"`
}

func (r BlogRating) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *BlogRating) Scan(value interface{}) error {
	if value == nil {
		*r = BlogRating{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, r)
}
