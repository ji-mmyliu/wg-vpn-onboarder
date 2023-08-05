Wireguard VPN Setup Tool
==
A program designed to quickly set up a Wireguard VPN using a Hub and Spoke configuration.
To get started, either [build the program from source](https://github.com/ji-mmyliu/wg-vpn-onboarder/releases/tag/v1.0.2) or follow the quick installation.

### Quick Installation
#### Method 1: using `go install` (Recommended)
```bash
go install github.com/ji-mmyliu/wg-vpn-onboarder@latest
alias wgv=wg-vpn-onboarder # If you'd like, add this line to your shell profile (e.g. ~/.bashrc)
```

#### Method 2: Download release binary (works on `amd` systems only)
1. Download the program executable 
```bash
curl --output /tmp/wgv https://github.com/ji-mmyliu/wg-vpn-onboarder/releases/download/v1.0.6/wgv
chmod +x /tmp/wgv
```
2. Copy the executable into a folder in your system's `PATH`:
```bash
sudo mv /tmp/wgv /usr/bin/wgv
```

### Setting up a VPN
1. Prepare the VPN server and follow the instruction prompts
```bash
wgv server new
```
2. Start up the server:
```bash
wgv server up
```
3. Generate a client configuration
```bash
wgv client new
```
4. Get the client's Wireguard configuration through a method of your choice and apply it to the client device's Wireguard application:
```bash
wgv client connect
```

## Summary of Commands
```
wgv server new      Creates a new VPN network interface
wgv server up       Starts a VPN server
wgv server down     Shuts down a VPN server

wgv client new      Prepares a client configuration for a server
wgv client connect  Exports a client configuration file

wgv help            Displays this help message
wgv version         Displays the version of wgv installed
```

Written in Golang.
Copyright (C) 2023 Jimmy Liu
