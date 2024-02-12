check:
	GOPRIVATE=github.com/ubbeg2000 buf breaking --against https://github.com/ubbeg2000/mybuf.git#branch=main

build:
	buf generate

mock: build
	./mock.sh