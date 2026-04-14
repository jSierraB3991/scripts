#!/bin/bash

# Dirección MAC del dispositivo (ajusta con la tuya)
MAC="00:76:45:33:C7:63"

# Obtener la salida cruda de bluetoothctl info
info=$(bluetoothctl info "$MAC")

# Función para extraer campos simples
extract_field() {
  echo "$info" | grep "$1" | head -n1 | cut -d ':' -f2- | sed 's/^ *//'
}

# Función para extraer UUIDs
extract_uuids() {
  echo "$info" | grep "UUID:" | sed 's/^ *UUID: //' | awk -F '[(|)]' '{printf "  - %-30s %s\n", $1, $2}'
}

# Función para extraer ManufacturerData y AdvertisingData
extract_data_block() {
  local key="$1"
  echo "$info" | awk "/$key:/,/^[^ ]/" | grep -v "$key:" | sed 's/^ *//'
}

# Mostrar los datos
echo "🎧 Información de Dispositivo Bluetooth"
echo "--------------------------------------"
echo "Nombre:              $(extract_field 'Name:')"
echo "Alias:               $(extract_field 'Alias:')"
echo "MAC:                 $MAC"
echo "Clase:               $(extract_field 'Class:')"
echo "Icono:               $(extract_field 'Icon:')"
echo "Emparejado:          $(extract_field 'Paired:')"
echo "Conectado:           $(extract_field 'Connected:')"
echo "Confiable:           $(extract_field 'Trusted:')"
echo "Bloqueado:           $(extract_field 'Blocked:')"
echo "Batería:             $(extract_field 'Battery Percentage:')"

echo
echo "📡 Servicios Soportados (UUIDs):"
extract_uuids

echo
echo "🔧 ManufacturerData:"
extract_data_block "ManufacturerData"

echo
echo "📢 AdvertisingData:"
extract_data_block "AdvertisingData"

echo
echo "✔ Hecho."

