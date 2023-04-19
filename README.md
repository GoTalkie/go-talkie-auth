# go-talkie-auth
Authorization service

# creation of build images
docker buildx build -t gotalkie/go-talkie-auth --push  --platform=linux/amd64,linux/arm64,linux/arm/v7,linux/arm/v6 .