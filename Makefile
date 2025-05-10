
.PHONY: push
push:
	@echo "Running go fmt..."
	go fmt ./...
	@echo "Running go test..."
	go test ./... -v
	@echo "Tests completed."
	@echo "Adding changes to git..."
	git add .
	@echo "Committing changes..."
	git commit -m "Update shared-dtos"
	@echo "Pushing changes to remote repository..."
	git push origin main
	@echo "Changes pushed successfully."
	@echo "All tasks completed successfully."
