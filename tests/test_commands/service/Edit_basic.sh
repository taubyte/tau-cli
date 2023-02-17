# pre-run
tau new -y service \
    --name someService \
    --description "some service description" \
    --tags "tag1, tag2,   tag3" \
    --protocol /testprotocol/v1 \
    --color never

# command
tau edit -y service \
    --name someService \
    --description "some newwenwenwenwen description" \
    --tags "tag1, tag23,   tag3" \
    --protocol /testprotocol/v2

