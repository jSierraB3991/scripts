#! /bin/bash

function get_battery() {
    battery=$1
    if [[ $battery -le 10 ]]; then
        echo "10"
    elif [[ $battery -le 20 ]]; then
        echo "20"
    elif [[ $battery -le 30 ]]; then
        echo "30"
    elif [[ $battery -le 40 ]]; then
        echo "40"
    fi
}

batstat="$(/usr/bin/cat /sys/class/power_supply/BAT1/status)"
while true; do
    battery="$(/usr/bin/cat /sys/class/power_supply/BAT1/capacity)"
    if [ $batstat == 'Unknown' ]; then
        notify-send "This computer not contains battery, finalizaing script" -u low
        killall verify_battery.sh
    elif [[ $battery -ge 1 ]] && [[ $battery -le 40 ]]; then
        batstat="$(/usr/bin/cat /sys/class/power_supply/BAT1/status)"
        icon="battery-level-$(get_battery $battery)-symbolic.symbolic.png"
        if [ "$batstat" == "Charging" ]; then
            icon="battery-level-$(get_battery $battery)-charging-symbolic.symbolic.png"
        fi
        icon=$(echo "/usr/share/icons/Adwaita/64x64/status/$icon")
        number=$(cat /sys/class/power_supply/BAT1/capacity)
        notify-send "Battery Verify" "Battery low percentaje: $number%" -u critical -i $icon -h string:x-canonical-private-synchronous:low_power
    fi
    sleep 1
done
