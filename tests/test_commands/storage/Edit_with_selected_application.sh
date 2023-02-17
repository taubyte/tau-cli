# pre-run
tau new -y storage \
    --name someStorage \
    --description "some storage description" \
    --tags "tag1, tag2,   tag3" \
    --bucket Streaming \
    --ttl 20s \
    --no-regex \
    --match test/v1 \
    --public \
    --size 10 \
    --size-unit GB

# children
tau edit -y storage \
    --name someStorage \
    --description "some new storage description" \
    --tags "tag1, tag2,   tag343" \
    --bucket Streaming \
    --ttl 25s \
    --no-public \
    --no-regex \
    --match some/match \
    --size 15 \
    --size-unit KB

# command
tau query storage someStorage

