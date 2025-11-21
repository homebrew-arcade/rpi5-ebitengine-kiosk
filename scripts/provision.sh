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

# clone repo into ~
echo "Cloning repo homebrew-arcade/rpi5-ebitengine-kiosk into /home/arcade/rpi5-ebitengine-kiosk"
git clone https://github.com/homebrew-arcade/rpi5-ebitengine-kiosk.git /home/arcade/rpi5-ebitengine-kiosk

# build and run sample app
echo "Building sample Ebitengine app"
go build -C /home/arcade/rpi5-ebitengine-kiosk ./cmd/ebitentest/ebitentest.go
mv /home/arcade/rpi5-ebitengine-kiosk/cmd/ebitentest/ebitentest /home/arcade/
chmod 755 ebitentest
echo "Dumping Kiosk ENV vars to /home/arcade/.bashrc"
echo '' >> /home/arcade/.bashrc
echo '# Kiosk settings' >> /home/arcade/.bashrc
echo 'export KIOSK_APP_PATH="/home/arcade/ebitentest"' >> /home/arcade/.bashrc
echo 'export KIOSK_APP_WATCH=1' >> /home/arcade/.bashrc
source /home/arcade/.bashrc
echo "Running sample app directly in Cage kiosk"
echo "Will exit after 15 seconds. Press Q or ESC to exit immediately"
read -n 1 -s -p "Press any key to run..."
cage /home/arcade/ebitentest

# configure startup service
echo "Building Kiosk app"
go build -C /home/arcade/rpi5-ebitengine-kiosk/cmd/kiosk kiosk.go
mv /home/arcade/rpi5-ebitengine-kiosk/cmd/kiosk/kiosk /home/arcade/
chmod 755 kiosk
echo "Configuring dumb boot in bashrc with tty1 check"
echo "# Boot Kiosk on TTY1 start" >> /home/arcade/.bashrc
echo "if [[ \"$(tty)\" == \"/dev/tty1\"]]; then" >> /home/arcade/.bashrc
echo "  /home/arcade/kiosk" >> /home/arcade/.bashrc
echo "fi" >> /home/arcade/.bashrc

# configure auto-login
echo "Using raspi-config to configure automatic login for arcade user"
echo "System Options > Boot / Auto Login > B1 Console Autologin"
read -n 1 -s -p "Press any key to continue..."
sudo raspi-config

# All done
echo "All done! Should start boot. Edit .bashrc for ENV updates"
read -n 1 -s -p "Press any key to reboot..."
sudo reboot