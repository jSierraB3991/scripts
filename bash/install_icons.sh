#! /bin/bash

FOLDER="$HOME/Descargas/icons"
COUNT_FILES="$(ls $FOLDER | wc -l )"
FOLDER_ICONS="$HOME/.icons/"

function descompress() {
    file=$1
    7z l $FOLDER/$file 2>/dev/null 1>/dev/null
    if [ "$?" == "0" ]; then
        echo "Descompress file $file"
        7z x $FOLDER/$file 1>/dev/nul
        rm -r $FOLDER/$file
    fi

}

function main() {
    for i in $(seq 1 $COUNT_FILES); do
        file=$(ls $FOLDER | head -$i | tail -1)
        if [ -f $FOLDER/$file ]; then
            if [ "$(echo \"$FOLDER/$file\" | grep "\ ")" != ""  ]; then
                new_name=$(echo $file | sed 's/ /_/g')
                echo "rename $file to $new_name"
                mv  "$FOLDER/$file" "$FOLDER/$new_name"
            fi
    
            descompress $file
        fi
    done
}

function moveToFolderCursorIcons() {

    for i in $(seq 1 $COUNT_FILES); do
        file=$(ls $FOLDER | head -$i | tail -1)
        if [ -d $FOLDER/$file ]; then
            mv $FOLDER/$file $FOLDER_ICONS
        fi
    done
}

while true; do
    hasCompressFile=false
    for element in $(lsd $FOLDER); do
        7z l $FOLDER/$element 1>/dev/null 2>/dev/null
        if [ "$?" != "2" ];then
            hasCompressFile=true
        fi
    done

    if [ "$hasCompressFile" == "true" ]; then
        main
    else
        break
    fi
done

moveToFolderCursorIcons
