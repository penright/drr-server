FROM golang:alpine
ENV TZ="America/Chicago"
RUN date
WORKDIR /app
COPY . .
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /ddr-server
EXPOSE 9091
ENTRYPOINT [ "/ddr-server" ]
