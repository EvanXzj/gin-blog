FROM scratch

WORKDIR $GOPATH/src/github.com/EvanXzj/gin-blog
COPY . .

EXPOSE 3000

ENTRYPOINT [ "./server" ]
