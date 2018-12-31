# https://gist.github.com/linse/87f3b44d59a9d298309c
SUBDIRS = ./src/monkey

.PHONY: all clean

all clean:
	for dir in $(SUBDIRS); do \
		$(MAKE) -C $$dir -f Makefile $@; \
  done