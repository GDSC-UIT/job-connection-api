package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string  `gorm:"primarykey" json:"id"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
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
	gorm.Model
	ID          string
	Name        string
	Email       string
	Photo       string
	Address     string
	WorkingTime string
	Description string
}

type Skill struct {
	gorm.Model
	ID   int
	Name string
}

type Job struct {
	gorm.Model
	Company
	Title       string
	Address     string
	Description string
	SkillIds    []int
}

type ApplyRequest struct {
	gorm.Model
	User
	Company
	CV       string
	Note     string
	Closed   bool
	ClosedAt time.Time
}
