FROM golang

WORKDIR /app

COPY . .

RUN go get download

RUN go build -o app.exe .

EXPOSE 8080

CMD [ "./app.exe" ]