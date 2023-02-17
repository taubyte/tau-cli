# pre-run
tau new -y library \
    --name someLibrary1 \
    --description "some library description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --path / \
    --repository-name tb_website_reactdemo \
    --no-clone \
    --branch master \
    --provider github

tau new -y library \
    --name someLibrary2 \
    --description "some library description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --path / \
    --repository-name tb_website_reactdemo \
    --no-clone \
    --branch master \
    --provider github

tau new -y library \
    --name someLibrary3 \
    --description "some library description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --path / \
    --repository-name tb_website_reactdemo \
    --no-clone \
    --branch master \
    --provider github

tau delete -y library \
    --name someLibrary3

tau new -y library \
    --name someLibrary4 \
    --description "some library description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --path / \
    --repository-name tb_website_reactdemo \
    --no-clone \
    --branch master \
    --provider github

tau new -y library \
    --name someLibrary5 \
    --description "some library description" \
    --tags "tag1, tag2,   tag3" \
    --no-generate-repository \
    --path / \
    --repository-name tb_website_reactdemo \
    --no-clone \
    --branch master \
    --provider github

# command
tau query library \
    --list

