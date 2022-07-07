#!/bin/bash

if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
fi

sed -i "s|GIN_MODE.*|GIN_MODE=release|" $(pwd | sed 's|IF-ELSE-Backend-2022.*|IF-ELSE-Backend-2022|')/.env
sudo cp -v $(pwd | sed 's|IF-ELSE-Backend-2022.*|IF-ELSE-Backend-2022|')/Other/if-else.service /etc/systemd/system
sudo cp -rfv $(pwd | sed 's|IF-ELSE-Backend-2022.*|IF-ELSE-Backend-2022|') /opt
sed -i "s|GIN_MODE.*|GIN_MODE=debug|" $(pwd | sed 's|IF-ELSE-Backend-2022.*|IF-ELSE-Backend-2022|')/.env

sudo systemctl daemon-reload
sudo systemctl start if-else.service
sudo systemctl status if-else.service
