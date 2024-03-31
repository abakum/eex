module github.com/abakum/eex

go 1.21.4

replace internal/tool => ./internal/tool

replace public/tool => ./public/tool

//replace github.com/abakum/embed-encrypt => ../embed-encrypt

require (
	github.com/abakum/embed-encrypt v0.0.0-20240330115809-059354cfa29a
	internal/tool v0.0.0-00010101000000-000000000000
	public/tool v0.0.0-00010101000000-000000000000
)
