FROM golang:1.9-alpine

ENV GOPATH="/home/app"
ENV PATH=$PATH:$GOPATH/bin

WORKDIR "$GOPATH/src/github.com/dmitryk-dk/raspberry"

RUN mkdir $GOPATH/bin
RUN apk update && \
    apk add git curl yarn && \
    rm -rf /var/cache/apk/*

COPY . .

RUN curl https://glide.sh/get | sh && \
    glide i

RUN rm -rf node_modules && \
    yarn install --non-interactive --silent && \
    yarn build

EXPOSE 3000

CMD ["go", "run", "main.go"]
