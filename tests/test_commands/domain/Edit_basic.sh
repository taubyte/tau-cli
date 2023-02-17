# pre-run
tau new -y domain \
    --name someDomain \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn domain-name4.com \
    --cert-type auto \
    --no-generated-fqdn

# command
tau edit -y domain \
    --name someDomain \
    --description "some newwenwenwenwen description" \
    --tags "tag1, tag23,   tag3" \
    --fqdn domain-name4.com \
    --cert-type auto

