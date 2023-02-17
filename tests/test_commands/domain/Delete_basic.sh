# pre-run
tau new -y domain \
    --name someDomain \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn domain-name8.com \
    --cert-type auto \
    --no-generated-fqdn

# command
tau delete -y domain \
    --name someDomain

