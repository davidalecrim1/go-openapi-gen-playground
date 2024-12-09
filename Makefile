goapi-gen-spec:
	@if goapi-gen --version 2>/dev/null; then \
		: ; \
	else \
		echo "The \"goapi-gen\" package is not installed" && \
		exit 1; \
	fi

	@if mkdir ./internal/api/generated 2>/dev/null; then \
		echo "Creating \"./internal/api/generated\" folder... "; \
	else \
		echo "The \"./internal/api/generated\" is already created. Moving on..."; \
	fi
	
	@goapi-gen --package generated -o internal/api/generated/todo.gen.go internal/api/openapi.yaml && \
	echo "Successfully completed!"

run:
	go run ./cmd/api/main.go

healthz:
	curl -v http://localhost:8080/todos

.PHONY: goapi-gen-spec
