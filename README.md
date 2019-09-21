## CRUD psu server

## Installation

1. Download project
```bash
$ go get github.com/Speakerkfm/web-crud
$ cd $GOPATH/src/github.com/Speakerkfm/web-crud
```
2. Create database
```bash
$ make migrations
```

3. Install dependencies
```bash
$ dep ensure
```

4. Configure .env
```bash
$ cp .env.dist .env
```

5. Run server
```bash
$ make run
```

6. Run client in another terminal
```bash
$ go run cmd/client/main.go
```