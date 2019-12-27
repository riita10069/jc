lint:
    go tool vet app
    golint -set_exit_status app/...
fmt: lint
    goimports -w app
build: fmt lint
    go build -o main main.go
run:
    realize start -n main
