FROM golang:1.18-alpine 

RUN  mkdir /app 

COPY  . /app 

WORKDIR /app 

RUN CGO_ENABLED=0 go build -o front-end ./cmd/web

CMD [ "/app/front-end" ]
