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
    echo -e "\n${redColour}[*] Exiting...\n${endColour}"
    tput cnorm
    exit 1
}

function help_panel() {
    echo -e "\n${redColour}[!]Using ./btcAnalyzer.sh ${endColour}"
    for in in $(seq 1 80); do echo -ne "${redColour}-"; done; echo -e "${endColour}"
    echo -e "\t${grayColour}[-e]${endColour}\t${yellowColour}Exploration Modo${endColour}"
    echo -e "\t\t${purpleColour}unconfirmed_transaction${endColour}\t\t${yellowColour}Unconformied Transactions List${endColour}"
    echo -e "\t\t${purpleColour}inspect_transaction${endColour}\t\t${yellowColour}Inspect of Hash of transaction${endColour}"
    echo -e "\t\t${purpleColour}inspect_address${endColour}\t\t\t${yellowColour}Inspect of transaction of Address${endColour}"
    tput cnorm; exit 1
}

function unconfirmed_transaction {
    curl -s $unconfirmed_transaction_url | html2text > out.tmp
    hashes=$(cat out.tmp | grep "Hash" -A 1 | grep -v -E "Hash|--|Time")
    
    echo "Hash-Amount-BTC-Time" > out.table
    for hash in $hashes; do
        cantidad=$(cat out.tmp | grep "$hash" -A 6 | tail -n 1)
        bitcoin=$(cat out.tmp | grep "$hash" -A 4 | tail -n 1)
        time=$(cat out.tmp | grep "$hash" -A 2 | tail -n 1)
        echo "$hash-$cantidad-$bitcoin-$time" >> out.table
    done
    ./print_tables.sh '-' "$(cat out.table)"
    rm out.* 2>/dev/null
}

function inspect_transaction {
    cowsay -f vader "hello"
}

function inspect_address {
    cowsay -f www "hello"
}

unconfirmed_transaction_url="https://www.blockchain.com/btc/unconfirmed-transactions"

inspect_transaction_url="https://www.blockchain.com/btc/tx/"
inspect_address_transaction_url="https://www.blockchain.com/btc/address/"


parameter_counter=0
while getopts "e:h:" arg; do
    case $arg in
        "e") exploration_mode=$OPTARG; let parameter_counter+=1;;
    esac
done

tput civis

if [ $parameter_counter -eq 0 ]; then
    help_panel
fi

case $exploration_mode in
    "unconfirmed_transaction") unconfirmed_transaction;;
    "inspect_transaction") inspect_transaction ;;
    "inspect_address") inspect_address ;;
    *) cowsay -f vader "Invalid parameter $exploration_mode"; tput cnorm ; exit 1 ;;
esac
tput cnorm
