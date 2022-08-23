#! /bin/bash

function get_battery() {
    battery=$1
    if [[ $battery -le 10 ]]; then
        echo "0"
    elif [[ $battery -le 20 ]]; then
        echo "10"
    elif [[ $battery -le 30 ]]; then
        echo "20"
    else
        echo "100"
    fi
}

batstat="$(/usr/bin/cat /sys/class/power_supply/BAT1/status)"
notify_id=0
while true; do
    battery="$(/usr/bin/cat /sys/class/power_supply/BAT1/capacity)"
    if [ $batstat == 'Unknown' ]; then
        notify-send "This computer not contains battery, finalizaing script" -u low
        killall verify_battery.sh
    elif [[ $battery -ge 1 ]] && [[ $battery -le 30 ]]; then
        batstat="$(/usr/bin/cat /sys/class/power_supply/BAT1/status)"
        icon="battery-level-$(get_battery $battery)"
        if [ "$batstat" == "Charging" ]; then
            icon="battery-level-$(get_battery $battery)-charging"
        fi
        #icon=$(echo "/usr/share/icons/Adwaita/64x64/status/$icon")
        number=$(cat /sys/class/power_supply/BAT1/capacity)
        notify_id=$(notify-send "Battery Verify" "Battery low percentaje: $number%" -u critical -i $icon --replace-id $notify_id -p)
    fi
    sleep 1
done
