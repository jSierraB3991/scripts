#! /bin/bash

folder_icons="/usr/share/icons/Papirus-Dark/24x24/panel"

function get_battery() {
    battery=$1
    if [[ $battery -le 10 ]]; then
        echo "010"
    elif [[ $battery -le 20 ]]; then
        echo "020"
    elif [[ $battery -le 30 ]]; then
        echo "030"
    elif [[ $battery -le 40 ]]; then
        echo "040"
    elif [[ $battery -le 50 ]]; then
        echo "050"
    elif [[ $battery -le 60 ]]; then
        echo "060"
    elif [[ $battery -le 70 ]]; then
        echo "070"
    elif [[ $battery -le 80 ]]; then
        echo "080"
    elif [[ $battery -le 90 ]]; then
        echo "090"
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

        if [ "$batstat" == "Charging" ]; then
            icon="$folder_icons/battery-$(get_battery $battery)-charging.svg"
        else
            icon="$folder_icons/battery-$(get_battery $battery).svg"
        fi
        #icon=$(echo "/usr/share/icons/Adwaita/64x64/status/$icon")
        number=$(cat /sys/class/power_supply/BAT1/capacity)
        notify_id=$(notify-send "Battery Verify" "Battery low percentaje: $number%" -u critical -i $icon --replace-id $notify_id -p)
    fi
    sleep 1
done
