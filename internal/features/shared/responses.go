package shared

type SliceResponse[T any] struct {
	Data []T `json:"data"`
}
