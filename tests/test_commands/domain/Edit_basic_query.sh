# pre-run
tau new -y domain \
    --name someDomain \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn domain-name7.com \
    --cert-type auto \
    --no-generated-fqdn

tau edit -y domain \
    --name someDomain \
    --description "some newwenwenwenwen description" \
    --tags "tag1, tag23,   tag3" \
    --fqdn domain-name7.com \
    --type auto

# command
tau query domain \
    --name someDomain

