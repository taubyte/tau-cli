# pre-run
tau new -y domain \
    --name test_domain_1 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn hal.computers.com \
    --cert-type auto \
    --no-generated-fqdn

tau new -y function \
    --name test_function1 \
    --description "some function description" \
    --tags "tag1, tag2,   tag3" \
    --timeout 10s \
    --memory 10 \
    --memory-unit GB \
    --type http \
    --use-template \
    --lang go \
    --template ping_pong \
    --domains test_domain_1 \
    --method get \
    --paths / \
    --source . \
    --call ping

tau new -y function \
    --name test_function2 \
    --description "some function description" \
    --tags "tag1, tag2,   tag3" \
    --timeout 10s \
    --memory 10 \
    --memory-unit GB \
    --type http \
    --use-template \
    --lang go \
    --template ping_pong \
    --domains test_domain_1 \
    --method get \
    --paths / \
    --source . \
    --call ping

tau new -y function \
    --name test_function3 \
    --description "some function description" \
    --tags "tag1, tag2,   tag3" \
    --timeout 10s \
    --memory 10 \
    --memory-unit GB \
    --type http \
    --use-template \
    --lang go \
    --template ping_pong \
    --domains test_domain_1 \
    --method get \
    --paths / \
    --source . \
    --call ping

tau delete -y function \
    --name test_function3

tau new -y function \
    --name test_function4 \
    --description "some function description" \
    --tags "tag1, tag2,   tag3" \
    --timeout 10s \
    --memory 10 \
    --memory-unit GB \
    --type http \
    --use-template \
    --lang go \
    --template ping_pong \
    --domains test_domain_1 \
    --method get \
    --paths / \
    --source . \
    --call ping

tau new -y function \
    --name test_function5 \
    --description "some function description" \
    --tags "tag1, tag2,   tag3" \
    --timeout 10s \
    --memory 10 \
    --memory-unit GB \
    --type http \
    --use-template \
    --lang go \
    --template ping_pong \
    --domains test_domain_1 \
    --method get \
    --paths / \
    --source . \
    --call ping

# command
tau query function \
    --list

