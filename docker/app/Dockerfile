FROM golang:1.16.5

RUN go get github.com/cespare/reflex
COPY reflex.conf /

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]