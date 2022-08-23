IMBUE_GENERATED_FILES += decorate.gen.go
IMBUE_GENERATED_FILES += invoke.gen.go
IMBUE_GENERATED_FILES += waitgroup.gen.go
IMBUE_GENERATED_FILES += with.gen.go
IMBUE_GENERATED_FILES += withgrouped.gen.go
IMBUE_GENERATED_FILES += withnamed.gen.go

GENERATED_FILES += $(IMBUE_GENERATED_FILES)

-include .makefiles/Makefile
-include .makefiles/pkg/go/v1/Makefile

.makefiles/%:
	@curl -sfL https://makefiles.dev/v1 | bash /dev/stdin "$@"

$(IMBUE_GENERATED_FILES): $(shell find internal/generate -type f)
	go run internal/generate/main.go -- $@
