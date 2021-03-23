FROM golang:latest

ENV DB_HOST=movie_mysql:3306
ENV DB_USER=root
ENV DB_PASSWORD=pass
ENV DB_NAME=movie_info

RUN mkdir -p -m 775 /workspace/golang_backend
RUN go get -u bitbucket.org/liamstask/goose/cmd/goose

WORKDIR /workspace/golang_backend

ADD ./ /workspace/golang_backend

RUN ["go", "build"]
CMD ["go", "run", "./main.go"]
