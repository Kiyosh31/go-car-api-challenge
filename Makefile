BINEXE = ./bin/api

build:
	echo "Building..."
	go build -o $(BINEXE)

dev: build $(BINDIR)
	$(BINEXE)