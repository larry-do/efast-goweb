package goweb

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
)

func (resp Response) RespondJson(statusCode int, data any) {
	resp.ContentType("application/json").Code(statusCode).Json(data)
	log.Debug().Any("body", data).Int("code", statusCode).Msg("")
}

func (resp Response) RespondPlainText(statusCode int, data string) {
	resp.ContentType("text/plain").Code(statusCode).PlainText(data)
	log.Debug().Any("body", data).Int("code", statusCode).Msg("")
}

func (resp Response) ContentType(contentType string) Response {
	resp.Header().Set("Content-Type", contentType)
	return resp
}

func (resp Response) Json(object any) {
	var err = json.NewEncoder(resp).Encode(object)
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}
}

func (resp Response) PlainText(str string) {
	_, err := resp.Write([]byte(str))
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}
}

func (resp Response) Code(statusCode int) Response {
	resp.WriteHeader(statusCode)
	return resp
}

func (resp Response) HtmlView(viewTemplate string, model any) (Response, error) {
	var tmpl = template.Must(template.ParseFiles(viewTemplate))
	var err = tmpl.Execute(resp, model)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type Response struct {
	http.ResponseWriter
}
