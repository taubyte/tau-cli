# pre-run
tau new -y library \
    --name someLibrary \
    --description "some library description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --path / \
    --repository-name tb_website_reactdemo \
    --no-clone \
    --branch master \
    --provider github

# command
tau query library someLibrary

