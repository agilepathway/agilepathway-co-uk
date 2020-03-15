#!/bin/bash

test_site_url="https://${CIRCLE_BRANCH}--agilepathway-co-uk.netlify.com"
echo "export TEST_SITE_URL=$test_site_url" >> "$BASH_ENV"
# shellcheck source=/dev/null
source "$BASH_ENV"
echo "TEST_SITE_URL environment variable set: $test_site_url"