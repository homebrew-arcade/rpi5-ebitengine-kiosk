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
mv /home/arcade/rpi5-ebitengine-kiosk/ebitentest /home/arcade/
chmod 755 /home/arcade/ebitentest
echo "Running sample app directly in Cage kiosk"
echo "Will exit after 15 seconds. Press Q or ESC to exit immediately"
read -n 1 -s -p "Press any key to run..."
cage /home/arcade/ebitentest

# building kiosk process manager
echo "Building Kiosk Watch app"
go build -C /home/arcade/rpi5-ebitengine-kiosk/cmd/kiosk kiosk.go
mv /home/arcade/rpi5-ebitengine-kiosk/kiosk /home/arcade/
chmod 755 /home/arcade/kiosk

# configure startup service
echo "Configuring kiosk startup service"
sudo cp /home/arcade/rpi5-ebitengine-kiosk/scripts/cage@.service /etc/systemd/system/cage@.service
sudo cp /home/arcade/rpi5-ebitengine-kiosk/scripts/pamconf /etc/pam.d/cage
sudo systemctl daemon-reload
sudo systemctl enable cage@tty1.service
sudo systemctl set-default graphical.target
echo "Re-running sample app from kiosk cage@.service"
echo "Will exit after 15 seconds. Press Q or ESC to exit immediately"
read -n 1 -s -p "Press any key to run..."
sudo systemctl start cage@.service
sudo systemctl stop cage@.service

# configure auto-login
#echo "Using raspi-config to configure automatic login for arcade user"
#echo "System Options > Boot / Auto Login > B1 Console Autologin"
#read -n 1 -s -p "Press any key to continue..."
#sudo raspi-config

# All done
echo "All done! Should start on boot."
echo "Edit /etc/systemd/system/cage@.service for kiosk settings"
read -n 1 -s -p "Press any key to reboot..."
sudo reboot