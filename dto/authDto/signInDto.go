package authdto

type SignInRequestBodyDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponseBodyDto struct{
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
