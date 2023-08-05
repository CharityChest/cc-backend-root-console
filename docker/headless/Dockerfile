ARG GOLANG_VERSION
ARG ALPINE_VERSION

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS executor

WORKDIR /app

COPY ./docker/headless/scripts/loop.sh ./

RUN ["chmod", "+x", "./loop.sh"]

RUN mkdir /code

WORKDIR /code

ENTRYPOINT ["/app/loop.sh"]