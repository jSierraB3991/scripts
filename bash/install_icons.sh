#! /bin/bash

FOLDER="$HOME/Descargas/icons"
COUNT_FILES="$(ls $FOLDER | wc -l )"
FOLDER_ICONS="$HOME/.icons/"

function descompress() {
    file=$1

    7z l $FOLDER/$file 2>/dev/null 1>/dev/null
    if [ "$?" == "0" ]; then
        echo "Descompress file $file"
        echo "Descomprimiendo $element"
        7z x -o$FOLDER $FOLDER/$file 1>/dev/null
        rm -r $FOLDER/$file
    fi

}

function main() {
    echo "Corrigiendo nombres"
    for i in $(seq 1 $COUNT_FILES); do
        file=$(ls $FOLDER | head -$i | tail -1)
        if [ "$(echo \"$FOLDER/$file\" | grep "\ ")" != ""  ] || [ "$file" != "\ "  ]; then
            new_name=$(echo $file | sed 's/ /_/g')
            echo "rename $file to $new_name"
            mv  "$FOLDER/$file" "$FOLDER/$new_name"
        fi
    done
    

    while true; do
        hasCompressFile=false
        echo "Preguntando si es un elemento descomprimible"
        for element in $(ls $FOLDER); do
            7z l $FOLDER/$element 1>/dev/null 2>/dev/null
            if [ "$?" != "2" ];then
                echo "el elemento $element es descomprimible"
                hasCompressFile=true
                break
            fi
        done

        if [ "$hasCompressFile" == "true" ]; then
            for element in $(ls $FOLDER); do
                7z l $FOLDER/$element 1>/dev/null 2>/dev/null
                if [ "$?" != "2" ];then
                    echo "llamar a descomprimir el elemento $element"
                    descompress $element
                fi
            done
        else
            break
        fi
    done


    moveToFolderCursorIcons
}

function moveToFolderCursorIcons() {

    for i in $(seq 1 $COUNT_FILES); do
        file=$(ls $FOLDER | head -$i | tail -1)
        if [ -d $FOLDER/$file ]; then
            mv $FOLDER/$file $FOLDER_ICONS
        fi
    done
}

main
