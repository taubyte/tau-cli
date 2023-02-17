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

# command
tau select application \
    --name someapp1 \
    --color never

