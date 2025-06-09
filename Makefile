

### up: docker container を立ち上げます
.PHONY: up
up:
	docker run -it --rm --name crypt -v $$(pwd):/workspace crypt
