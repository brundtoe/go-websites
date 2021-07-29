module example.com/websites

go 1.16

replace example.com/bookstore => ../bookstore

require (
	example.com/bookstore v0.0.0-00010101000000-000000000000 // indirect
	github.com/gin-gonic/gin v1.7.2 // indirect
)
