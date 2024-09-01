# Stage 1: Complite the binary in a containerized Golang environment
#
FROM golang:1.23 as build

# Copy the sous files from the host
COPY . /src

# Set the working directory to the same place we copied the code
WORKDIR /src

# Bilding the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o kvs

# Stage 2: Build the Key-Value Store image proper
#
# Use a "scratch" image, which contains no distribution files
FROM scratch

# Tell Docker we'll be using port 8080
EXPOSE 8080

# Tell Docker to execute this command on a "docker run"
CMD ["/kvs"]