package authdto

type SignUpRequestBodyDto struct {
	FirstName    string `json:"firstName" validate:"required,alpha,min=3,max=50"`
	LastName     string `json:"lastName" validate:"required,alpha,min=3,max=50"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=6"`
	IsRememberMe *bool  `json:"isRememberMe" validate:"boolean"`
}
