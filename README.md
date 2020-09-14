# gq

go clone of yq (https://github.com/kislyuk/yq)

engine libjq will use embedded libjq ,and it's default option
````
./gq -y --engine libjq '.a.b="你好"' sample/test.yml
./yq -y --engine libjq '.a.b="你好"' sample/test.yml

./gq -y '.a.b="你好"' sample/test.yml
./yq -y '.a.b="你好"' sample/test.yml
````
engine libjq will call external jq like yq
````
./gq -y --engine jq '.a.b="你好"' sample/test.yml
./yq -y --engine jq '.a.b="你好"' sample/test.yml
````
