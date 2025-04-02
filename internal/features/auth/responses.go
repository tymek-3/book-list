package auth

type jwtTokenResponse struct {
	Token string `json:"token" binding:"required"`
}

type sliceResponse[T any] struct {
	Data []T `json:"data"`
}
