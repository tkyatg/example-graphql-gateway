#!/bin/sh
git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

cd /app &&\
    fresh