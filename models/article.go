package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)

// Article is used by pop to map your articles database table to your go code.
type Article struct {
	ID          int          `json:"id" db:"id"`
	Guid        string       `json:"guid" db:"guid"`
	FeedURL     string       `json:"feed_url" db:"feed_url"`
	Title       string       `json:"title" db:"title"`
	Description string       `json:"description" db:"description"`
	Content     string       `json:"content" db:"content"`
	Link        string       `json:"link" db:"link"`
	ImageURL    nulls.String `json:"image_url" db:"image_url"`
	PublishedAt time.Time    `json:"published_at" db:"published_at"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (a Article) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Articles is not required by pop and may be deleted
type Articles []Article

// String is not required by pop and may be deleted
func (a Articles) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Article) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: a.Guid, Name: "Guid"},
		&validators.StringIsPresent{Field: a.FeedURL, Name: "FeedURL"},
		&validators.StringIsPresent{Field: a.Title, Name: "Title"},
		&validators.StringIsPresent{Field: a.Description, Name: "Description"},
		&validators.StringIsPresent{Field: a.Content, Name: "Content"},
		&validators.StringIsPresent{Field: a.Link, Name: "Link"},
		&validators.TimeIsPresent{Field: a.PublishedAt, Name: "PublishedAt"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Article) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Article) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
