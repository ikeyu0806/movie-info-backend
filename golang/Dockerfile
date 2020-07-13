FROM golang:latest

RUN mkdir -p -m 775 /workspace/golang_backend

WORKDIR /workspace/golang_backend

ADD ./ /workspace/golang_backend

CMD ["go", "build"]
CMD ["go", "run", "./cmd/demo/main.go"]
