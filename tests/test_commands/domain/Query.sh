# pre-run
tau new -y domain \
    --name someDomain1 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn domain1-name0.com \
    --cert-type auto \
    --no-generated-fqdn

tau new -y domain \
    --name someapp2 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn domain2-name0.com \
    --cert-type auto \
    --no-generated-fqdn

tau new -y domain \
    --name someapp3 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn domain3-name0.com \
    --cert-type auto \
    --no-generated-fqdn

tau new -y domain \
    --name someapp13 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn domain4-name0.com \
    --cert-type auto \
    --no-generated-fqdn

tau delete -y domain \
    --name someapp13

tau new -y domain \
    --name someapp4 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn domain5-name0.com \
    --cert-type auto \
    --no-generated-fqdn

tau new -y domain \
    --name someapp5 \
    --description "some domain description" \
    --tags "tag1, tag2,   tag3" \
    --fqdn domain6-name0.com \
    --cert-type auto \
    --no-generated-fqdn

# command
tau query domain \
    --list

