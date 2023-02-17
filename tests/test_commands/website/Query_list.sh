# pre-run
tau new -y domain \
    --name test_domain_1 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn hal.computers.com \
    --cert-type auto \
    --no-generated-fqdn

tau new -y website \
    --name someWebsite1 \
    --description "some website description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --paths / \
    --repository-name tb_website_reactdemo \
    --no-clone \
    --branch master \
    --provider github \
    --domains test_domain_1

tau new -y website \
    --name someWebsite2 \
    --description "some website description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --paths / \
    --repository-name tb_website_reactdemo \
    --no-clone \
    --branch master \
    --provider github \
    --domains test_domain_1

tau new -y website \
    --name someWebsite3 \
    --description "some website description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --paths / \
    --repository-name tb_website_reactdemo \
    --no-clone \
    --branch master \
    --provider github \
    --domains test_domain_1

tau delete -y website \
    --name someWebsite3

tau new -y website \
    --name someWebsite4 \
    --description "some website description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --paths / \
    --repository-name tb_website_reactdemo \
    --no-clone \
    --branch master \
    --provider github \
    --domains test_domain_1

tau new -y website \
    --name someWebsite5 \
    --description "some website description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --paths / \
    --repository-name tb_website_reactdemo \
    --no-clone \
    --branch master \
    --provider github \
    --domains test_domain_1

# command
tau query website \
    --list

