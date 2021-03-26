FROM golang:latest

ENV DB_HOST=movie_mysql:3306
ENV DB_USER=root
ENV DB_PASSWORD=pass
ENV DB_NAME=movie_info
ENV DATA_SOURCE=root:pass@tcp(mysql:3306)/movie_info?charset=utf8&parseTime=True&loc=Local

RUN mkdir -p -m 775 /workspace/golang_backend

WORKDIR /workspace/golang_backend

ADD ./ /workspace/golang_backend

RUN go get -v github.com/rubenv/sql-migrate/...

RUN ["go", "build"]
CMD ["go", "run", "./main.go"]
