FROM golang:1.25.5

ARG GROUP=1000
ARG USER=1000

WORKDIR /app

COPY platform/go.mod platform/go.sum ./platform/
COPY shared/go.mod shared/go.sum ./shared/
COPY auth/go.mod auth/go.sum ./auth/

WORKDIR /app/auth
RUN go mod download

WORKDIR /app

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN groupadd -g $GROUP user
RUN useradd -u $USER -ms /bin/bash -g user user

USER user
