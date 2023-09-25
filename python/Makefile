.DEFAULT_GOAL := build

clean:

update:

update-release:

format: 
	yapf -ri src
#	yapf -ri test
	yapf -ri examples

build: format

test: build

vet: 

lint: 

build-all: test vet lint

release: build-all 
	rm -rf dist/*
	python3 -m build
# 	python3 -m twine check dist/* 

publish: release
	echo "Releasing version $(VERSION)"
# 	gh release create "$(VERSION)" dist/*.tar.gz --draft --prerelease --title "$(VERSION)-beta" --notes-file release-notes.md
# 	python3 -m twine upload --repository testpypi -u __token__ --skip-existing dist/*
#	python3 -m twine upload --repository pypi     -u __token__ --skip-existing dist/*

debug: build
	cd examples/basic && export PYTHONPATH=../../src && python3 main.py --data data.txt

usage: build

run: build
