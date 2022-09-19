package httputils

type EmptyResponse struct{}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
