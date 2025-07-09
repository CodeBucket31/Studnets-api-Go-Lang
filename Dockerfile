FROM golang:1.24.4
WORKDIR /test
COPY . /test
RUN go build /test
EXPOSE 8083
ENTRYPOINT [  "./dockergin" ]