package goweb

type StandardResponse struct {
	Data any    `json:"data,omitempty"`
	Msg  string `json:"msg,omitempty"`
}
