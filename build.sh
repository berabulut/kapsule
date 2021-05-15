docker build . -t kapsule-service
docker run -i -t -p 8080:8080 kapsule-service