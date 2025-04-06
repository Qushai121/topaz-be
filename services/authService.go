package services

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Qushai121/topaz-be/dto"
	authdto "github.com/Qushai121/topaz-be/dto/authDto"
	"github.com/Qushai121/topaz-be/entities"
	"github.com/Qushai121/topaz-be/models"
	"github.com/Qushai121/topaz-be/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IAuthService interface {
	SignIn(requestBody *authdto.SignInRequestBodyDto) (*dto.SuccessDto[*authdto.SignInResponseBodyDto], *dto.ErrorDto[any])
	SignUp(requestBody *authdto.SignUpRequestBodyDto) (*dto.SuccessDto[any], *dto.ErrorDto[any])
	PostNewAccessToken(requestBody *authdto.PostNewAccessTokenRequestBodyDto) (*dto.SuccessDto[*authdto.PostNewAccessTokenResponseBodyDto], *dto.ErrorDto[any])
}

type authService struct {
	dbTopaz *gorm.DB
}

func NewAuthService(dbTopaz *gorm.DB) IAuthService {
	return &authService{
		dbTopaz: dbTopaz,
	}
}

func (a *authService) SignIn(requestBody *authdto.SignInRequestBodyDto) (*dto.SuccessDto[*authdto.SignInResponseBodyDto], *dto.ErrorDto[any]) {
	user := models.User{}

	res := a.dbTopaz.Where(models.User{Email: requestBody.Email}).First(&user)

	if res.Error != nil {
		return nil, dto.NewErrorDto[any]("User not found", http.StatusNotFound, nil)
	}

	errComparePass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))

	if errComparePass != nil {
		return nil, dto.InternalServerError()
	}

	payload := entities.UserToken{
		UserId: user.ID,
	}

	accessToken, errAccessToken := utils.GenerateToken(payload, time.Hour*2, utils.AccessTokenKey)

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	refreshToken, errRefreshToken := utils.GenerateToken(payload, time.Hour*(24*7), utils.RefreshTokenKey)

	if errRefreshToken != nil {
		return nil, errRefreshToken
	}

	responseData := authdto.SignInResponseBodyDto{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}

	return dto.NewSuccessDto("Sign in successfully executed", int(http.StatusCreated), &responseData), nil
}

func (a *authService) SignUp(requestBody *authdto.SignUpRequestBodyDto) (*dto.SuccessDto[any], *dto.ErrorDto[any]) {
	hashedPass, errHash := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)

	if errHash != nil {
		log.Printf("Error : %v", errHash.Error())
		return nil, dto.NewErrorDto[any]("Error : when try to hash password", int(http.StatusInternalServerError), nil)
	}

	user := models.User{
		FirstName: requestBody.FirstName,
		LastName: sql.NullString{
			String: requestBody.LastName,
			Valid:  requestBody.LastName != "",
		},
		Password: string(hashedPass),
		Email:    requestBody.Email,
	}

	res := a.dbTopaz.FirstOrCreate(&user, models.User{
		Email: requestBody.Email,
	})

	if res.Error != nil {
		return nil, dto.NewErrorDto[any]("Database Error : when try to create new user", http.StatusInternalServerError, nil)
	}

	if res.RowsAffected == 0 {
		return nil, dto.NewErrorDto[any]("Email Have been used", http.StatusBadRequest, nil)
	}

	return dto.NewSuccessDto[any]("Sign up successfully executed", int(http.StatusCreated), nil), nil
}

// Service To get new access token using refresh token from request body
func (a *authService) PostNewAccessToken(requestBody *authdto.PostNewAccessTokenRequestBodyDto) (*dto.SuccessDto[*authdto.PostNewAccessTokenResponseBodyDto], *dto.ErrorDto[any]) {
	encodeRefreshToken, errEncodeRefreshToken := utils.EncodeToken[entities.UserToken](requestBody.RefreshToken, utils.RefreshTokenKey)

	if errEncodeRefreshToken != nil {
		return nil, errEncodeRefreshToken
	}

	accessToken, errAccessToken := utils.GenerateToken(encodeRefreshToken.Data, time.Hour*2, utils.AccessTokenKey)

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	responseData := authdto.PostNewAccessTokenResponseBodyDto{
		AccessToken: *accessToken,
	}

	return dto.NewSuccessDto("Refresh Access Token successfully executed", http.StatusCreated, &responseData), nil
}
