#!/usr/bin/env bash

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

docker run --rm -e SONAR_HOST_URL="${SONARQUBE_URL}" \
    -e PROJECT_VERSION="$(git rev-parse --short=7 HEAD)" \
    -e BRANCH_NAME="$(git describe --all | sed 's=.*/==' | sed -r 's/[//]+/:/g')" \
    -e SONAR_LOGIN="$(pass dumbi/sonar-key)" \
    -e SONAR_PROJECT_BASE_DIR="." \
    -v "$DIR:/usr/src" \
    sonarsource/sonar-scanner-cli

    http://localhost:9000/