package dto

import "github.com/google/uuid"

type Register struct {
	Name                      string `json:"name" validate:"required"`
	Email                     string `json:"email" validate:"required,email"`
	Password                  string `json:"password" validate:"required,min=6"`
	PhotoProfile              string `json:"photo_profile"`
	Gender                    string `json:"gender" validate:"required"`
	Birthday                  string `json:"birthday" validate:"required"`
	EducationalLevel          string `json:"educational_level" validate:"required"`
	Institution               string `json:"institution" validate:"required"`
	Departmen                 string `json:"departmen" validate:"required"`
	StartEducation            string `json:"start_education" validate:"required"`
	StillInSchool             bool   `json:"still_in_school"`
	Graduate                  string `json:"graduate"`
	WorkExperience            string `json:"work_experience"`
	WorkPlace                 string `json:"work_place"`
	Position                  string `json:"position"`
	StartPeriod               string `json:"start_period"`
	EndPeriod                 string `json:"end_period"`
	SelfDescription           string `json:"self_description"`
	WorkExperienceDescription string `json:"work_experience_description" `
	WorkArea                  string `json:"work_area" validate:"required"`
	WorkRoles                 string `json:"work_roles" validate:"required"`
	CV                        string `json:"cv" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserParam struct {
	ID    uuid.UUID
	Email string
}
