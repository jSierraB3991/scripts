#! /bin/bash

mkdir -p $HOME/.local/data
mkdir -p $HOME/Descargas

#Verify table
echo "CREATE TABLE IF NOT EXISTS programs(name varchar(100), version varchar(50), date varchar(40))" | sqlite3 ~/.local/data/ejemplo.db 1>/dev/null 2>/dev/null


function get_actual_version {
    echo "SELECT version FROM programs WHERE name = '$1'" | sqlite3 ~/.local/data/ejemplo.db
}

function update_program {
    program=$1
    version=$(get_actual_version $program)
    new_version=$2
    date_now=$(date)
    function_download=$3
    if [ "$version" != "$new_version" ]; then
        echo "Update version $version to $new_version"

        notify-send "New Version of $program $new_version" "update notifier" -u CRITICAL 

        if [ "$version" != "" ]; then
            echo "UPDATE programs SET version='$new_version', date='$date_now' WHERE name='$program'" | sqlite3 ~/.local/data/ejemplo.db
        else
            echo "INSERT INTO programs VALUES('$program', '$new_version', '$date_now')" | sqlite3 ~/.local/data/ejemplo.db
        fi
        $function_download $new_version
    fi
}

function downloading_brave {
    version=$1
    remove_v=$(echo $version | sed 's/v//g')
    url="https://github.com/brave/brave-browser/releases/download/$version/brave-browser-$remove_v-linux-amd64.zip"
    
    echo "Creating folder for downloading Brave $version"
    cd $HOME/Descargas/programs
    mkdir brave
    cd brave

    echo "Downloading new version of Brave $version"
    wget -nv $url 1>/dev/null
    echo "Unzip for file Brave $version"
    unzip *zip 1>/dev/null
    rm *.zip
    cd ..
    echo "Deleteting previous version of brave"
    sudo rm -rf /opt/brave
    echo "Moving version version of brave"
    sudo mv brave /opt/
}

function downloading_dbeaver {
    version=$1
    url="https://dbeaver.io/files/dbeaver-ce-latest-linux.gtk.x86_64.tar.gz"
     cd $HOME/Descargas/programs
    
    echo "Downloading new version of dbeaver $version"
    wget -nv $url

    echo "Descompress for file Dbeaver $version"
    tar -xzf dbeaver-ce-**-linux.gtk.x86_64.tar.gz
    rm dbeaver-ce-**-linux.gtk.x86_64.tar.gz
    echo "Moving version $version of dbeaver"
    sudo rm -rf /opt/dbeaver
    sudo mv dbeaver /opt/
}

function downloading_insomnia {
    version=$1
    url="https://github.com/Kong/insomnia/releases/download/core@$version/Insomnia.Core-$version.tar.gz"
    echo "Creating folder for downloading Insomnia $version"
    cd $HOME/Descargas/programs
    
    echo "Downloading new version of Insomnia $version"
    wget -nv $url

    echo "Descompress for file Insomnia $version"
    tar -xzf Insomnia.Core-$version.tar.gz
    rm Insomnia.Core-$version.tar.gz

    echo "Moving version $version of insomnia"
    sudo rm -rf /opt/insomnia
    sudo mv Insomnia.Core-$version /opt/insomnia

}

function downloading_linux_notification_center {
    version=$1
    echo "Creating folder for downloading linux notification center"
    cd $HOME/Descargas/programs
    
    echo "Downloading linux notification center $version"
    wget -nv https://github.com/phuhl/linux_notification_center/archive/refs/tags/$version.tar.gz

    echo "Descompress for file linux_notification_center$version"
    tar -xzf $version.tar.gz

    sudo rm -rf /opt/linux_notification_center
    sudo mv linux_notification_center-$version /opt/linux_notification_center
    rm $version.tar.gz

    echo "Installing linux_notification_center $version"
    killall deadd-notification-center

    cd /opt/linux_notification_center
    echo "downloading daemon set $version"
    wget -nv https://github.com/phuhl/linux_notification_center/releases/download/$version/deadd-notification-center
    mkdir -p .out
    mv deadd-notification-center .out
    sudo make install
    
    deadd-notification-center &
}

version_of_jdk=17
function downloading_graalvm {
    version=$1
    url="https://github.com/graalvm/graalvm-ce-builds/releases/download/vm-$version/graalvm-ce-java$version_of_jdk-linux-amd64-$version.tar.gz"

    echo "Creating folder for downloading GraalVM $version_of_jdk $version"
    cd $HOME/Descargas/programs
    
    echo "Downloading new version of GraalVM $version"
    wget -nv $url

    echo "Descompress for file GraalVM $version"
    tar -xzf graalvm-ce-java$version_of_jdk-linux-amd64-$version.tar.gz
    rm graalvm-ce-java$version_of_jdk-linux-amd64-$version.tar.gz
    sudo rm -rf /opt/graalvm
    sudo mv graalvm-ce-java$version_of_jdk-$version /opt/graalvm


}

function download_libre_sprite {
    version=$1
    url="https://github.com/LibreSprite/LibreSprite/releases/download/$version/libresprite-development-linux-x86_64.zip"
    
    echo "change folder to download"
    cd $HOME/Descargas/programs

    echo "Downloading LibreSprite Version $version"
    wget -nv $url -O libre_sprite.zip
    unzip libre_sprite.zip
    rm -rf libre_sprite.zip
    mkdir libre_sprite

    sudo rm -rf /opt/libre_sprite
    mv **.AppImage libre_sprite/libre_sprite.AppImage
    chmod +x libre_sprite/libre_sprite.AppImage
    sudo mv libre_sprite /opt/libre_sprite


    
}

function downloading_vscode {
    url=$(curl -s 'https://code.visualstudio.com/sha/download?build=stable&os=linux-x64' | awk -F"/" '{print $NF}')
    name_file="code.tar.gz"   
    cd $HOME/Descargas/programs
    
    echo "Downloading new version of Visual Studio Code $version"
    wget -nv $(curl -s 'https://code.visualstudio.com/sha/download?build=stable&os=linux-x64' | awk '{print $NF}') -O $name_file

    echo "Descompress for file VS Code $version"
    tar -xzf $name_file
    rm $name_file
    sudo rm -rf /opt/code
    sudo mv VSCode-linux-x64/ /opt/code
}

function checking_libre_sprite {
    echo "Get last version of libre sprite"
    last_release=$(curl -s -I https://github.com/LibreSprite/LibreSprite/releases/latest | grep location | awk '{print $2}' | awk -F"/" '{print $NF}' | tr -d '\r')
    update_program "libre_sprite" $last_release download_libre_sprite
}

function checking_brave {
    #brave
    echo "Verifing Brave"
    new_version=0

    new_version=""
    x=0
    while [ "$new_version" == ""  ]; do
        x=$((x+1))
        echo "https://github.com/brave/brave-browser/releases?page=$x"
        count_release=$(curl -s https://github.com/brave/brave-browser/releases?page=$x | grep /tag/ | awk -F '>' '{print $3}' | grep Release | awk '{print $2}' | wc -l)
        if [[ "$count_release" -ge 0 ]]; then
            #for i in {1..$count_release}; do
            for i in $(seq 1 1 $count_release); do
                relese_version=$(curl -s https://github.com/brave/brave-browser/releases?page=$x | grep /tag/ | awk -F '>' '{print $3}' | grep Release | awk '{print $2}' | sed -n "$i"p)
                echo "$i" 
                if [ "$relese_version" != "" ]; then
                    remove_v=$(echo $relese_version | sed 's/v//g')      
                    is_success=$(curl -s -I https://github.com/brave/brave-browser/releases/download/$relese_version/brave-browser-$remove_v-linux-amd64.zip | grep 404)
                if [ "$is_success" == "" ]; then
                    new_version=$relese_version
                fi
            fi
            done
        fi
    done
    update_program "brave" $new_version downloading_brave
}

function checking_dbeaver {
    #dbeaver
    echo "Verifing Dbeaver"
    new_version=$(curl -s https://dbeaver.io/download/ | grep ">DBeaver Community ")
    new_version=$(echo $new_version | sed -e 's/<[^>]*>//g')
    update_program "dbeaver" "$new_version" downloading_dbeaver
}

function checking_insomnia {
    #insomnia
    echo "Verifing Insomnia"
    new_version=""
    new_version=$(curl -s -L github.com/Kong/insomnia/releases/latest/ | grep "/Kong/insomnia/releases/tag/core" | head -1 | grep -oP 'core@\K[0-9.]+')
    update_program "insomnia" $new_version downloading_insomnia
}

#Linux notification Center
#echo "Verifing Linux Notification Center"
#new_version=$(curl -s https://github.com/phuhl/linux_notification_center/releases | grep tree | awk 'BEGIN{FS="\""}{print $2}' | head -1 | awk 'BEGIN{FS="/"}NF{print $NF}')
#update_program "linux_notification_center" "$new_version" downloading_linux_notification_center

function checking_graalvm_java {
    #dbeaver
    echo "Verifing $version_of_jdk GraalVM"
    new_version=$(curl -s https://github.com/graalvm/graalvm-ce-builds/releases | grep Edition | sed -e 's/<[^>]*>//g' | grep -v OpenJDK | grep -v container | head -1 | awk '{print $4}')
    update_program "graalvm" "$new_version" downloading_graalvm
}

function checking_vscode {
    #VsCode
    echo "Verifing VsCode for Linux"
    new_version=$(curl -s 'https://code.visualstudio.com/sha/download?build=stable&os=linux-x64' | awk -F"/" '{print $NF}')
    update_program "VsCode" $new_version downloading_vscode
}

BOLD="\e[1m"
DIM="\e[2m"
GREEN="\e[32m"
CYAN="\e[36m"
YELLOW="\e[33m"
RESET="\e[0m"

PROGRAMS=(
  "dbeaver|DBeaver (DB Manager)"
  "insomnia|Insomnia (API Client)"
  "java|Java GraalVM"
  "sprite|LibreSprite"
)

show_help() {
    echo -e "${BOLD}update_automatic.sh${RESET}"
    echo -e "${DIM}Descarga software fuera de dnf.${RESET}\n"

    echo -e "${BOLD}Uso:${RESET}"
    echo -e "  ./update_automatic.sh ${CYAN}[programas...]${RESET}\n"

    echo -e "${BOLD}Programas disponibles:${RESET}"
    for entry in "${PROGRAMS[@]}"; do
        IFS="|" read -r name desc <<< "$entry"
        printf "  ${GREEN}%-10s${RESET} %s\n" "$name" "$desc"
    done

    echo -e "\n${BOLD}Opciones:${RESET}"
    echo -e "  ${YELLOW}-h, --help${RESET}  Mostrar esta ayuda"
}


if [ $# -eq 0 ]; then
    checking_dbeaver
    checking_insomnia
    checking_libre_sprite
elif [ "$1" == "-u"  ]; then
    if [ $# -eq 2 ]; then
        if [ "$2" == "dbeaver" ]; then
            checking_dbeaver
        elif [ "$2" == "brave"  ]; then
            checking_brave
        elif [ "$2" == "insomnia" ]; then
            checking_insomnia
        elif [ "$2" == "java" ]; then
            checking_graalvm_java
        elif [ "$2" == "sprite" ]; then
            checking_libre_sprite
        elif [ "$2"  == "vscode" ] || [ "$2" == "code" ]; then
            checking_vscode
        fi
    else
        cowsay "I need two data"
    fi
elif [ "$1" == "-h" ]; then
    show_help
fi
