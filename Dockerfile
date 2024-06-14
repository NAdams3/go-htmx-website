FROM golang:1.22 as base
FROM base as dev
RUN curl https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
WORKDIR /opt/app
CMD ["air"]

