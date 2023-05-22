package goweb

type StdResponseData[E any] struct {
	Data E    `json:"data,omitempty"`
	Msg  string `json:"msg,omitempty"`
}
