.PHONY: sdk-doc

sdk-doc:
	@go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	@gomarkdoc --header '# SDK文档' --exclude-dirs ./internal/... -o SDK_DOC.md ./...
