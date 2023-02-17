# pre-run
tau new -y domain \
    --name test_domain_1 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn hal.computers.com \
    --cert-type auto \
    --no-generated-fqdn

# command
tau new -y website \
    --name someWebsite \
    --description "some website description" \
    --tags "tag1, tag2,   tag3" \
    --generate-repository \
    --private \
    --template html \
    --branch master \
    --paths / \
    --domains test_domain_1 \
    --provider github \
    --no-embed-token

