# children
tau new -y domain \
    --name someDomain \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --type auto \
    --generated-fqdn

# command
tau query domain someDomain

