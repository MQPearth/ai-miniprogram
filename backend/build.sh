# product
go build -ldflags "-s -w" -tags release -o output/ai-backend
docker cp output/ai-backend ai-backend:/ai-backend
docker restart ai-backend
docker logs --since 20m -f ai-backend