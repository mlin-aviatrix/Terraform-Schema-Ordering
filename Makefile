run:
	go mod vendor
	go run cmd/terraform-schema/main.go -- example.go

plugin:
	go build -trimpath -buildmode=plugin -o terraform-schema-ordering.so ./pkg/plugin

test:
	go test -v ./...
