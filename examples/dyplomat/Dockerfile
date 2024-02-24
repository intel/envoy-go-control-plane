FROM golang:1.22@sha256:7b297d9abee021bab9046e492506b3c2da8a3722cbf301653186545ecc1e00bb

WORKDIR /go/src/dyplomat
COPY . /go/src/dyplomat

RUN go install -v ./...
CMD ["dyplomat"]
