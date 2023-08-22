#!/bin/bash

function show_append_path()
{
MSG="$(cat <<EOF
Then, need to set environment variables:
Manually append below to ~/.profile
if [ -d "/usr/local/go/bin" ] ; then
    export PATH=\$PATH:/usr/local/go/bin
fi
EOF
)"

echo "$MSG"
}

VER=1.21.0
FN=go${VER}.linux-amd64.tar.gz
DLPATH=https://dl.google.com/go/${FN}
rm -rf --preserve-root ${FN}
echo download ${DLPATH} ...
#curl -O ${DLPATH} || exit 1
wget ${DLPATH} || exit 1
tar xvf ${FN} || exit 1
sudo chown -R root:root ./go || exit 1
sudo rm -rf --preserve-root /usr/local/go || exit 1
sudo mv go /usr/local || exit 1
rm -f --preserve-root ${FN} || exit 1

mkdir -p $HOME/go || exit 1

# show_append_path
sed -i '/\usr\/local\/go\/bin/d' ~/.profile
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile

exit 0

