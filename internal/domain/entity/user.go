package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                        uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
	Name                      string    `json:"name" gorm:"type:varchar(255);not null"`
	Email                     string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password                  string    `json:"password" gorm:"type:varchar(255);not null"`
	PhotoProfile              string    `json:"photo_profile" gorm:"type:text"`
	Gender                    string    `json:"gender" gorm:"type:text;not null"`
	Birthday                  string    `json:"birthday" gorm:"type:text;not null"`
	EducationalLevel          string    `json:"educational_level" gorm:"type:text;not null"`
	Institution               string    `json:"institution" gorm:"type:text;not null"`
	Departmen                 string    `json:"departmen" gorm:"type:text;not null"`
	StartEducation            string    `json:"start_education" gorm:"type:text;not null"`
	StillInSchool             bool      `json:"still_in_school" gorm:"type:boolean"`
	Graduate                  string    `json:"graduate" gorm:"type:text;not null"`
	WorkExperience            string    `json:"work_experience" gorm:"type:text"`
	WorkPlace                 string    `json:"work_place" gorm:"type:text;not null"`
	Position                  string    `json:"position" gorm:"type:text;not null"`
	StartPeriod               string    `json:"start_period" gorm:"type:text;not null"`
	EndPeriod                 string    `json:"end_period" gorm:"type:text;not null"`
	SelfDescription           string    `json:"self_description" gorm:"type:text"`
	WorkExperienceDescription string    `json:"work_experience_description" gorm:"type:text"`
	WorkArea                  string    `json:"work_area" gorm:"type:text"`
	WorkRoles                 string    `json:"work_roles" gorm:"type:text"`
	CV                        string    `json:"cv" gorm:"type:text"`
	IsAdmin                   bool      `json:"is_admin" gorm:"type:boolean"`
	CreatedAt                 time.Time    `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
}
