docker-compose -p 'multi' build

docker image ls --format "table {{.Repository}}\t{{.Size}}" | grep "multi"
