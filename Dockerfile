FROM golang

RUN apt-get update
RUN apt-get install lsof

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

WORKDIR /app

EXPOSE 8080