# Copyright 2014 Dirk Jablonowski. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

all: build test install

GO=$(GOROOT)/bin/go

SRCDIR = $(shell pwd)

DIRS=\
	.\
	net\
	net/base58\
	net/head\
	net/payload\
	net/packet\
	net/optionaldata\
	event\
	subscription\
	connector\
	connector/simple\
	connector/buffered\
	connector/virtual\
	util/hash\
	util/generator\
	util/ks0066\
	util/lcdcharacter\
	util/miscellaneous\
	device\
	device/identity\
	device/name\
	device/enumerate\
	device/bricklet/ambientlight\
	device/bricklet/analogin\
	device/bricklet/analogout\
	device/bricklet/dualbutton\
	device/bricklet/dualrelay\
	device/bricklet/humidity\
	device/bricklet/io16\
	device/bricklet/io4\
	device/bricklet/lcd20x4\
	device/bricklet/piezobuzzer\
	device/bricklet/piezospeaker\
	device/bricklet/temperature\
	device/bricklet/tilt

test.dirs: $(addsuffix .test, $(DIRS))
deeptest.dirs: $(addsuffix .deeptest, $(DIRS))
clean.dirs: $(addsuffix .clean, $(DIRS))
build.dirs: $(addsuffix .build, $(DIRS))
install.dirs: $(addsuffix .build, $(DIRS))

%.clean:
	+@echo clean $*
	+@cd $*; $(GO) clean ; cd $(SRCDIR)

%.install:
	+@echo intall $*
	+@cd $*; $(GO) install ; cd $(SRCDIR)

%.test:
	+@echo test $*
	+@cd $*; $(GO) test ; cd $(SRCDIR)

%.deeptest:
	+@echo test $*
	+@cd $*; $(GO) test -v ; cd $(SRCDIR)

%.build:
	+@echo build $*
	+@cd $*; $(GO) build ; cd $(SRCDIR)

build: build.dirs

clean: clean.dirs

install: install.dirs

test: test.dirs

deeptest: deeptest.dirs

echo-dirs:
	@echo $(DIRS)
