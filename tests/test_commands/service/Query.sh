# pre-run
tau new -y service \
    --name someService1 \
    --description "some service description" \
    --tags "tag1, tag2,   tag3" \
    --protocol /testprotocol/v1 \
    --color never

tau new -y service \
    --name someapp2 \
    --description "some service description" \
    --tags "tag1, tag2,   tag3" \
    --protocol /testprotocol/v1 \
    --color never

tau new -y service \
    --name someapp3 \
    --description "some service description" \
    --tags "tag1, tag2,   tag3" \
    --protocol /testprotocol/v1 \
    --color never

tau new -y service \
    --name someapp13 \
    --description "some service description" \
    --tags "tag1, tag2,   tag3" \
    --protocol /testprotocol/v1 \
    --color never

tau delete -y service \
    --name someapp13

tau new -y service \
    --name someapp4 \
    --description "some service description" \
    --tags "tag1, tag2,   tag3" \
    --protocol /testprotocol/v1 \
    --color never

tau new -y service \
    --name someapp5 \
    --description "some service description" \
    --tags "tag1, tag2,   tag3" \
    --protocol /testprotocol/v1 \
    --color never

# command
tau query service \
    --list

