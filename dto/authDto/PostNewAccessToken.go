package authdto

type PostNewAccessTokenRequestBodyDto struct{
	RefreshToken    string `json:"refreshToken" validate:"required"`
}

type PostNewAccessTokenResponseBodyDto struct{
	AccessToken    string `json:"accessToken" validate:"required"`
}
