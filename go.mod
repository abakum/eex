module github/abakum/eex

go 1.21.4

require (
	github.com/abakum/embed-encrypt v0.0.0-20240309090251-5db710414591 // indirect
	internal/tool v0.0.0-00010101000000-000000000000 // indirect
)

replace internal/tool => ./internal/tool
//replace github.com/abakum/embed-encrypt => ../embed-encrypt
