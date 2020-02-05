# crawler


## setp 1
```docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.5.1```
will be the same image version on your docker

## step 2
start engin ```go run main.go```

## step 3
start fontend view, cd `/frontend`,  ```go run starter.go```
