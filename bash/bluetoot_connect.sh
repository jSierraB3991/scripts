#!/bin/bash

# Selección
connect_bluetooth() {
    device=$(bluetoothctl devices | fzf | awk '{print $2}')

    # Conectar
    if [ -n "$device" ]; then
        device_name=$(bluetoothctl devices | grep $device | awk '{print $4}')
        echo "Connecting bluetooth with $device_name device with MAC ADRESS $device"
        bluetoothctl pair "$device"
        bluetoothctl trust "$device"
        bluetoothctl connect "$device"
    else
        echo "No seleccionaste ningún dispositivo"
    fi
}


connect_bluetooth
