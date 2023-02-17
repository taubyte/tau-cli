# children
tau new -y function \
    --name test_function \
    --description "some function description" \
    --tags tag1,tag2,tag3 \
    --timeout 10m \
    --memory 50 \
    --memory-unit MB \
    --type pubsub \
    --no-use-template \
    --channel doPing \
    --no-local \
    --paths / \
    --source inline \
    --call ping

# command
tau query function test_function

