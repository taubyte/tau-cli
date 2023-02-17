# pre-run
tau new -y domain \
    --name test_domain_1 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn hal.computers.com \
    --cert-type auto \
    --no-generated-fqdn

tau new -y service \
    --name test_service_1 \
    --description "some service description" \
    --tags "tag1, tag2,   tag3" \
    --protocol /test/v1 \
    --color never

tau new -y function \
    --name test_function \
    --description "some function description" \
    --tags "tag1, tag2,   tag3" \
    --timeout 10s \
    --memory 10 \
    --memory-unit GB \
    --type http \
    --use-template \
    --lang go \
    --template ping_pong \
    --domains test_domain_1 \
    --method get \
    --paths / \
    --source . \
    --call ping

# children
tau edit -y function \
    --name test_function \
    --description "some function description" \
    --tags tag1,tag2,tag3 \
    --timeout 10m \
    --memory 50 \
    --memory-unit MB \
    --type p2p \
    --command doPing \
    --no-local \
    --protocol /test/v1 \
    --source inline \
    --call ping

# command
tau query function test_function

