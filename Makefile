all: build

.PHONY: install
install:
	@npm install gitbook-cli@latest -g
	@npm install || (cd node_modules/canvas && node-gyp rebuild)
	@gitbook install
	@# install ebook-convert
	@# 		on arch: sudo pacman -S calibre

.PHONY: serve
serve:
	@gitbook serve

.PHONY: build
build:
	@gitbook build

.PHONY: deploy
deploy:
	@./deploy.sh

.PHONY: deploy-all
deploy-all: build ebooks ebooks-cp deploy

.PHONY: ebooks-cp
ebooks-cp:
	@cp ethereum-development-with-go* _book

.PHONY: ebooks
ebooks: pdf ebook mobi

pdf: pdf-en pdf-zh

.PHONY: pdf-en
pdf-en:
	@gitbook pdf ./en ethereum-development-with-go.pdf

.PHONY: pdf-zh
pdf-zh:
	@gitbook pdf ./zh ethereum-development-with-go-zh.pdf

ebook: ebook-en ebook-zh

.PHONY: ebook-en
ebook-en:
	@gitbook epub ./en ethereum-development-with-go.epub

.PHONY: ebook-zh
ebook-zh:
	@gitbook epub ./zh ethereum-development-with-go-zh.epub

mobi: mobi-en mobi-zh

.PHONY: mobi-en
mobi-en:
	@gitbook mobi ./en ethereum-development-with-go.mobi

.PHONY: mobi-zh
mobi-zh:
	@gitbook mobi ./zh ethereum-development-with-go-zh.mobi

.PHONY: plugins-install
plugins-install:
	@gitbook install
