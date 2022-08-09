#! /bin/bash

set -e

if [ -z $REPOS_HOME ]; then
    echo "The enviroment variable REPOS_HOME is required"
    exit 1
fi

CURRENT_DIR=${PWD}
if [ -d $REPOS_HOME ]; then

    for repository in $(ls $REPOS_HOME); do
        if [ -d $REPOS_HOME/$repository ] && [ -d $REPOS_HOME/$repository/.git ]
        then
            cd $REPOS_HOME/$repository
            echo "The actually branch in the repository: $repository" \
                 " is: $(git branch --show-current)"
        fi
    done
fi

cd $CURRENT_DIR

