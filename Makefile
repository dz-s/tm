#
# Makefile for tm.txt-cli
#

INSTALL = /usr/bin/install

prefix = /usr/local

# The directory to install todo.sh in.
bindir = $(prefix)/bin
binlocal = bin/tm

build:
	mkdir -p bin
	go build -o $(binlocal)

install:
	cp $(binlocal) $(bindir)

clean:
	@echo "Cleaning up..."
	rm $(bindir)/tm