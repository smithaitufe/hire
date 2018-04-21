FROM golang:latest
ENV SRC_DIR /go/src/github.com/smithaitufe/hire

# Add the source fode
COPY . $SRC_DIR
WORKDIR $SRC_DIR

# Set environment configuration
RUN export PATH=$PATH:$GOPATH
RUN export GOBIN=$GOPATH/bin
RUN export PATH=$PATH:$GOBIN

# Build it
RUN go get -v github.com/smithaitufe/hire/...
RUN go get -v ./
RUN go build
CMD ["hire"]

# git config remote.config.url "https://{8bcf24a8940acef8720ea16cd2baf19537c8c372}@github.com/smithaitufe/hire.git"