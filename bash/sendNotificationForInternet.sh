#! /bin/bash

while true
do
    curl -d "Have internet" ntfy.sh/linux_desktop 2>/dev/null

    if [ "$?" != "0" ]
    then
        echo "No having Internet"
    else
        cowsay "Have internet"
        break
    fi
    sleep 3
done
