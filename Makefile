.PHONY: push
push:
	@echo "Running go fmt..."
	go fmt ./...

	@echo "Running go test..."
	go test ./... -v
	@if [ $$? -ne 0 ]; then \
		echo "❌ Tests failed. Aborting push."; \
		exit 1; \
	fi

	@echo "✅ Tests passed."
	@echo "Adding changes to git..."
	git add .

	@echo "Committing changes..."
	git commit -m "Update shared-dtos" || echo "No changes to commit."

	@echo "Pushing changes to remote repository..."
	git push origin main

	@echo "✅ All tasks completed successfully."
