package models

import (
	"time"

	pq "github.com/lib/pq"
	"gorm.io/gorm"
)

type AccountType struct {
	UserID string `gorm:"primarykey" json:"user_id"`
	Type   string `json:"type"`
}

type User struct {
	ID            string  `gorm:"primarykey" json:"id"`
	Name          string  `json:"name"`
	Email         string  `json:"email"  gorm:"unique"`
	Phone         string  `json:"phone"`
	Photo         string  `json:"photo"`
	Bio           string  `json:"bio"`
	Location      string  `json:"location"`
	FacebookURL   string  `json:"facebook_url"`
	LinkedinURL   string  `json:"linkedin_url"`
	GPA           float64 `json:"gpa"`
	NumberOfYears int     `json:"number_of_years"`
	Degree        string  `json:"degree"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Company struct {
	ID          string         `json:"id" gorm:"primarykey"`
	Name        string         `json:"name"`
	Email       string         `json:"email" gorm:"unique"`
	Photo       string         `json:"photo"`
	Address     string         `json:"address"`
	WorkingTime string         `json:"working_time"`
	Description string         `json:"description"`
	Approved    bool           `json:"approved"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Skill struct {
	gorm.Model
	ID   int
	Name string
}

type Job struct {
	gorm.Model
	CompanyID   string
	Title       string
	Address     string
	Description string
	SkillIds    pq.Int64Array `gorm:"type:integer[]"`
}
type ApplyRequest struct {
	gorm.Model
	UserID    string
	CompanyID string
	CV        string
	Note      string
	Closed    bool
	ClosedAt  time.Time
}

type Experience struct {
	gorm.Model
	UserID      string
	CompanyID   string
	JobTitle    string
	Description string
	SkillIds    []int `gorm:"type:integer[]"`
	From        time.Time
	To          time.Time
}
