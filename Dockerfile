FROM golang:latest

RUN mkdir -p -m 775 /workspace/golang_backend
RUN go get -u bitbucket.org/liamstask/goose/cmd/goose

WORKDIR /workspace/golang_backend

ADD ./ /workspace/golang_backend

CMD ["go", "build"]
CMD ["go", "run", "./main.go"]
