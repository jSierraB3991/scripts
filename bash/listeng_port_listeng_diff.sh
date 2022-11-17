#! /bin/bash

icon_folder="/usr/share/icons/Papirus/64x64/apps"
old_ports=$(ss -tulpn | grep LISTEN | awk '{print $5}' | grep "^[1 | 0 | * ]")

function send_notification() {
    status=$1
    port=$2
    echo "notify-send PORT SCAN PORT $port is $status -a PORT_SCAN -i breeze-settings.svg"
    notify-send "PORT SCAN" "PORT $port is $status" -a PORT_SCAN -i $icon_folder/breeze-settings.svg
}

while true; do
    new_ports=$(ss -tulpn | grep LISTEN | awk '{print $5}' | grep "^[1 | 0 | * ]")
    type_change=$(diff <(echo "$old_ports") <(echo "$new_ports") | head -1)
    port=$(diff <(echo "$old_ports") <(echo "$new_ports") | tail -1 | awk 'BEGIN{FS=":"} {print $2}' )
    if [ "$type_change" != "" ] && [ "$port" != "" ]; then
        echo "type: $type_change port: $port"
        if [ "$(echo $type_change | grep '[0-9]d[0-9]')" != "" ]  ; then
            send_notification "DOWN" $port
        elif [ "$(echo $type_change | grep '[0-9]a[0-9]')" != "" ]  ; then
            send_notification "UP" $port
        fi
    fi
    old_ports=$new_ports
    sleep 1
done
