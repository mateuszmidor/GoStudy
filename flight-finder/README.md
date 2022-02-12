# Flight Finder

Find flight connections between two given airports.  
Using gin-gonic web framework.



## Structure
![Logo](media/structure.png)

## Sequence
![Logo](media/sequence.png)

## Showcase

### 2-segments connection: Kraków-Las Palmas
![Logo](media/krk-mad-lpa.png)

### 3-segments connection: Kraków-Colombo
![Logo](media/krk-amm-mct-cmb.png)

## Installation script for AWS EC2 or similar
```bash
#!/usr/bin/env bash

# install docker
curl -fsSL https://get.docker.com | sh

# install systemd service
cat << EOF > /etc/systemd/system/flight-finder.service
[Unit] 
Description=Flight Finder Web Server 
After=network.target 

[Service] 
Type=simple 
Restart=always  
ExecStart=docker run --rm --name=flight-finder -p=80:80 mateuszmidor/flight-finder:latest
ExecStop=docker stop flight-finder 
                                   
[Install] 
WantedBy=multi-user.target
EOF

systemctl daemon-reload    
systemctl enable flight-finder
systemctl start flight-finder
systemctl status flight-finder
```