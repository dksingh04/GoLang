# Docker-go

### Build docker Image
docker build -t docker-go-first .

### Run created docker Image
docker run -it  -d -p 8080:8080 docker-go-first

### Access the application
http://localhost:8080/