package login

type LoginDTO struct {
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required"`
}
