# pre-run
tau new -y service \
    --name test_service_1 \
    --description "some service description" \
    --tags "tag1, tag2,   tag3" \
    --protocol /test/v1 \
    --color never

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
    --source inline \
    --call ping

tau edit -y function \
    --name test_function \
    --description "some function description" \
    --tags tag4 \
    --timeout 10m \
    --memory 50 \
    --memory-unit MB \
    --type p2p \
    --command doPing \
    --local \
    --protocol /test/v1 \
    --source test_library \
    --call test_library.ping

# command
tau query function test_function

