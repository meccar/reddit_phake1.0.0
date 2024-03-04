# # Use the official Golang image as the base image
# FROM golang:latest

# # Set the working directory inside the container
# WORKDIR /go_project3

# # Copy the local source files to the working directory inside the container
# COPY . .

# # Build the Go executable
# RUN go build -o tafviet

# # Command to run the executable when the container starts
# CMD ["./main"]

# Build stage
FROM golang:latest AS builder
WORKDIR /go_project3
COPY . .
RUN CGO_ENABLED=0 go build -o main main.go

# Run stage
FROM debian:latest
WORKDIR /go_project3
COPY --from=builder /go_project3/main .
COPY app.env .
# COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE 1515 5432
CMD [ "/go_project3/main" ]
# ENTRYPOINT [ "/go_project3/start.sh" ]