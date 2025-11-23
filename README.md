# Provisioning Raspberry Pi OS Lite to Run Ebitengine Apps in Kiosk Mode
[pi-os-provisioning.sh](https://github.com/homebrew-arcade/rpi5-ebitengine-kiosk/scripts/pi-os-provisioning.sh)

This bash script can be run from a fresh Raspberry Pi OS Lite install on an RPi5 device.

It will install all necessary packages to run compiled Ebitengine binaries in a kiosk mode.

# Goals for Arcade Kiosk Mode Deployment
- Run homebrew arcade games on a low-powered low-cost device
- Tightly couple to Raspberry Pi for market share and availability over SBC offerings like odroid
- Start from a minimal no-desktop distro and add required packages for framebuffer display
- Kiosk mode should allow non-interactive boot directly into jailed process
- Turning on device power should bring up a default Ebitengine application without the ability to swap apps via regular input
- Ability to hot-reload when Kiosk binary is updated
- Initally target keyboard inputs with MAME mappings and [IPac interface](https://www.ultimarc.com/control-interfaces/i-pacs/)

# Raspberry Pi OS Imager Settings
Confirmed from Windows 10, Raspberry Pi Imager v 1.9.6

**Raspberry Pi Device:** Raspberry Pi 5\
**Operating System:** Raspberry Pi OS Lite (64-bit)\
**Storage:** MicroSD with >= 8GB storage

## Settings

### General
**Set hostname:** (checked) `arcade` (.local)

**Set username and password:** (checked)
**Username:** `arcade`

**Configure Wireless LAN:** (checked and configured)

### Services
**Enable SSH:** (checked)
(Public-key authentication recommended)

# First Boot

The provisioning script will attempt to run a sample Ebitengine app. This requires the script\
to be run from the device directly with a plugged in monitor and keyboard.

If the configured Kiosk application is exited the service will attempt to reload automatically.

This can actually be useful during development. If you build in an exit input sequence you can reload on-demand.
This can be used to rebuild updates git/rsync and trigger reload from the application.

```
bash -c "$(curl -fsSL https://raw.githubusercontent.com/homebrew-arcade/rpi5-ebitengine-kiosk/main/scripts/provision.sh)"
```

# Configuration and Development
