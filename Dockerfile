FROM golang:1.19.0-buster as build
WORKDIR /app
COPY . .
RUN go build -v

FROM ubuntu:22.04
WORKDIR /app
EXPOSE 80

COPY --from=build /app/erupe-ce .
COPY www ./www
COPY savedata ./savedata
COPY bin ./
COPY RoadShopItems.csv ./
RUN chmod +x erupe-ce
CMD ["/app/erupe-ce"]
