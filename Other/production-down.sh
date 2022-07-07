#!/bin/bash

if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
fi

sudo systemctl stop if-else.service
sudo rm -v /etc/systemd/system/if-else.service
sudo systemctl daemon-reload
sudo rm -rfv /opt/IF-ELSE-Backend-2022
