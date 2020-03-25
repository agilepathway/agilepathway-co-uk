#!/bin/bash

test_site_url=$(curl -H "Authorization: Bearer ${NETLIFY_ACCESS_TOKEN}" -H "User-Agent: CircleCI-${CIRCLE_PROJECT_USERNAME}-${CIRCLE_PROJECT_REPONAME}/${CIRCLE_BUILD_NUM}" "$NETLIFY_DEPLOY_PREVIEW_FUNCTION_URL"?commit="${CIRCLE_SHA1}"\&siteid="${NETLIFY_SITE_ID}" )
test_site_url=https://agilepathway-co-uk.netlify.com
echo "export TEST_SITE_URL=$test_site_url" >> "$BASH_ENV"
# shellcheck source=/dev/null
source "$BASH_ENV"
echo "TEST_SITE_URL environment variable set: $test_site_url"
