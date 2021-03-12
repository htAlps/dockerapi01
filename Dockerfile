FROM golang:latest 

LABEL maintainer="rhalpizar@gmail.com"
LABEL description="Basic Docker App"

WORKDIR /app

# For Configuration Control 
COPY go.mod .
COPY go.sum .
# RUN go mod download 

# Copy All Files 
COPY . .

# Specify Environment Vars 
ENV PORT 8080

# Compile Program 
RUN go build 

# Run Executable
CMD ["./dockerapi01"]

