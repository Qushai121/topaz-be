package services

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Qushai121/topaz-be/dto"
	authdto "github.com/Qushai121/topaz-be/dto/authDto"
	"github.com/Qushai121/topaz-be/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IAuthService interface {
	SignIn(requestBody *authdto.SignInRequestBodyDto) (*dto.SuccessDto[any], *dto.ErrorDto[any])
	SignUp(requestBody *authdto.SignUpRequestBodyDto) (*dto.SuccessDto[*authdto.SignUpResponseBodyDto], *dto.ErrorDto[any])
}

type authService struct {
	dbTopaz *gorm.DB
}

func NewAuthService(dbTopaz *gorm.DB) IAuthService {
	return &authService{
		dbTopaz: dbTopaz,
	}
}

func (a *authService) SignIn(requestBody *authdto.SignInRequestBodyDto) (*dto.SuccessDto[any], *dto.ErrorDto[any]) {
	return nil, nil
	// hashedPass, errHash := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	// if errHash != nil {
	// 	return nil, dto.InternalServerError()
	// }

	// bcrypt.CompareHashAndPassword()

}

func (a *authService) SignUp(requestBody *authdto.SignUpRequestBodyDto) (*dto.SuccessDto[*authdto.SignUpResponseBodyDto], *dto.ErrorDto[any]) {
	hashedPass, errHash := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)

	if errHash != nil {
		log.Printf("Error : %v",errHash.Error())
		return nil, dto.NewErrorDto[any]("Error : when try to hash password", int(http.StatusInternalServerError), nil)
	}

	user := models.User{
		FirstName: requestBody.FirstName,
		LastName: sql.NullString{
			String: requestBody.LastName,
			Valid:  requestBody.LastName != "",
		},
		Password:      string(hashedPass),
		Email:         requestBody.Email,
		RememberToken: "tokenasdasdasdasd",
	}

	res := a.dbTopaz.Create(&user)

	if res.Error != nil {
		log.Printf("Error : %v",errHash.Error())
		return nil, dto.NewErrorDto[any]("Database Error : when try to create new user", int(http.StatusInternalServerError), nil)
	}

	responseData := authdto.SignUpResponseBodyDto{
		AccessToken:  "accessTokenasdsadsdsdsa",
		RefreshToken: "refreshTokenasdasdvberthyrer4",
	}

	return dto.NewSuccessDto("Sign up successfully executed", int(http.StatusCreated), &responseData), nil
}
