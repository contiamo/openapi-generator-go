FROM --platform=${TARGETPLATFORM:-linux/amd64} gcr.io/distroless/static:nonroot

ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

COPY dist/openapi-generator-go_${TARGETOS}_${TARGETARCH}/openapi-generator-go /

ENTRYPOINT ["/openapi-generator-go"]