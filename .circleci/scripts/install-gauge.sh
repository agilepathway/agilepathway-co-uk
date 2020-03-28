#!/bin/bash

cd e2etests || exit
npm set prefix=/home/circleci/npm && echo "export PATH=/home/circleci/npm/bin:$PATH" >> "$BASH_ENV"
# shellcheck source=/dev/null
source "$BASH_ENV"
npm install -g npm@latest
npm install
npm install -g @getgauge/cli --unsafe-perm
gauge install
gauge install html-report
gauge install screenshot
gauge install xml-report
gauge install js
npm install -g taiko --allow-root --unsafe-perm
gauge init js
