# children
tau new -y messaging \
    --name someMessaging \
    --description "some messaging description" \
    --tags "tag1, tag2,   tag3" \
    --match xpath \
    --no-local \
    --no-mqtt \
    --no-ws \
    --no-regex

# command
tau query messaging someMessaging

