#!/bin/bash

if [ "${TRAVIS_PULL_REQUEST}" = "false" ]; then
    echo "Skip, because not a PR"
    exit 0
fi

DIR=`dirname $0`

#!/bin/bash

set -xe

DIR=`dirname $0`

apt-get -yqq update && apt-get install -y jq
go get -u github.com/ewgRa/ci-utils/utils/diff_liner
go get -u github.com/ewgRa/ci-utils/utils/github_comments_diff
go get -u github.com/ewgRa/ci-utils/utils/github_comments_send

git config --global core.quotepath false

git diff HEAD^ --name-status | grep "^D" -v | sed 's/^.\t//g' | grep "\.md$" > /tmp/changed_files

curl -sH "Accept: application/vnd.github.v3.diff.json" https://api.github.com/repos/$TRAVIS_REPO_SLUG/pulls/$TRAVIS_PULL_REQUEST > /tmp/pr.diff
cat /tmp/pr.diff | diff_liner > /tmp/pr_liner.json

rm -f /tmp/comments.json
touch /tmp/comments.json

while read FILE; do
    COMMIT=$(git log --pretty=format:"%H" -1 "$FILE");

    // Made checks and store result to >> /tmp/comments.json, content of this file will be similar to comments.example

    echo
done < /tmp/changed_files

jq -s '[.[][]]' /tmp/comments.json > /tmp/comments_array.json

cat /tmp/comments_array.json

OUTPUT=$(cat /tmp/comments_array.json | grep "\[]");
EXIT_CODE=$?

if [ $EXIT_CODE -ne 0 ]; then
    curl -s https://api.github.com/repos/$TRAVIS_REPO_SLUG/pulls/$TRAVIS_PULL_REQUEST/comments > /tmp/pr_comments.json

    github_comments_diff -comments /tmp/comments_array.json -exists-comments /tmp/pr_comments.json > /tmp/send_comments.json

    curl -XPOST "https://github-api-bot.herokuapp.com/send_review?repo=$TRAVIS_REPO_SLUG&pr=$TRAVIS_PULL_REQUEST&body=Thanks%20for%20PR" -d @/tmp/send_comments.json
fi

exit $EXIT_CODE
