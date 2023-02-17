# pre-run
tau new -y application \
    --name someapp1 \
    --description "some app desc" \
    --tags "some, other, tags" \
    --color never

tau new -y application \
    --name someapp2 \
    --description "some app desc" \
    --tags "some, other, tags" \
    --color never

tau new -y application \
    --name someapp3 \
    --description "some app desc" \
    --tags "some, other, tags" \
    --color never

tau new -y application \
    --name someapp13 \
    --description "some app desc" \
    --tags "some, other, tags" \
    --color never

tau delete -y application \
    --name someapp13

tau new -y application \
    --name someapp4 \
    --description "some app desc" \
    --tags "some, other, tags" \
    --color never

tau new -y application \
    --name someapp5 \
    --description "some app desc" \
    --tags "some, other, tags" \
    --color never

# command
tau query application \
    --list

