#! /bin/bash

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
    wget $url 1>/dev/null
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
    url="https://dbeaver.io/files/dbeaver-ce-latest-linux.gtk.x86_64-nojdk.tar.gz"
    echo "Descargando version de DBEAVER"
}

function downloading_insomnia {
    version=$1
    url="https://github.com/Kong/insomnia/releases/download/core@$version/Insomnia.Core-$version.tar.gz"
    echo "Creating folder for downloading Insomnia $version"
    cd $HOME/Descargas/programs
    
    echo "Downloading new version of Insomnia $version"
    wget $url

    echo "Descompress for file Insomnia $version"
    tar -xzf Insomnia.Core-$version.tar.gz
    rm Insomnia.Core-$version.tar.gz
    sudo rm -rf /opt/insomnia
    sudo mv Insomnia.Core-$version /opt/insomnia

}


version_of_jdk=17
function downloading_graalvm {
    version=$1
    url="https://github.com/graalvm/graalvm-ce-builds/releases/download/vm-$version/graalvm-ce-java$version_of_jdk-linux-amd64-$version.tar.gz"

    echo "Creating folder for downloading GraalVM $version_of_jdk $version"
    cd $HOME/Descargas/programs
    
    echo "Downloading new version of GraalVM $version"
    wget $url

    echo "Descompress for file GraalVM $version"
    tar -xzf graalvm-ce-java$version_of_jdk-linux-amd64-$version.tar.gz
    rm graalvm-ce-java$version_of_jdk-linux-amd64-$version.tar.gz
    sudo rm -rf /opt/graalvm
    sudo mv graalvm-ce-java$version_of_jdk-$version /opt/graalvm


}

#brave
echo "Verifing Brave"
new_version=$(curl -s https://brave.com/latest/| grep "Release Notes <"  | head -1  | awk '{print $4}')
new_version=$(echo $new_version | sed -e 's/<[^>]*>//g')
update_program "brave" $new_version downloading_brave

#dbeaver
echo "Verifing Dbeaver"
new_version=$(curl -s https://dbeaver.io/download/ | grep ">DBeaver Community ")
new_version=$(echo $new_version | sed -e 's/<[^>]*>//g')
update_program "dbeaver" "$new_version" downloading_dbeaver

#insomnia
echo "Verifing Insomnia"
new_version=$(curl -s https://github.com/Kong/insomnia/releases | grep Insomnia | head -1 | sed -e 's/<[^>]*>//g' | awk '{print $2}')
update_program "insomnia" $new_version downloading_insomnia



#dbeaver
echo "Verifing $version_of_jdk GraalVM"
new_version=$(curl -s https://github.com/graalvm/graalvm-ce-builds/releases | grep Edition | sed -e 's/<[^>]*>//g' | head -1 | awk '{print $4}')
update_program "graalvm" "$new_version" downloading_graalvm
