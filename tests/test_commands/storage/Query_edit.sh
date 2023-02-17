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
    --tags "tag543, tag422,   tag341" \
    --bucket Streaming \
    --ttl 25s \
    --public \
    --no-regex \
    --size 15 \
    --size-unit KB \
    --match test/v1

# command
tau query storage someStorage

