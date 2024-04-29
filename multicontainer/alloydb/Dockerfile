ARG GO_VERSION=1.22
ARG APP=todos

FROM golang:${GO_VERSION}-alpine AS build

WORKDIR /${APP}

COPY main.go handlers.go go.mod go.sum Makefile ./
COPY public public
COPY views views

RUN apk --no-cache add --update make libx11-dev gcc libc-dev curl

# Intentionally un packed
RUN make build

FROM gcr.io/distroless/static AS final

LABEL maintainer="Julien BREUX <julien.breux@gmail.com>"
USER nonroot:nonroot

COPY --from=build --chown=nonroot:nonroot /${APP}/bin/app /app

ENTRYPOINT ["/app"]