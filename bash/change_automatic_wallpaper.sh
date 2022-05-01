#! /bin/bash

FOLDER_WALLPAPER="$HOME/Source/Wallpapers"
#Tiemp en minutes
TIME=300

while true; do
    for wallpaper in $(ls $FOLDER_WALLPAPER); do
        feh --bg-fil $FOLDER_WALLPAPER/$wallpaper
        sleep $TIME
    done
done
