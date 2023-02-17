# pre-run
tau new -y domain \
    --name test_domain_1 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn hal.computers.com \
    --cert-type auto \
    --no-generated-fqdn

tau new -y domain \
    --name test_domain2 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn test.computers.com \
    --cert-type auto \
    --no-generated-fqdn

tau new -y website \
    --name someWebsite \
    --description "some website description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --paths / \
    --repository-name tb_website_reactdemo \
    --no-clone \
    --branch master \
    --provider github \
    --domains test_domain_1

# children
tau edit -y website someWebsite \
    --description "some new website description" \
    --tags "tag1, tag2,   tag4" \
    --paths / \
    --no-clone \
    --branch master \
    --domains test_domain2

# command
tau query website someWebsite

