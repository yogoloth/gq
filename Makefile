


static:
	go build --gcflags '-I ./include ' --ldflags '-linkmode external -extldflags "-Wl,-Bstatic  -L ./libs -ljq -lonig -Wl,-Bdynamic -lm -lc"'

dynamic:
	go build
