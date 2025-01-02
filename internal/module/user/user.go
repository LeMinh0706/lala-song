package user

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type Register struct {
	Username string `json:"username" validate:"required,min=6,max=16"`
	Password string `json:"password" validate:"required,min=6,max=16"`
	Fullname string `json:"fullname" validate:"required,min=6,max=16"`
	Gender   int32  `json:"gender" validate:"required,min=0,max=1"`
}
