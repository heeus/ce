#! /bin/bash

_user=cect
_url=https://github.com/heeus/ce
_service="ce"

if [[ $OSTYPE = linux* ]]; then
  echo "your OS is linux"
else
  echo not supported OS: $OSTYPE 
  exit
fi

createFService(){
  fName=./ce.service
  echo "[Unit]" > $fName
  echo "Description=heeus CE" >> $fName
  echo >> $fName
  echo "[Service]" >> $fName
  echo "Type=idle" >> $fName
  echo "ExecStart=/usr/local/bin/ce" >> $fName
  echo "KillMode=process" >> $fName
  echo "Restart=on-failure" >> $fName
  echo >> $fName
  echo "[Install]" >> $fName
  echo "WantedBy=multi-user.target" >> $fName
}

terminateScript(){
  sudo rm ./ce
  sudo rm ./ce.service
  echo "install terminated"
  exit
}

stopSerciceIfExists(){
  service="ce"
  s=$(sudo systemctl is-active $service)
  if [ "$s" == "active" ]; then
    sudo systemctl stop $service
    echo "the service $service is stopped"
  fi
}

needCreateUser(){
  if grep $_user/etc/passwd; then
    return true
  else
    return false
  fi
}

createUser(){
  if sudo useradd $_user --system --no-create-home; then
    echo "User $_user created!"
  fi
}

getFiles(){
  echo Getting a distribution from a remote repository
  if sudo wget "${_url}/releases/download/v0.5.0/ce_v0.5.0_linux_amd64.tar.gz"; then
    echo Unpacking the downloaded archive
    if sudo tar -x -f ce_v0.5.0_linux_amd64.tar.gz; then
      echo Successfully!
      sudo rm ./ce_v0.5.0_linux_amd64.tar.gz
    fi
  fi
}

serviceExists(){
  s=$(sudo systemctl is-active ${_service})
  if [ "$s" == "active" ]; then
    return true
  else
    return false
  fi
}

stopService(){
  sudo systemctl stop $service
  echo "the service $service is stopped"
}

setupService(){
  echo "copy ./ce to /usr/local/bin/"         
  if !(sudo cp --force ./ce /usr/local/bin/); then
    return
  fi

  if !(sudo chown $_user /usr/local/bin/ce); then
    return
  fi

  echo "copy ./ce.service to /lib/systemd/system/" 
  if !(sudo cp --force ./ce.service  /lib/systemd/system/); then
    return
  fi

  echo "adding a service $service to startup"
  if !(sudo systemctl enable ce.service); then
    return
  fi

  if !(sudo systemctl start ce.service); then
    return
  fi


}


if $needCreateUser; then
  if !(createUser); then
    terminateScript
  fi
fi

if $serviceExists; then
  if !(stopService); then
    terminateScript
  fi
fi

if !(getFiles); then
  terminateScript
fi

if !(createFService); then
  terminateScript
fi

if !(setupService); then
  terminateScript
fi

echo "Congratulations! The service $service has been successfully installed!"

sudo rm ./ce
sudo rm ./ce.service

sudo ce help
