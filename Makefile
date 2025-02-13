# Copyright 2025 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
# All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

all: build test

build: main.go
	go build -o build/agicli .

run: build/agicli
	./build/agicli

test: main_test.go
	go test -v

clean: build/agicli
	rm -rf build/
