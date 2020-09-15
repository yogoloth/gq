
build: fresh static

static:
	go build --gcflags '-I ./modules/include ' --ldflags '-linkmode external -extldflags "-Wl,-Bstatic  -L ./modules/lib -ljq -lonig -Wl,-Bdynamic -lm -lc"'

dynamic:
	go build

fresh:
	git submodule update --init
	(cd modules/jq; git checkout jq-1.6 )
	(cd modules/jq; git submodule update --init)
	(cd modules/jq; autoreconf -if )
	(cd modules/jq;	./configure CFLAGS=-fPIC --disable-maintainer-mode --enable-all-static \
		--disable-shared --disable-docs --disable-valgrind --with-oniguruma=builtin -prefix=$(shell pwd)/modules)
	(cd modules/jq;	make )
	(cd modules/jq; make install-libLTLIBRARIES install-includeHEADERS;)
	cp ./modules/jq/modules/oniguruma/src/.libs/libonig.a modules/lib/
	cp ./modules/jq/modules/oniguruma/src/.libs/libonig.la modules/lib/
	cp ./modules/jq/modules/oniguruma/src/.libs/libonig.lai modules/lib/
	cp ./modules/jq/modules/oniguruma/src/oniguruma.h modules/include/

