#!/bin/bash

set -eu

: ${LIBRARY:?required}
: ${RELEASE_OUT:?required}
: ${NOTIFICATION_OUT:?required}
: ${REPO_URL:?required}

REPO_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )/../.." && pwd )"

VERSION_FROM=${VERSION_FROM:-$REPO_ROOT/VERSION}
VERSION=$(cat ${VERSION_FROM})

###############################################################

go get github.com/cloudfoundry/libbuildpack/packager/buildpack-packager

RELEASE_OUT=$PWD/$RELEASE_OUT
mkdir -p ${RELEASE_OUT}/artifacts
echo "v${VERSION}"             > ${RELEASE_OUT}/tag
echo "${LIBRARY} v${VERSION}"  > ${RELEASE_OUT}/name

pushd $REPO_ROOT
source .envrc

rm -f *.zip
buildpack-packager summary
buildpack-packager summary     > ${RELEASE_OUT}/notes.md
buildpack-packager build -cached -stack cflinuxfs3
mv *.zip                         ${RELEASE_OUT}/artifacts
popd


cat > ${NOTIFICATION_OUT}/message <<EOF
New ${LIBRARY} v${VERSION} released. <${REPO_URL}/releases/tag/v${VERSION}|Release notes>.
EOF
