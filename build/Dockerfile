# syntax=docker/dockerfile:1

# Build stage: compile the operator for the requested target platform.
# BUILDPLATFORM keeps the compiler running natively (fast) while cross-compiling to
# TARGETOS/TARGETARCH, so `docker buildx --platform linux/arm64,linux/amd64` produces a
# correct binary per arch instead of copying a single prebuilt one.
FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /workspace

# Cache module downloads before copying the full tree.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath -tags=v1 -o /out/kubedge-base-operator ./cmd/...

# Runtime stage.
FROM alpine:3.20

ENV OPERATOR=/usr/local/bin/kubedge-base-operator \
    USER_UID=1001 \
    USER_NAME=kubedge-base-operator

# install operator binary
COPY --from=builder /out/kubedge-base-operator ${OPERATOR}

COPY build/bin /usr/local/bin

RUN  /usr/local/bin/user_setup

COPY build/test-pods /opt/kubedge-operators/test-pods

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
