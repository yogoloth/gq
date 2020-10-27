# gq

go version of yq (https://github.com/kislyuk/yq)

# compile

- link with libjq.a and libonig.a in modules/lib
````
make static
````
- link with system libjq.so and libonig.so
````
make dynamic
````
- compile libjq with modules/jq(jq-1.6) and do static link
````
make build
````

# usage

engine libjq will use embedded libjq ,and it's default option
````
./gq -y --engine libjq '.a.b="你好"' sample/test.yml
./yq -y --engine libjq '.a.b="你好"' sample/test.yml

./gq -y '.a.b="你好"' sample/test.yml
./yq -y '.a.b="你好"' sample/test.yml
````
engine jq will call external jq like yq
````
./gq -y --engine jq '.a.b="你好"' sample/test.yml
./yq -y --engine jq '.a.b="你好"' sample/test.yml
````
