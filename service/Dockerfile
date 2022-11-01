FROM golang:1.16.3-buster

WORKDIR /app/src
RUN go get github.com/cespare/reflex 
RUN echo "cd /app/src && reflex -r '\.go\$' -s -- sh -c 'go mod tidy && go run main.go'" >> /run.sh
RUN sed -i 's/\r//' /run.sh
ENTRYPOINT ["sh", "/run.sh"]