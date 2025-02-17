APP = samsamoohooh

.PHONY: swag
swag:
	@swag init -g cmd/samsamoohooh/samsamoohooh.go -o ./docs/swagger