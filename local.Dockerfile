FROM golang:1.19.3

RUN apt update && apt upgrade -y && \
	apt install -y git \
	make openssh-client

# all the code lives here. We gonna mount the volume here
WORKDIR /products

# copy the go.mod and go.sum files then download the dependencies
ADD go.mod go.sum /products/
RUN go mod download

# install the air tool
RUN go install github.com/cosmtrek/air@latest

# install the dlv debugger
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# will run from an entrypoint.sh file
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]
