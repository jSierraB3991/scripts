#! /bin/bash

LOG_DUNST="/home/juan-sierra/.cache/dunst.log"

HOUR=""
URGENCY=""
ICON=""
BODY=""
TITLE=""
APP=""
DEMAS=""
while read -r line;
do
    if [[ "$line" =~ ^(0[1-9]|1[0-2]):([0-5][0-9])[[:space:]]([AaPp][Mm])$ ]]; then
        HOUR=$line
    elif [ "$HOUR" != "" ] && [ "$URGENCY" == "" ]; then
        URGENCY=$line
    elif [ "$URGENCY" != "" ] && [ "$ICON" == "" ]; then
        ICON=$line
    elif [ "$URGENCY" != "" ] && [ "$ICON" != "" ] && [ "$BODY" == "" ]; then
        BODY=$line
    elif [ "$URGENCY" != "" ] && [ "$ICON" != "" ] && [ "$BODY" != "" ] && [ "$TITLE" == "" ]; then
#        if [ "$line" == ""  ]; then
#            TITLE="Notification"
#        else
            TITLE=$line
#        fi
    elif [ "$URGENCY" != "" ] && [ "$ICON" != "" ] && [ "$BODY" != "" ] && [ "$TITLE" != "" ] && [ "$APP" == "" ]; then
        APP=$line
        echo "INSERT INTO Notifications(hora, title, body, urgency, icon, program) VALUES('$HOUR', '$TITLE', '$BODY', '$URGENCY', '$ICON', '$APP')" | sqlite3 $HOME/.local/data/ejemplo.db
        HOUR=""
        URGENCY=""
        ICON=""
        BODY=""
        TITLE=""
        APP=""
    else
        if [ "$line" != "" ]; then
            DEMAS=$(echo "$DEMAS $line\n")
        fi
    fi


done < $LOG_DUNST

echo $DEMAS
echo $DEMAS > $LOG_DUNST
