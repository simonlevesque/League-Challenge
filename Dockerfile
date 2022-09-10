FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/league/challenge
WORKDIR /go/src/league/challenge

RUN go mod init
RUN go mod tidy

CMD go run .
	
EXPOSE 8080