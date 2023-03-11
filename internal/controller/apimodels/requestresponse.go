package apimodels

type SignInRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email"`
}
type NewWalletRequest struct {
	UserID  int `json:"user_id"`
	Balance int `json:"balance"`
}
type GetBalanceRequest struct {
	UserId int `json:"user_id"`
}
type GetBalanceResponse struct {
	Balance int `json:"balance"`
}
type AddMoneyRequest struct {
	UserID int `json:"user_id"`
	Amount int `json:"amount"`
}
type AddMoneyResponse struct {
	ReferenceID int `json:"reference_id"`
}
type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type RemoveUserRequest struct {
	UserName string `json:"user_name"`
}
