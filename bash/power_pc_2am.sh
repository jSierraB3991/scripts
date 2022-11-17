#! /bin/bash

icon_folder="/usr/share/icons/Papirus/64x64/apps"
id_notification=0
while true; do
    date_hour=$(date "+%H")
   
    if [ $date_hour -lt 06 ] || [ $date_hour -ge 22 ]; then
        id_notification=$(notify-send "Go To Sleep" -i "$icon_folder/xfce-system-info.svg" --replace-id $id_notification -p )
    fi

    if [ $date_hour -ge 02 ] && [ $date_hour -lt 06 ]; then
        id_notification=0
        for num in $(echo "1 2 3"); do

            faltan=30
            if [ $num -eq 2 ]; then
                faltan=20
            elif [ $num -eq 3 ]; then
                faltan=10
            fi

            id_notification=$(notify-send "The Pc shutdown in $faltan seconds" -u critical  -i "$icon_folder/system-error.svg" --replace-id $id_notification -p)
            sleep 10
        done
        systemctl poweroff
    fi
    sleep 5
done
