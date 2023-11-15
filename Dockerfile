FROM hub.in.yottacloud.cn/iot/go-builder:1.17-alpine AS builder

ARG COMMIT_ID
ARG VERSION=""
ARG VCS_BRANCH=""
ARG GRPC_STUB_REVISION=""
ARG PROJECT_NAME=authory
ARG DOCKER_PROJECT_DIR=/build
ARG EXTRA_BUILD_ARGS=""
ARG GOCACHE=""
ARG GOMODCACHE

WORKDIR $DOCKER_PROJECT_DIR
COPY . $DOCKER_PROJECT_DIR

ENV GOPRIVATE=git.yottacloud.cn
ENV GOSUMDB=sum.golang.google.cn
ENV CGO_ENABLED=0

RUN git config --global url."git@git.yottacloud.cn:".insteadOf "https://git.yottacloud.cn/" \
    && go install git.yottacloud.cn/yottacloud/chaos/chaos@latest \
    && mkdir -p /output \
    && make dep \
    && make build -e GOCACHE=$GOCACHE \
    -e GOMODCACHE=$GOMODCACHE \
    -e COMMIT_ID=$COMMIT_ID -e OUTPUT_FILE=/output/authory \
    -e VERSION=$VERSION -e VCS_BRANCH=$VCS_BRANCH -e EXTRA_BUILD_ARGS=$EXTRA_BUILD_ARGS


## Uncomment lines when you write goa design
# 生成文档 .html
#FROM hub.in.yottacloud.cn/library/redoc-cli:latest as doc-render
#WORKDIR /build
#COPY --from=builder /build/ ./
#RUN make goa-doc -e OUTPUT_DIR=/output


FROM alpine:3.14
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk --no-cache --update add ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

ENV TZ=Asia/Shanghai

USER 1000

COPY --from=builder /output/authory /usr/local/bin/authory
COPY --chown=1000 config/config.sample.toml /app/
## Uncomment below lines when you have migrations
#COPY --chown=1000 migrations/ /app/migrations/
#COPY --from=migrate/migrate:v4.14.1 /usr/local/bin/migrate /usr/local/bin/migrate

## Uncomment lines when you write goa design
#COPY --from=doc-render --chown=1000 /output/apidoc.html /app/doc/apidoc.html



WORKDIR /app

EXPOSE 8080
CMD ["authory", "help"]
