FROM golang:1.25.5-alpine as build
WORKDIR /api
COPY . .
RUN go mod download && go mod verify
RUN GOOS=linux GOARCH=amd64 go build -o api_executable
RUN ls

FROM alpine:latest
WORKDIR /
COPY --from=build /api/api_executable .
EXPOSE 5000

CMD ["./api_executable"]