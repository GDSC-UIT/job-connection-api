package models

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	ID    	string
	Name  	string
	Email 	string
	Photo 	string
}

type CompanyProfile struct {
	gorm.Model
	ID				string
	Name 			string
	Email 			string
	Photo 			string
	Address 		string
	WorkingTime 	string
	Description		string
}

type Skill struct {
	gorm.Model
	ID 		int
	Name 	string
}

type Job struct {
	gorm.Model
	CompanyProfile
	Title 		string
	Address		string
	Description string
	SkillIds	[]int
}

type ApplyRequest struct {
	gorm.Model
	UserProfile
	CompanyProfile
	CV			string
	Note 		string
	Closed		bool
	ClosedAt	time.Time
}