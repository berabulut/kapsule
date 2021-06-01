docker build . -t kapsule-service
docker run -i -t -p 8081:8081 -p 8080:8080 kapsule-service