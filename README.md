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
# example

1. convert yaml config
````
./gq -y 'if .spring.redis then del(.spring.redis.host) | del(.spring.redis.port) | .spring.redis.sentinel.master="sentinel-10.224.63.xxx-xx"  |.spring.redis.sentinel.nodes="10.224.63.xxxx:xxxx,10.224.63.xxx:xxxx,10.224.63.xxx:xxxx"  |.spring.redis.password="xxxx" else . end |del(.hadoop.zookeeper.master)|if .spring.profiles|type|test("string") then .spring.profiles="prod" else if .spring.profiles.active then .spring.profiles.active="prod" else . end end' uat/xxx-dealer-maintenance.yml >prod/xxx-dealer-maintenance.yml
'
````
2. modify k8s yaml resource
````
./gq -y '.spec.containers[0].env[.spec.containers[0].env|length]={name: "REQUEST_MEM",valueFrom:{resourceFieldRef:{containerName: .spec.containers[0].name, divisor: 0,resource: "requests.memory"}}}|".spec.containers[0].resources.requests={cpu: "100m",memory: "1Gi"}"' some_pod.yml
````
see my other project https://github.com/yogoloth/autopatch-operator

# full tutor see https://stedolan.github.io/jq/manual/
