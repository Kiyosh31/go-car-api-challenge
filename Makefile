GO = go
RUN = run
BUILD = build
TEST = test
BINDIR = bin
EXE = api

build:
	$(GO) $(BUILD) -o $(BINDIR)/$(EXE)

dev: main.go
	$(GO) $(RUN) main.go

run: build $(BINDIR)
	$(BINDIR)/$(EXE)

test:
	$(GO) $(TEST) -v ./...