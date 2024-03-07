

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
WORKDIR /reddit_phake
COPY . .
RUN CGO_ENABLED=0 go build -o main main.go

# Run stage
FROM debian:latest
WORKDIR /reddit_phake
COPY --from=builder /reddit_phake/main .
COPY app.env .
# COPY start.sh .
COPY templates ./templates
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE 1515 5432
CMD [ "/reddit_phake/main" ]
# ENTRYPOINT [ "/reddit_phake/start.sh" ]