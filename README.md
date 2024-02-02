# AC API

This is the backend API of Toastmasters District 67 Annual Conference website.

## Format:

```
go fmt $(go list ./... | grep -v /vendor/)
go vet $(go list ./... | grep -v /vendor/)
go test -race $(go list ./... | grep -v /vendor/)
```

## Build

```
mkdir -p mybinaries
go build -o mybinaries ./...
```
