# Build
FROM golang:1.20-alpine AS Build
WORKDIR /src
COPY . .
# RUN go mod download
EXPOSE 8080
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o=debug-container main.go

# deploy

FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app
RUN chmod 777 /app
COPY --from=Build /src/debug-container /app/debug-container

RUN addgroup -S me
RUN adduser -S me -G me
RUN chown -R me:me /app
USER me:me

EXPOSE 8080
CMD ["./debug-container"]
