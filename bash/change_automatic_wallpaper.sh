#! /bin/bash

FOLDER_WALLPAPER="$HOME/Source/Wallpapers"
#Tiemp en minutes
TIME=500

while true; do
    picture_number=$(echo $((1 + $RANDOM % $(ls $FOLDER_WALLPAPER | wc -l) )))
    echo "image number $picture_number"
    image=""
    count=1
    for wallpaper in $(ls $FOLDER_WALLPAPER); do
        if [ $count == $picture_number ]; then
            image=$wallpaper
            break
        fi
        let count+=1
    done
    echo "change background with $image"
    feh --bg-fill $FOLDER_WALLPAPER/$image
    sleep $TIME
done
