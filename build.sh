docker build  -t kapsule-api -f ./cmd/api/Dockerfile .
docker build  -t kapsule-redirect -f ./cmd/redirect/Dockerfile .

docker run -d -i -t -p 8080:8080 kapsule-api
docker run -d -i -t -p 8081:8081 kapsule-redirect