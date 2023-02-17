# pre-run
tau new -y smartops \
    --name test_smartops1 \
    --description "some smartops description" \
    --tags "tag1, tag2,   tag3" \
    --ttl 10s \
    --memory 10 \
    --memory-unit GB \
    --no-use-template \
    --source . \
    --call ping

tau new -y smartops \
    --name test_smartops2 \
    --description "some smartops description" \
    --tags "tag1, tag2,   tag3" \
    --ttl 10s \
    --memory 10 \
    --memory-unit GB \
    --no-use-template \
    --source . \
    --call ping

tau new -y smartops \
    --name test_smartops3 \
    --description "some smartops description" \
    --tags "tag1, tag2,   tag3" \
    --ttl 10s \
    --memory 10 \
    --memory-unit GB \
    --no-use-template \
    --source . \
    --call ping

tau delete -y smartops \
    --name test_smartops3

tau new -y smartops \
    --name test_smartops4 \
    --description "some smartops description" \
    --tags "tag1, tag2,   tag3" \
    --ttl 10s \
    --memory 10 \
    --memory-unit GB \
    --no-use-template \
    --source . \
    --call ping

tau new -y smartops \
    --name test_smartops5 \
    --description "some smartops description" \
    --tags "tag1, tag2,   tag3" \
    --ttl 10s \
    --memory 10 \
    --memory-unit GB \
    --no-use-template \
    --source . \
    --call ping

# command
tau query smartops \
    --list

