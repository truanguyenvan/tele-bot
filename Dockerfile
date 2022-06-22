##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /tele-bot

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /tele-bot /tele-bot
COPY --from=build /app/config.json /


EXPOSE 5000

USER nonroot:nonroot

ENTRYPOINT ["/tele-bot"]