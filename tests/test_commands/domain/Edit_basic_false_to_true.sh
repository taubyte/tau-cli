# pre-run
tau new -y domain \
    --name someDomain \
    --description "some newwenwenwenwen description" \
    --tags "tag1, tag23,   tag3" \
    --fqdn domain-name6.com \
    --type auto \
    --no-generated-fqdn

# command
tau edit -y domain \
    --name someDomain \
    --description "some false description" \
    --tags "tag1, tag23,   tag3" \
    --fqdn domain-name6.com \
    --type auto

