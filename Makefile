


#static:
#	go build --gcflags '-I /root/gq/libjq-go/out/include ' --ldflags '-linkmode external -extldflags "-static -L/root/gq/libjq-go/out/lib"'
#half_static:
#	go build --gcflags '-I /root/gq/libjq-go/out/include ' --ldflags '-linkmode external -extldflags "-Wl,-Bstatic  -L/root/gq/libjq-go/out/lib -ljq -lonig -Wl,-Bdynamic -lm"'
static:
	go build --gcflags '-I ./include ' --ldflags '-linkmode external -extldflags "-Wl,-Bstatic  -L ./libs -ljq -lonig -Wl,-Bdynamic -lm -lc"'
