FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/restaurant




# FROM gcr.io/distroless/base-debian11
FROM alpine:latest

WORKDIR /

COPY --from=build /app/bin/restaurant /usr/local/bin/restaurant
COPY --from=build /app/migrations /migrations
COPY --from=build /app/restaurantData /restaurantData

# EXPOSE 5000

CMD ["sh", "-c", "restaurant migrate up --config /app/prod.env && restaurant start --config /app/prod.env"]
