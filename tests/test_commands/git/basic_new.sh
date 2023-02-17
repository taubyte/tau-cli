# pre-run
tau login \
    --name taubyte-test \
    --provider github \
    --token <git-token>

tau select project testproject

# command
tau clone project \
    --no-embed-token \
    --branch master \
    --color never

