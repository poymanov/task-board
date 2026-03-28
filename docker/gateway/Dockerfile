FROM golang:1.25.5

ARG GROUP=1000
ARG USER=1000

WORKDIR /app

COPY gateway/go.mod gateway/go.sum ./gateway/
COPY platform/go.mod platform/go.sum ./platform/
COPY shared/go.mod shared/go.sum ./shared/

WORKDIR /app/gateway
RUN go mod download

WORKDIR /app

RUN groupadd -g $GROUP user
RUN useradd -u $USER -ms /bin/bash -g user user

USER user
