# children
tau new -y smartops \
    --name test_smartops \
    --description "some smartops description" \
    --tags "tag1, tag2,   tag3" \
    --ttl 10s \
    --memory 10 \
    --memory-unit GB \
    --use-template \
    --template confirm_http \
    --lang go \
    --source . \
    --call confirmHttp

# command
tau query smartops test_smartops

