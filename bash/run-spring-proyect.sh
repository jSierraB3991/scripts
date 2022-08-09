#!/bin/bash

SHOW_WITH="echo"

set -e
if [ -z $REPOS_HOME ]; then 
    $SHOW_WITH "REPOS_HOME not configurate this enviroment variable"
    exit 1
fi

if [ $# -eq 1 ] || [ $# -eq 2 ]; then
    $SHOW_WITH "$REPOS_HOME/$1"
    if [ -d $REPOS_HOME/$1 ] && [ -d $REPOS_HOME/$1/.git ]; then
        if [ $# -eq 2 ]; then
            if [ "$2" == "-y" ] || [ "$2" == "-Y" ] || 
                [ "$2" == "-s" ] || [ "$2" == "-S" |]; then
                if [ ! -d $HOME/logs ]; then
                    mkdir $HOME/logs
                fi
                $SHOW_WITH "" > "$HOME/logs/$1.log"

                cd $REPOS_HOME/$1
                $SHOW_WITH "clean project && install dependencies, see file ~/logs/$1.log"
                if [ -d ~/target ]; then
        	    rm -rf ./target
                fi
                mvn -U -V -B -s settings.xml clean install -DskipTests=false -l "$HOME/logs/$1.log"
            fi
        fi
        if [ "$(ls $REPOS_HOME/$1/target/**.jar)" != "" ]; then
            jar=$(ls $REPOS_HOME/$1/target/**.jar)
            java -jar $jar
        else
            $SHOW_WITH "Not exist file $1.jar"
        fi
    else
        $SHOW_WITH "the folder $REPOS_HOME/$1 not exists"
    fi
else
    $SHOW_WITH "The script need almost one parameter"
fi
