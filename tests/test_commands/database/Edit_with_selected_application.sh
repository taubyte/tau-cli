# pre-run
tau new -y database \
    --name someDB \
    --description "some database description" \
    --tags "tag1, tag2,   tag3" \
    --no-local \
    --encryption \
    --match someMatch \
    --no-regex \
    --key somekey \
    --min 10 \
    --max 112 \
    --size 10 \
    --size-unit GB

# children
tau edit -y database \
    --name someDB \
    --description "some hola description" \
    --tags "tag1, tag28913,   tag3" -local \
    --no-encryption \
    --regex \
    --match someMatch \
    --min 42 \
    --max 432 \
    --size 43GB

# command
tau query database someDB

