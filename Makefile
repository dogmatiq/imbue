GENERATED_FILES += imbue.gen.go

-include .makefiles/Makefile
-include .makefiles/pkg/go/v1/Makefile

.makefiles/%:
	@curl -sfL https://makefiles.dev/v1 | bash /dev/stdin "$@"

imbue.gen.go: $(shell find internal/generate -type f)
	go run internal/generate/main.go -- $@
