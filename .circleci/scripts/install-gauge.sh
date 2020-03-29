#!/bin/bash

curl -SsL https://downloads.gauge.org/stable | sh
gauge install html-report
gauge install screenshot
gauge install xml-report
