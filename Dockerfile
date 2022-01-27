FROM golang AS builder

WORKDIR /opt
 
COPY go.mod go.sum ./
RUN go mod download && go mod verify && go mod tidy

COPY . .

RUN go build -o application .


FROM ubuntu AS production

RUN apt-get update && \
    apt-get -y install pdftk 

WORKDIR /opt
RUN mkdir data
ADD f8949.pdf .

COPY --from=builder /opt/application ./
VOLUME [ "/opt/data" ]
CMD ["./application"]