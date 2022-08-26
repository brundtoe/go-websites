module github.com/brundtoe/websites

go 1.16

replace brundtoe/bookstore => ../go-bookstore

require (
	brundtoe/bookstore v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.8.1
)
