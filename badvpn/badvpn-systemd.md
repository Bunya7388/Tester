# badvpn-udpgw systemd and setup

This document contains ready-to-use commands and configuration for running badvpn/udpgw for DNSTT support.

## 1) Download prebuilt binary (if available)
```bash
curl -LO https://github.com/ambrop72/badvpn/releases/download/1.999.130/badvpn-1.999.130-linux-x86_64.zip
unzip badvpn-1.999.130-linux-x86_64.zip
sudo mv badvpn-udpgw /usr/local/bin/
sudo chmod +x /usr/local/bin/badvpn-udpgw
```

## 2) Run directly for tests
```bash
/usr/local/bin/badvpn-udpgw --listen-addr 127.0.0.1:7300 --max-clients 256 --log-level 1
```

## 3) Optional systemd unit
Create `/etc/systemd/system/badvpn-udpgw.service`:
```ini
[Unit]
Description=badvpn UDPGW service
After=network.target

[Service]
ExecStart=/usr/local/bin/badvpn-udpgw --listen-addr 127.0.0.1:7300 --max-clients 256 --log-level 1
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

## 4) Enable and start
```bash
sudo systemctl daemon-reload
sudo systemctl enable --now badvpn-udpgw.service
sudo systemctl status badvpn-udpgw.service
```

## 5) Use with DNSTT
- Configure DNSTT client to point to your badvpn UDPGW endpoint when direct UDP is blocked.
- Example: `dns of upstream = 127.0.0.1:7300` or `resolv.conf` through local proxy.
