.PHONY: gen-docs
gen-docs:
	@swag init -g ./web/main.go -d cmd,internal && swag fmt