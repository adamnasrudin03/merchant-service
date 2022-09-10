package dto

//LoginDTO is a model that used by client when POST from /login url
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
