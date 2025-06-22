#====================================================================
# Reqquires https://github.com/muquit/go-xbuild-go for cross compiling
# for other platforms.
# Mar-29-2025 muquit@muquit.com 
#====================================================================
README_ORIG=./docs/README.md
README=./README.md
GLOSSARY_FILE=./docs/glossary.txt

all: build_all

build:
	echo "*** Compiling ...."
	go build -o example-cli cmd/cli/main.go
	go build -o example-server cmd/cli/main.go

build_all:
	echo "*** Cross Compiling ...."
	/bin/rm -rf bin/*
	go-xbuild-go -additional-files 'foo.txt' -config build-config.json

release:
	go-xbuild-go -release

doc:
	echo "*** Generating README.md with TOC ..."
	chmod 600 $(README)
	markdown-toc-go -i $(README_ORIG) -o $(README) --glossary ${GLOSSARY_FILE} -f
	chmod 444 $(README)

clean:
	/bin/rm -f example-cli example-server
	/bin/rm -rf ./bin/*
