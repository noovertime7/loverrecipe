package domain

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
