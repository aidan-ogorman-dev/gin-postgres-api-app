# specify a base image from which to to build the application
FROM golang:1.19-alpine

# set the working directory
WORKDIR /app

# copy files required to run the app into the working directory
COPY app .

# download modules into cache
RUN go mod download

# build the app and output to gin-api directory
RUN go build -o /gin-api

CMD [ "/gin-api" ]
