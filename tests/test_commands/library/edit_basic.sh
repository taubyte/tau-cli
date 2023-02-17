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

# children
tau edit -y library someLibrary \
    --description "some new library description" \
    --tags "tag1, tag2,   tag4" \
    --path /new \
    --no-clone \
    --branch master \
    --domains hal.computers.com

# command
tau query library someLibrary

