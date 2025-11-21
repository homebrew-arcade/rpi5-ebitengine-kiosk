#!/usr/bin/bash

# Readme
# https://github.com/homebrew-arcade/rpi5-ebitengine-kiosk/scripts/pi-os-provisioning.md

# Install package dependencies
echo "Updating apt"
sudo apt update
sudo apt upgrade
echo "Installing required packages"
sudo apt install gcc git zsh golang-go xorg cage libc6-dev libgl1-mesa-dev libxcursor-dev \
    libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config gldriver-test

# install ohmyzsh and set as default shell
echo "Installing oh-my-zsh. Set ZSH to default shell"
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
zsh

# clone repo into ~
echo "Cloning repo homebrew-arcade/rpi5-ebitengine-kiosk into /home/arcade/rpi5-ebitengine-kiosk"
git clone https://github.com/homebrew-arcade/rpi5-ebitengine-kiosk.git /home/arcade/rpi5-ebitengine-kiosk

# build and run sample app
echo "Building sample Ebitengine app"
go build /home/arcade/rpi5-ebitengine-kiosk/cmd/ebitentest/ebitentest.go /home/arcade/ebitentest
echo "Dumping Kiosk ENV vars to /home/arcade/.zshrc"
echo '' >> /home/arcade/.zshrc
echo '# Kiosk settings' >> /home/arcade/.zshrc
echo 'export KIOSK_APP_PATH="/home/arcade/ebitentest"' >> /home/arcade/.zshrc
echo 'export KIOSK_APP_WATCH=1' >> /home/arcade/.zshrc
source /home/arcade/.zshrc
echo "Running sample app directly in Cage kiosk"
echo "Will exit after 15 seconds. Press Q or ESC to exit immediately"
read -n 1 -s -p "Press any key to run..."
cage $KIOSK_APP_PATH

# configure startup service
echo "Building Kiosk app"
go build /home/arcade/rpi5-ebitengine-kiosk/cmd/kiosk/kiosk.go /home/arcade/kiosk
echo "Configuring systemd process for boot"
sudo cp /home/arcade/rpi5-ebitengine-kiosk/scripts/kiosk.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable kiosk.service
echo "Re-running sample app from kiosk service"
echo "Will exit after 15 seconds. Press Q or ESC to exit immediately"
read -n 1 -s -p "Press any key to run..."
sudo systemctl start kiosk.service
sudo systemctl stop kiosk.service

# configure auto-login
echo "Using raspi-config to configure automatic login for arcade user"
echo "System Options > Boot / Auto Login > B1 Console Autologin"
read -n 1 -s -p "Press any key to continue..."
sudo raspi-config

# All done
echo "All done!"
read -n 1 -s -p "Press any key to reboot..."
sudo reboot

