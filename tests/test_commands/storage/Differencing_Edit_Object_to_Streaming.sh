# pre-run
tau new -y storage \
    --name someStorage \
    --description "some new storage description" \
    --tags "tag1, tag2,   tag3" \
    --bucket Object \
    --no-public \
    --versioning \
    --no-regex \
    --match some/match \
    --size 10 \
    --size-unit GB

tau edit -y storage \
    --name someStorage \
    --description "some new storage description" \
    --tags "tag1, tag2,   tag3" \
    --bucket Streaming \
    --ttl 25s \
    --public \
    --no-regex \
    --match some/match \
    --size 10 \
    --size-unit GB

# command
tau query storage someStorage

