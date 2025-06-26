# build: docker container をキャッシュ無しでbuildします
.PHONY: build
build:
	docker build --no-cache -t crypt .

# up: docker container を立ち上げます
.PHONY: up
up:
	docker run -it --init --rm --name crypt -v $$(pwd):/workspace crypt

# down: docker container を停止し削除します
.PHONY: down
down:
	docker kill crypt
