package domain

type CreateUserInput struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Password string `json:"password" validate:"required"`
	Avatar   string `json:"avatar"`
	Status   int64  `json:"status" validate:"required,oneof=0 1"`
}

type CreateUserOutput struct {
	Token string `json:"token"`
}
