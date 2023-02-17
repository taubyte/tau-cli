# pre-run
tau new -y service \
    --name test_service_1 \
    --description "some service description" \
    --tags "tag1, tag2,   tag3" \
    --protocol /test/v1 \
    --color never

tau new -y domain \
    --name test_domain_2 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn hal.computers.com \
    --cert-type auto \
    --no-generated-fqdn

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
tau new -y function \
    --name test_function \
    --description "some function description" \
    --tags tag1,tag2,tag3 \
    --timeout 10m \
    --memory 50 \
    --memory-unit MB \
    --type p2p \
    --no-use-template \
    --command doPing \
    --local \
    --protocol /test/v1 \
    --source inline \
    --call ping

tau edit -y function \
    --name test_function \
    --description "some function description" \
    --tags tag4 \
    --timeout 10m \
    --memory 50 \
    --memory-unit MB \
    --type http \
    --domains test_domain_2 \
    --method get \
    --paths /,/test \
    --source test_library \
    --call test_library.ping

# command
tau query function test_function

