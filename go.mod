module goweb

go 1.20

replace goutils => ../efast-goutils

require (
	github.com/gorilla/mux v1.8.0
	github.com/rs/zerolog v1.29.1
	goutils v0.0.0-00010101000000-000000000000
)

require (
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/sys v0.1.0 // indirect
)
