docker build . -t capsule-service
docker run -i -t -p 8080:8080 capsule-service