#! /bin/bash

FOLDER_WALLPAPER="$HOME/Source/Wallpapers"
TIME=5 #in minutes

while true; do
    cowsay "CI $(date)"
    feh --randomize --bg-fill $FOLDER_WALLPAPER/*
    sleep $((TIME * 60))
done
