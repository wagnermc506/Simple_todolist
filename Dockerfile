FROM golang

WORKDIR /app

COPY ./src .

WORKDIR /app/todolist

RUN go build -o /

CMD ["/todolist"]
