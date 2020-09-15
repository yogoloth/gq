


static:
	go build --gcflags '-I ./include ' --ldflags '-linkmode external -extldflags "-Wl,-Bstatic  -L ./libs -ljq -lonig -Wl,-Bdynamic -lm -lc"'

dynamic:
	go build

fresh:
	git submodule update --init 
	(cd modules/jq; git checkout jq-1.6 )
	(cd modules/jq; git submodule update --init)
	(cd modules/jq; autoreconf -if )
	(cd modules/jq;	./configure CFLAGS=-fPIC --disable-maintainer-mode --enable-all-static \
		--disable-shared --disable-docs --disable-valgrind --with-oniguruma=builtin -prefix=`pwd`/modules/out)
	(cd modules/jq;	make )
	(cd modules/jq; make install-libLTLIBRARIES install-includeHEADERS;)
