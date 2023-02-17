# pre-run
tau new -y database \
    --name SomeDB1 \
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

tau new -y database \
    --name SomeDB2 \
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

tau new -y database \
    --name SomeDB3 \
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

tau delete -y database \
    --name SomeDB3

tau new -y database \
    --name SomeDB4 \
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

tau new -y database \
    --name SomeDB5 \
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
tau query database \
    --list

