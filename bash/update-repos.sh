#! /bin/bash
FOLDER=$(echo "${PWD}")
SHOW_WITH="cowsay -animal tux"

set -e

if [ $# -eq 1 ] && [ "$1" != "-y" ]; then
    $SHOW_WITH "The parameter $1 not recognize"
    exit 1
fi

FAIL_ENV=""
if [ -z $REPOS_HOME ]; then
    FAIL_ENV="REPOS_HOME"
fi

if [ "$FAIL_ENV" != "" ]; then
    $SHOW_WITH "The enviroment variable $FAIL_ENV is required"
    exit 1
fi

for repo in $(ls $REPOS_HOME)
do
    if [ -d $REPOS_HOME/$repo ] && [ -d $REPOS_HOME/$repo/.git ]; then

	cd $REPOS_HOME/$repo
        if [ "$repo" == "$CONFIGURATION_REPO" ]; then
            if [ $# -eq 1 ] && [ "$1" == "-y" ]; then
                for branch in $(echo  $CONFIGURATION_BRANCHS)
                do
    		    if [ ! -f $branch ]; then
                        git checkout $branch
                        git pull origin $branch
                    fi
                done
            fi
        else
            branch=$(git branch --show-current)
            #branch="develop"
            git checkout $branch
            echo "------------------------------------------------------------------------"
            echo -e "-----------------Updating $repo--------------------------\n"
            git pull origin $branch
        fi

    else
        $SHOW_WITH "$repo not is a repository"
    fi
done
cd $FOLDER
