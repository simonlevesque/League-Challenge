FROM golang

COPY . /go/src/league/challenge
WORKDIR /go/src/league/challenge

RUN go mod init
RUN go mod tidy

# Make sure the script that start the code hot reloader can run from
# inside the container

RUN chmod +x dev.sh

CMD go run .
	
EXPOSE 8080