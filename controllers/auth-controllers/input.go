package authControllers

type SessionInput struct {
	Cpf      string `json:"cpf" validate:"required,lowercase,len=11"`
	Password string `json:"password" validate:"required"`
}
