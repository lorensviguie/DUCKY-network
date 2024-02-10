FROM golang:latest

# Set the working directory
WORKDIR /go/src/app


# Install required tools
RUN go get -u golang.org/x/tools/gopls \
    gotests \
    github.com/fatih/gomodifytags \
    github.com/josharian/impl \
    github.com/haya14busa/goplay/cmd/goplay \
    github.com/go-delve/delve/cmd/dlv \
    honnef.co/go/tools/cmd/staticcheck

CMD ["tail", "-f", "/dev/null"]
