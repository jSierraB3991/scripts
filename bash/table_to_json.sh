#! /bin/bash

function help_function() {
    echo "function de ayuda para el script"
}

databases_availables=(mysql mariadb postgres)
hosts_availables=(local podman docker)
function validate_data() {
    if [ "$TYPE_DATABASE" == "" ]; then
        echo "The database is import '-y'"
        echo "Database availables is: ${databases_availables[*]}"
        exit 1
    fi

    if [ "$DATABASE" == "" ]; then
        echo "The name of database is important '-d'"
        exit 1
    fi

    if [ "$TABLE" == "" ]; then
        echo "The name of table of database is important '-t'"
        exit 1
    fi
}

function validate_cli_program(){
    if [ "$(which $1 2>/dev/null)" == "" ]; then
        echo "$1 is not installed"
        exit 1
    fi
}

function get_comman_container() {
    if [ "$HOST" == "local" ]; then
        echo ""

    elif [ "$HOST" == "docker" ]; then
        id_db=$(sudo docker ps | grep "$TYPE_DATABASE" | head -1 | awk '{print $1}')
        if [ "$id_db" == "" ]; then
            echo "In docker, not exists database $TYPE_DATABASE"
            exit 1
        fi
        echo "sudo docker exec $id_db"

    elif [ "$HOST" == "podman" ]; then
        id_db=$(sudo podman ps | grep "$TYPE_DATABASE" | head -1 | awk '{print $1}')
        if [ "$id_db" == "" ]; then
            echo "In podman, not exists database $TYPE_DATABASE"
            exit 1
        fi
        echo "sudo podman exec $id_db"
    else
        echo "Invalid host"
    fi
}


#TODO: VALIDATE WITH HOST AND VERIFY LOCAL
function validate_host() {

    if [ "$HOST" == "local" ]; then
        echo "using host in local NOOB"

    elif [ "$HOST" == "docker" ]; then
        validate_cli_program "docker"
        echo "Get conainer for docker"
        id_db=$(sudo docker ps | grep "$TYPE_DATABASE" | head -1 | awk '{print $1}')
        if [ "$id_db" == "" ]; then
            echo "In docker, not exists database $TYPE_DATABASE"
            exit 1
        fi

    elif [ "$HOST" == "podman" ]; then
        validate_cli_program "podman"
        echo "Get conainer for podman"
        id_db=$(sudo podman ps | grep "$TYPE_DATABASE" | head -1 | awk '{print $1}')
        if [ "$id_db" == "" ]; then
            echo "In podman, not exists database $TYPE_DATABASE"
            exit 1
        fi
    else
        echo "Invalid host"
    fi
}

function validate_database() {
    command_container=$(get_comman_container)

    db_in_host=""
    if [ "$TYPE_DATABASE" == "postgres" ]; then
        db_in_host=$($command_container psql -U postgres -c "\l" | awk '{print $1}' |grep $DATABASE)
    else
        echo "Type of database incorrect"
        exit 1
    fi
    if [ "$db_in_host" == "" ]; then
        echo "In host $HOST no exists the database $DATABASE of type $TYPE_DATABASE"
        exit 1
    fi
}

function validate_table() {
    command_container=$(get_comman_container)

    table_in_db=""
    if [ "$TYPE_DATABASE" == "postgres" ]; then
        read -p "What schema using? " SCHEMA
        $command_container psql -U postgres -d $DATABASE -c "SELECT * FROM $SCHEMA.$TABLE" 2>/dev/null 1>/dev/null
        db_in_host="$?"
    else
        echo "Database incorrect"
        exit 1
    fi
    if [ "$db_in_host" == "" ]; then
        echo "In database $DATABASE of type $TYPE_DATABASE no exists database $DATABASE"
        exit 1
    fi
}

HOST="local"
TYPE_DATABASE=""
DATABASE=""
TABLE=""
SCHEMA=""
while getopts "t:d:y:l:h" FLAG; do
    case "${FLAG}" in
        t) TABLE="${OPTARG}" ;;

        d) DATABASE="${OPTARG}" ;;

        y) TYPE_DATABASE=""
            argument=${OPTARG}
            for available in ${databases_availables[*]}; do
                if [ "$available" == "$argument" ]; then
                    TYPE_DATABASE="${OPTARG}"
                    break
                fi
            done
            if [ "$TYPE_DATABASE" == "" ]; then
                echo "The $argument is not database available"
                exit 1
            fi
            ;;

        l) HOST=""
            argument="${OPTARG}"
            for available in ${hosts_availables[*]}; do
                if [ "$available" == "$argument" ]; then
                    HOST=$argument
                    break
                fi
            done
            if [ "$HOST" == "" ]; then
                echo "The $argument is not host available"
                exit 1
            fi
            ;;
        h) 
            help_function 
            exit 0
            ;;

        *) echo "Uso invalido, por favor pasar opciones '-a' o '-b'" ;;
    esac
done
function extract_data() {
    command_container=$(get_comman_container)
    if [ "$TYPE_DATABASE" == "postgres" ]; then
        number_data=$($command_container psql -U postgres -d $DATABASE -c "SELECT * FROM $SCHEMA.$TABLE;" | wc -l)
        headers_db=$($command_container psql -U postgres -d $DATABASE -c "SELECT * FROM $SCHEMA.$TABLE;" | head -1)
        $command_container psql -U postgres -d $DATABASE -c "SELECT * FROM $SCHEMA.$TABLE;"  \
            | head -$((number_data-2)) \
            | tail -$((number_data-4)) > $HOME/.local/data/to_json.txt


        JSON="[ "
        while read -r line; do
            counter_data=0
            JSON=$(echo $JSON " { ")
            for header in $headers_db; do
                if [ $header != "|" ]; then
                    counter_data=$((counter_data+1))
                    dato_db=$(echo $line | awk 'BEGIN{FS="|"} {print $'$counter_data'}')
                    
                    if [ ! $counter_data -eq 1 ]; then
                        JSON=$(echo $JSON ",")
                    fi
                    JSON=$(echo $JSON " \"$header\": \"$dato_db\"")
                fi
            done
            JSON=$(echo $JSON "},")
        done < $HOME/.local/data/to_json.txt 
        echo $JSON " ]"


    else
        echo "Database incorrect"
        exit 1
    fi
}

validate_data
validate_host
validate_database
validate_table
extract_data
