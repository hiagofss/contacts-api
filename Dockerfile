FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o /server .

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /server /server

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]