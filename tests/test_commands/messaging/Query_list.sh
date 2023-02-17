# pre-run
tau new -y messaging \
    --name someMsging1 \
    --description "some messaging description" \
    --tags "tag1, tag2,   tag3" \
    --match xpath \
    --no-local \
    --no-mqtt \
    --no-ws \
    --no-regex

tau new -y messaging \
    --name someMsging2 \
    --description "some messaging description" \
    --tags "tag1, tag2,   tag3" \
    --match xpath \
    --no-local \
    --no-mqtt \
    --no-ws \
    --no-regex

tau new -y messaging \
    --name someMsging3 \
    --description "some messaging description" \
    --tags "tag1, tag2,   tag3" \
    --match xpath \
    --no-local \
    --no-mqtt \
    --no-ws \
    --no-regex

tau delete -y messaging \
    --name someMsging3

tau new -y messaging \
    --name someMsging4 \
    --description "some messaging description" \
    --tags "tag1, tag2,   tag3" \
    --match xpath \
    --no-local \
    --no-mqtt \
    --no-ws \
    --no-regex

tau new -y messaging \
    --name someMsging5 \
    --description "some messaging description" \
    --tags "tag1, tag2,   tag3" \
    --match xpath \
    --no-local \
    --no-mqtt \
    --no-ws \
    --no-regex

# command
tau query messaging \
    --list

