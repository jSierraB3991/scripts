#! /bin/bash

#brave
http GET  https://github.com/brave/brave-browser/releases |grep Release | grep Chromium | head -1 | awk 'BEGIN{FS=">"} { print $2}'
