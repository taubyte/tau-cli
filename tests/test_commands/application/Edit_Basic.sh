# pre-run
tau new -y application \
    --name someApp \
    --description "some app desc" \
    --tags "some, other, tags" \
    --color never

# children
tau edit -y application \
    --name someApp \
    --description "some nedwdadda" \
    --tags "some, wack, tags" \
    --color never

# command
tau query application \
    --name someApp

