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
	ID            string        `gorm:"primarykey" json:"id"`
	Name          string        `json:"name"`
	Email         string        `json:"email"  gorm:"unique"`
	Phone         string        `json:"phone"`
	Photo         string        `json:"photo"`
	Bio           string        `json:"bio"`
	Location      string        `json:"location"`
	FacebookURL   string        `json:"facebook_url"`
	LinkedinURL   string        `json:"linkedin_url"`
	GPA           float64       `json:"gpa"`
	NumberOfYears int           `json:"number_of_years"`
	Degree        string        `json:"degree"`
	HardSkillIds  pq.Int64Array `json:"hard_skill_ids" gorm:"type:integer[]"`
	SoftSkillIds  pq.Int64Array `json:"soft_skill_ids" gorm:"type:integer[]"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Company struct {
	ID          string `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Email       string `json:"email" gorm:"unique"`
	Photo       string `json:"photo"`
	Address     string `json:"address"`
	WorkingTime string `json:"working_time"`
	Description string `json:"description"`
	Approved    bool   `json:"approved"`
	Jobs        []Job  `json:"jobs"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Skill struct {
	ID        int            `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Job struct {
	ID          uint          `gorm:"primarykey" json:"id"`
	Company     Company       `json:"company"`
	CompanyID   string        `json:"company_id"`
	Title       string        `json:"title"`
	Address     string        `json:"address"`
	Description string        `json:"description"`
	SkillIds    pq.Int64Array `gorm:"type:integer[]" json:"skill_ids"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type ApplyRequest struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	User      *User     `json:"user"`
	UserID    string    `json:"user_id"`
	Job       *Job      `json:"job"`
	JobID     uint      `json:"job_id"`
	CompanyID string    `json:"company_id"`
	CV        string    `json:"cv"`
	Note      string    `json:"note"`
	Closed    bool      `json:"closed"`
	ClosedAt  time.Time `json:"closed_at"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Experience struct {
	gorm.Model
	UserID      string        `json:"user_id"`
	Company     Company       `json:"company"`
	CompanyID   string        `json:"company_id"`
	JobTitle    string        `json:"job_title"`
	Description string        `json:"description"`
	SkillIds    pq.Int64Array `gorm:"type:integer[]" json:"skill_ids"`
	From        *time.Time    `json:"from"`
	To          *time.Time    `json:"to"`
}
