module movie-info-backend

go 1.14

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/csrf v1.7.0
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/lib/pq v1.7.0
	local.packages/server v0.0.0-00010101000000-000000000000
)

replace local.packages/server => ./server
