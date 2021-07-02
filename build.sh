docker build  -t kapsule-api -f ./Dockerfile-api .
docker build  -t kapsule-redirect -f ./Dockerfile-redirect .

docker run -d -i -t -p 8080:8080 kapsule-api
docker run -d -i -t -p 8081:8080 kapsule-redirect