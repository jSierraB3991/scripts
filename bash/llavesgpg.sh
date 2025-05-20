#!/bin/bash

echo "Listado de llaves GPG instaladas en el sistema:"
echo "-----------------------------------------------"

# Obtiene todas las claves gpg instaladas
for key in $(rpm -q gpg-pubkey); do
    echo "🔑 $key"
    rpm -qi "$key" | grep -E "Summary|Build Date"
    echo "-----------------------------------------------"
done

