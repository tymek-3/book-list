package auth

type jwtTokenResponse struct {
	Token string `json:"token" binding:"required"`
}
