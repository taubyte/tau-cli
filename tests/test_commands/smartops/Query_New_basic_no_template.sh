# children
tau new -y smartops \
    --name test_smartops \
    --description "some smartops description" \
    --tags "tag1, tag2,   tag3" \
    --ttl 10s \
    --memory 10 \
    --memory-unit GB \
    --no-use-template \
    --source . \
    --call ping

# command
tau query smartops test_smartops

