# pre-run
tau new -y domain \
    --name test_domain_1 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn hal.computers.com \
    --cert-type auto \
    --no-generated-fqdn

# children
tau new -y function \
    --name test_function \
    --description "some function description" \
    --tags "tag1, tag2,   tag3" \
    --timeout 10s \
    --memory 10 \
    --memory-unit GB \
    --type http \
    --no-use-template \
    --domains test_domain_1 \
    --method get \
    --paths / \
    --source . \
    --call ping

# command
tau query function test_function

