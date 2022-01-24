FROM ubuntu_go_pdftk
RUN mkdir -p /go/pdf
WORKDIR /go/pdf
ADD go.mod .
ADD go.sum .
ADD main.go .
ADD f8949.pdf .
RUN go mod download && go mod verify
RUN go build -o /go/pdf/app
VOLUME [ "/go/pdf" ]
CMD ["./app"]