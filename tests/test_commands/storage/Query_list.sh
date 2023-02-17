# pre-run
tau new -y storage \
    --name someStrg1 \
    --description "some storage description" \
    --tags "tag1, tag2,   tag3" \
    --bucket Streaming \
    --ttl 20s \
    --no-regex \
    --match test/v1 \
    --public \
    --size 10 \
    --size-unit GB

tau new -y storage \
    --name someStrg2 \
    --description "some storage description" \
    --tags "tag1, tag2,   tag3" \
    --bucket Streaming \
    --ttl 20s \
    --no-regex \
    --match test/v1 \
    --public \
    --size 10 \
    --size-unit GB

tau new -y storage \
    --name someStrg3 \
    --description "some storage description" \
    --tags "tag1, tag2,   tag3" \
    --bucket Streaming \
    --ttl 20s \
    --no-regex \
    --match test/v1 \
    --public \
    --size 10 \
    --size-unit GB

tau delete -y storage \
    --name someStrg3

tau new -y storage \
    --name someStrg4 \
    --description "some storage description" \
    --tags "tag1, tag2,   tag3" \
    --bucket Streaming \
    --ttl 20s \
    --no-regex \
    --match test/v1 \
    --public \
    --size 10 \
    --size-unit GB

tau new -y storage \
    --name someStrg5 \
    --description "some storage description" \
    --tags "tag1, tag2,   tag3" \
    --bucket Streaming \
    --ttl 20s \
    --no-regex \
    --match test/v1 \
    --public \
    --size 10 \
    --size-unit GB

# command
tau query storage \
    --list

