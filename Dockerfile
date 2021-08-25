FROM golang:1.15

RUN mkdir -p $GOPATH/src/github.com/ncostamagna/axul_event
WORKDIR $GOPATH/src/github.com/ncostamagna/axul_event
COPY . .
RUN ls

RUN go get -d -v ./... 
RUN go install -v ./...
EXPOSE 8081

CMD ["axul_event"]