module github.com/ikeyu0806/movie-info-backend

go 1.14

require (
	bitbucket.org/liamstask/goose v0.0.0-20150115234039-8488cc47d90c // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/csrf v1.7.0
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/ikeyu0806/movie-info-backend/db v0.0.0-20200716141722-25051f327dd1
	github.com/jinzhu/gorm v1.9.14
	github.com/kylelemons/go-gypsy v0.0.0-20160905020020-08cad365cd28 // indirect
	github.com/lib/pq v1.7.0
	github.com/ziutek/mymysql v1.5.4 // indirect
	ikeyu0806/db v0.0.0-00010101000000-000000000000
)

replace ikeyu0806/db => github.com/ikeyu0806/movie-info-backend/db v0.0.0-20200715072438-137b842e20f0

replace movie-info-backend => github.com/ikeyu0806/movie-info-backend v0.0.0-20200716141722-25051f327dd1
