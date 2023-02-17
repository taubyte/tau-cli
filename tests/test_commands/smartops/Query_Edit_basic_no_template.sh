# pre-run
tau new -y library \
    --name test_library \
    --description "some library description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --path / \
    --repository-name tb_website_reactdemo \
    --repository-id 123456 \
    --no-clone \
    --branch master \
    --provider github

# children
tau new -y smartops \
    --name test_smartops \
    --description "some smartops description" \
    --tags "tag1, tag2,   tag3" \
    --ttl 10s \
    --memory 10 \
    --memory-unit GB \
    --no-use-template \
    --source . \
    --call ping

tau edit -y smartops \
    --name test_smartops \
    --description "some smartops description" \
    --tags tag4 \
    --ttl 10m \
    --memory 50 \
    --memory-unit MB \
    --source test_library \
    --call test_library.ping

# command
tau query smartops test_smartops

