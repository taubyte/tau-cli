# children
tau new -y storage \
    --name someStorage \
    --description "some storage description" \
    --tags "tag1, tag2,   tag3" -bucket Streaming \
    --ttl 20s \
    --no-regex \
    --match test/v1 -public -size 10 -size-unit GB

# command
tau query storage someStorage

