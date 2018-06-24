GO_CMD             ?= go
GO_DEP_CMD         ?= dep
APP_NAME        	 ?= psd2-server

install-dependencies:
	@echo "Install dependencies..."
	$(GO_CMD) get ./...

install: install-dependencies
	@echo "Install $(APP_NAME)"
	$(GO_CMD) install

run: install
	@echo "Run $(APP_NAME)"
	$(APP_NAME)