###########
## BUILD ##
###########

# specify a base image from which to to build the application
FROM golang:1.19-alpine AS build

# set the working directory
WORKDIR /app

# copy files required to run the app into the working directory
COPY app .

# download modules into cache
RUN go mod download

# build the app and output to gin-api directory, creating a static binary we can copy across
# without worrying about external dependencies
RUN CGO_ENABLED=0 go build -o /gin-api -ldflags="-w -s"

############
## DEPLOY ##
############

# use the official scratch image from https://hub.docker.com/_/scratch
FROM scratch

WORKDIR /

# copy the built binary into scratch environment
COPY --from=build /gin-api /gin-api

CMD [ "/gin-api" ]