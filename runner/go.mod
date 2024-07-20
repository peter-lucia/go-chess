module github.com/peter-lucia/go-chess/runner

go 1.22.5

require github.com/peter-lucia/go-chess/engine v0.0.0-00010101000000-000000000000

require github.com/peter-lucia/go-chess/ui v0.0.0-00010101000000-000000000000

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/labstack/echo/v4 v4.11.3 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
)

// ensure we test local changes instead of those on github
replace github.com/peter-lucia/go-chess/engine => ../engine

replace github.com/peter-lucia/go-chess/ui => ../ui
