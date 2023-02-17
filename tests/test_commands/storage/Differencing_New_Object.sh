# pre-run
tau new -y storage \
    --name someStorage \
    --description "some storage description" \
    --tags "tag1, tag2,   tag3" \
    --bucket Object \
    --public \
    --versioning \
    --no-regex \
    --match some/match \
    --size 10 \
    --size-unit GB

# command
tau query storage someStorage

