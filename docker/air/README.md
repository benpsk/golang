# running docker with air install

## build docker
```
docker build -t go-docker-app .
```

## run docker 
```
docker run -it --rm -p 8081:8081 --name my-running-app go-docker-app
```


