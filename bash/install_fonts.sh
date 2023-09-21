#! /bin/bash

FOLDER_DOWNLOADS=$HOME/Descargas
FOLDER_FONTS=$FOLDER_DOWNLOADS/fonts

for file in $(ls $FOLDER_DOWNLOADS); do
    mkdir -p $FOLDER_FONTS
    7z l $FOLDER_DOWNLOADS/$file 2>/dev/null 1>/dev/null

    if [ "$?" == "0" ]; then
        echo "$FOLDER_DOWNLOADS/$file"
        FOLDER=${PWD}
        
        cd $FOLDER_DOWNLOADS/
        7z x $file 1>/dev/null
        rm -f $file
        mv *.ttf $FOLDER_FONTS
        cd -
    fi
done

FOLDER_FONTS_SHARE=~/.local/share/fonts
mkdir  -p $FOLDER_FONTS_SHARE
mv $FOLDER_FONTS/**.ttf $FOLDER_FONTS_SHARE 2>/dev/null

fc-cache -v
