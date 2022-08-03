#! /bin/bash

SQL=""

#Colours
greenColour="\e[0;32m\033[1m"
endColour="\033[0m\e[0m"
yellowColour="\e[0;33m\033[1m"
grayColour="\e[0;37m\033[1m"

trap ctrl_c INT

function ctrl_c(){
    echo -e "\n${yellowColour}[*]${endColour}${grayColour}Exiting for interruption${endColour}"
    echo "$SQL"
    exit 0
}

echo -e "${greenColour}Creating Tables For sql${endColour}"
add_table=true
while $add_table; do
    read -p "What is name for table: " table_name
    add_column=true
    BODY_OF_COLUMN=""
    while $add_column; do
        read -p "\tWhat is name for column: " column_name
        read -p "What is type data for $column_name: " column_type
        read -p "¿Is $column_name null? y/n: " is_column_null
        read -p "¿Is $column_name with default value? y/n: " is_column_default_value

        null_value="NOT NULL"
        if [ "$is_column_null" == "y" ] || [ "$is_column_null" == "Y" ]; then
            null_value="NULL"
        fi
        default_value=""
        if [ "$is_column_default_value" == "y" ] || [ "$is_column_default_value" == "Y" ]; then
            read -p "¿What is of the default value for $column_name?: " value_default
            default_value="DEFAULT $value_default"
        fi
        
        comma=","
        read -p "¿Want do you add new column? y/n: " another_column
        if [ "$another_column" == "n" ] || [ "$another_column" == "N" ]; then
            add_column=false
            comma=""
        fi
        
        BODY_OF_COLUMN="$BODY_OF_COLUMN $column_name $column_type $null_value $default_value $comma"
    done

    SQL="$SQL CREATE TABLE $table_name ( $BODY_OF_COLUMN );"
    read -p "¿Want do you add new table? y/n: " another_table
    if [ "$another_table" == "n" ] || [ "$another_table" == "N" ]; then
        add_table=false
    fi
done

echo "$SQL"
