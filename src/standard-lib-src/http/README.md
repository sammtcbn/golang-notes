docker command to launch web_hello.
```
docker run --network=host --rm -it -v .:/go/src --workdir="/go/src" golang go run web_hello.go
```
