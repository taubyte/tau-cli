# pre-run
tau new -y messaging \
    --name someMessaging \
    --description "some messaging description" \
    --tags "tag1, tag2,   tag3" \
    --match xpath \
    --no-local \
    --no-mqtt \
    --no-ws \
    --no-regex

# children
tau edit -y messaging \
    --name someMessaging \
    --description "some new messaging description" \
    --tags "tag1, tag2,   tag341" \
    --local \
    --regex \
    --match xpdsaath \
    --mqtt \
    --web-socket

# command
tau query messaging someMessaging

