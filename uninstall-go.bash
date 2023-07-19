#!/bin/bash

function show_msg()
{
MSG="$(cat <<EOF
Then, need to manually remove /usr/local/go/bin from PATH environment variables in ~/.profile
EOF
)"

echo "$MSG"
}

rm -rf --preserve-root ~/go || exit 1
sudo apt-get remove golang-go
sudo rm -rf --preserve-root /usr/local/go || exit 1
show_msg
