ARG GOLANG_VERSION
ARG ALPINE_VERSION

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS builder

WORKDIR /app

COPY . ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

FROM alpine:${ALPINE_VERSION} AS runner

ARG APP_GROUP=executer
ARG APP_GROUP_ID=1100
ARG APP_USER=executer
ARG APP_USER_ID=1100

RUN addgroup -g "${APP_GROUP_ID}" -S "${APP_GROUP}"
RUN adduser -u "${APP_USER_ID}" -S -D -H -G "${APP_GROUP}" "${APP_USER}"
RUN adduser "${APP_USER}" tty

COPY --from=builder --chown=${APP_USER}:${APP_GROUP} /main /main

EXPOSE 8080

USER ${APP_USER}

CMD ["/main"]