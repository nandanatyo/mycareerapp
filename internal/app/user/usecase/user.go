package usecase

import (
	"errors"
	"mycareerapp/internal/app/user/repository"
	"mycareerapp/internal/domain/dto"
	"mycareerapp/internal/domain/entity"
	"mycareerapp/internal/infra/jwt"
	"mycareerapp/pkg/util"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseItf interface {
	Register(dto.Register) error
	Login(dto.Login) (string, error)
}

type UserUsecase struct {
	userRepo repository.UserMySQLItf
	jwt      jwt.JWTI
}

func NewUserUsecase(userRepo repository.UserMySQLItf, jwt jwt.JWTI) UserUsecaseItf {
	return &UserUsecase{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

func (u *UserUsecase) Register(register dto.Register) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	gender, err := util.CheckGender(register.Gender)
	if err != nil {
		return err
	}

	birthday, err := time.Parse("02/01/2006", register.Birthday)
	if err != nil {
		return err
	}

	user := entity.User{
		ID:                        uuid.New(),
		Name:                      register.Name,
		Email:                     register.Email,
		Password:                  string(hashedPassword),
		PhotoProfile:              register.PhotoProfile,
		Gender:                    gender,
		Birthday:                  birthday,
		EducationalLevel:          register.EducationalLevel,
		Institution:               register.Institution,
		Departmen:                 register.Departmen,
		StartEducation:            register.StartEducation,
		StillInSchool:             register.StillInSchool,
		Graduate:                  register.Graduate,
		WorkExperience:            register.WorkExperience,
		WorkPlace:                 register.WorkPlace,
		Position:                  register.Position,
		StartPeriod:               register.StartPeriod,
		EndPeriod:                 register.EndPeriod,
		SelfDescription:           register.SelfDescription,
		WorkExperienceDescription: register.WorkExperienceDescription,
		WorkArea:                  register.WorkArea,
		WorkRoles:                 register.WorkRoles,
		CV:                        register.CV,
		IsAdmin:                   false,
	}

	err = u.userRepo.Create(&user)

	return err
}

func (u *UserUsecase) Login(login dto.Login) (string, error) {
	var user entity.User

	err := u.userRepo.Get(&user, dto.UserParam{Email: login.Email})
	if err != nil {
		return "", errors.New("email atau password invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", errors.New("email atau password invalid")
	}

	token, err := u.jwt.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		return "", err
	}

	return token, nil
}
