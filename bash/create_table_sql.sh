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
        echo -e "${yellowColour} What is name for column: ${endColour}" 
        read column_name
        echo -e "${yellowColour} What is type data for $column_name: ${endColour}"
        read column_type
        echo -e "${yellowColour} ¿Is $column_name null? y/n: ${endColour}"
        read is_column_null
        echo -e "${yellowColour} ¿Is $column_name with default value? y/n: ${endColour}"
        read is_column_default_value

        null_value="NOT NULL"
        if [ "$is_column_null" == "y" ] || [ "$is_column_null" == "Y" ]; then
            null_value="NULL"
        fi
        default_value=""
        if [ "$is_column_default_value" == "y" ] || [ "$is_column_default_value" == "Y" ]; then
            read -p "¿What is of the default value for $column_name?: " value_default
            default_value="DEFAULT $value_default"
        fi
        
        read -p "¿Want do you add new column? y/n: " another_column
        if [ "$another_column" == "n" ] || [ "$another_column" == "N" ]; then
            add_column=false
        fi
        
        BODY_OF_COLUMN="$BODY_OF_COLUMN $column_name $column_type $null_value $default_value,"
    done

    sequence_sub="_seq"
    sequence_id="$table_name$sequence_sub"
    key_name_sub="_pkey"

    key_name=$( echo "$table_name""$key_name_sub")
    constraint_pk="CONSTRAINT $key_name PRIMARY KEY (id)"
    pk_external=""
    read -p "¿Want do you add external id column? y/n: " have_external_id
    if [ "$have_external_id" == "y" ] || [ "$have_external_id" == "Y" ]; then
        pk_external="external_id bigint NOT NULL,"
    fi

    register_data=""
    read -p "¿Want do you add data of register? y/n: " have_register_data
    if [ "$have_register_data" == "y" ] || [ "$have_register_data" == "Y" ]; then
        register_data="created_at timestamp NULL, created_by varchar(255) NULL, deleted bool NOT NULL, updated_at timestamp NULL, updated_by varchar(255) NULL,"
    fi

    constarints_fk=""
    fk_count=1
    read -p "¿Want do you add foreign key? y/n: " have_fk_const
    while [[ "$have_fk_const" == "y" ]] || [[ "$have_fk_const" == "Y" ]]; do
        echo -e "${yellowColour} What is name for table of foreign key: ${endColour}" 
        read table_name_fk
        echo -e "${yellowColour} What is name column of $table_name with foreign key $table_name_fk: ${endColour}"
        read column_name_fk

        echo $fk_count
        name_constraint=$(echo "$table_name""_fk0""$fk_count" )

        constraints_fk="$constraints_fk CONSTRAINT $name_constraint FOREIGN KEY ($column_name_fk) REFERENCES $table_name_fk(id),"
        ((fk_count + 1))
        read -p "¿Want do you add foreign key? y/n: " have_fk_const
    done

    

    SQL="$SQL CREATE TABLE IF NOT EXISTS $table_name (id bigint NOT NULL, $pk_external $register_data $BODY_OF_COLUMN $constraint_pk, $constraints_fk);"
    SQL="$SQL ALTER TABLE $table_name OWNER TO \${flyway:user};"
    
    SQL="$SQL CREATE SEQUENCE IF NOT EXISTS $sequence_id START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;"
    SQL="$SQL ALTER TABLE $sequence_id OWNER TO \${flyway:user};"
    SQL="$SQL ALTER TABLE ONLY $table_name ALTER COLUMN id SET DEFAULT nextval('$sequence_id'::regclass);"

    echo ""
    read -p "¿Want do you add new table? y/n: " another_table
    if [ "$another_table" == "n" ] || [ "$another_table" == "N" ]; then
        add_table=false
    fi
done

echo "$SQL"
