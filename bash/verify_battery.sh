#! /bin/bash

batstat="$(/usr/bin/cat /sys/class/power_supply/BAT1/status)"
id_notification=0
while true; do
    battery="$(/usr/bin/cat /sys/class/power_supply/BAT1/capacity)"
    if [ $batstat == 'Unknown' ]; then
        notify-send "This computer not contains battery, finalizaing script" -u low
        killall verify_battery.sh
    elif [[ $battery -ge 1 ]] && [[ $battery -le 30 ]]; then
        batstat="$(/usr/bin/cat /sys/class/power_supply/BAT1/status)"
        icon="battery-caution"
        if [ "$batstat" == "Charging" ]; then
            icon="battery-low-charging"
        fi
        number=$(cat /sys/class/power_supply/BAT1/capacity)

        new_id=$(notify-send -u critical -i $icon "Battery low percentaje: $number%" -p --replace-id $id_notification)
        id_notification=$new_id
    fi
    sleep 1
done
