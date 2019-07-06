package models

import (
	"time"
)

type Job struct {
	Title    *string    `json:"title,omitempty"`
	Company  *string    `json:"company,omitempty"`
	Location *string    `json:"location,omitempty"`
	Salary   *string    `json:"salary,omitempty"`
	Posted   *time.Time `json:"posted,omitempty"`
	Link     *string    `json:"link,omitempty"`
	Keywords []string   `json:"keywords,omitempty"`
	Source   *string    `json:"source,omitempty"`
}
