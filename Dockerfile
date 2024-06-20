# syntax=docker/dockerfile:1
ARG GO_VERSION=1.22.4
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine AS build

WORKDIR /src

COPY . .

RUN go build -o /bin/lci .

FROM alpine

COPY --from=build /bin/lci /bin/lci

ENTRYPOINT [ "/bin/lci" ]
