format:
	@find . -print | grep --regex '.*\go' | xargs goimports -w -local "github.com/tkanata/go-training-v2"
serve:
	@docker-compose -f build/docker-compose.yml up
