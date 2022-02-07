FROM golang:1.16-alpine AS build
WORKDIR ./app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /content

##
## Deploy
##
FROM scratch

WORKDIR /

COPY --from=build /content /content
ENV MONGO_URI=${MONGO_URI}
ENV PORT=${PORT}
ENV DATABASE=${DATABASE}
ENV COLLECTION=${COLLECTION}

EXPOSE ${PORT}
CMD [ "/content" ]