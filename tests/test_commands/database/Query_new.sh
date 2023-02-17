# children
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
    --path /SOMEPATH \
    --size 10 \
    --size-unit GB

# command
tau query database someDB

