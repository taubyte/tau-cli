# command
tau new -y library \
    --name someLibrary \
    --description "some library description" \
    --tags "tag1, tag2,   tag3" \
    --generate-repository \
    --private \
    --template empty \
    --branch master \
    --path / \
    --provider github \
    --no-embed-token

