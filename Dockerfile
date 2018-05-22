# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/coppetti/graphql-go

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/mnmtanish/go-graphiql
RUN go get github.com/graphql-go/graphql

RUN go install github.com/coppetti/graphql-go

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/graphql-go

# Document that the service listens on port 3000.
EXPOSE 3000