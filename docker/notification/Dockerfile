FROM golang:1.25.5

ARG GROUP=1000
ARG USER=1000

WORKDIR /app

COPY platform/go.mod platform/go.sum ./platform/
COPY shared/go.mod shared/go.sum ./shared/
COPY notification/go.mod notification/go.sum ./notification/

WORKDIR /app/notification
RUN go mod download

WORKDIR /app

RUN groupadd -g $GROUP user
RUN useradd -u $USER -ms /bin/bash -g user user

USER user
