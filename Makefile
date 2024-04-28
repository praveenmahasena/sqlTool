all: test
all: vet
all: package
all: package_race


test: vet
test: base_test

base_test:
	go test ./... -v

vet:
	go vet ./...

package: app

package_race: app_race

app:
	go build -o ./bin/sqlTool .

app_race:
	go build --race -o ./bin/sqlTool_race .
