#!/bin/bash

#Colours
greenColour="\e[0;32m\033[1m"
endColour="\033[0m\e[0m"
redColour="\e[0;31m\033[1m"
blueColour="\e[0;34m\033[1m"
yellowColour="\e[0;33m\033[1m"
purpleColour="\e[0;35m\033[1m"
turquoiseColour="\e[0;36m\033[1m"
grayColour="\e[0;37m\033[1m"


trap ctrl_c INT

function ctrl_c(){
    echo -e "\n${yellowColour}[*]${endColour}${grayColour}Saliendo${endColour}"
    tput cnorm
    exit 0
}

function helpPanel(){
    echo -e "\n${yellowColour}[*]${endColour}${grayColour} Uso: ./s4viPwnWifi.sh${endColour}"
    echo -e "\n\t${purpleColour}a)${endColour}${yellowColour} Modo de ataque${endColour}"
    echo -e "\t\t${redColour}Handshake${endColour}"
    echo -e "\t\t${redColour}PKMID${endColour}"
    echo -e "\t${purpleColour}n)${endColour}${yellowColour} Nombre de la tarjeta de red${endColour}"
    echo -e "\t${purpleColour}h)${endColour}${yellowColour} Mostrar este panel de ayuda${endColour}\n"
    exit 0
}

function dependencies(){
    tput civis
    clear
    dependencies=(aircrack-ng macchanger)

    echo -e "${yellowColour}[*]${endColour}${grayColour} Comprobando programas necesarios...${endColour}"
    sleep 2
    for program in "${dependencies[@]}"; do
        echo -ne "\n${yellowColour}[*]${endColour}${blueColour} Herramienta${endColour}${purpleColour} $program${endColour}${blueColour}...${endColour}"

        test -f /usr/bin/$program
	if [ "$(echo $?)" == "0" ]; then
            echo -e " ${greenColour}(V)${endColour}"
        else
            echo -e " ${redColour}(X)${endColour}"
            echo -e "${yellowColour}[*]${endColour}${grayColour} Instalando herramienta ${endColour}${blueColour}$program${endColour}${yellowColour}...${endColour}"
            sleep 2
        fi; sleep 1
    done
}

function startAttack(){
    echo -e "\n${yellowColour}[*] Lol ${endColour}"
}

# Main Function

if [ "$(id -u)" == "0" ]; then
    declare -i parameter_counter=0; while getopts ":a:n:h:" arg; do
        case $arg in
            a) attack_mode=$OPTARG; let parameter_counter+=1 ;;
            n) networkCard=$OPTARG; let parameter_counter+=1 ;;
	    h) helpPanel;;
	esac
    done

    if [ $parameter_counter -ne 2 ]; then
	helpPanel
    else
        dependencies
        startAttack
        tput cnorm
    fi
else
    echo -e "\n${redColour}[*] No soy root${endColour}"
fi
