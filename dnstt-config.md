# DNSTT Performance Optimization Configuration

This file provides recommended DNSTT and badvpn/udpgw configurations to reduce latency and improve throughput in DNS tunnel setups.

## 1) DNSTT server (udp mode)

- host: `150.136.168.240`
- nameserver tunnel domain: `g.tz1.qd.je`
- upstream resolver: `169.255.187.58`

Example command:
```bash
./sldns-server -udp 150.136.168.240:53 -privkey-file server.key g.tz1.qd.je 169.255.187.58:53
```

If using `dnstt-server` instead:
```bash
dnstt-server --listen 0.0.0.0:53 --nameserver g.tz1.qd.je --resolver 169.255.187.58 --protocol udp --mtu 1232 --allow-keepalive --verbose
```

## 2) DNSTT client

Example command:
```bash
dnstt-client --server 150.136.168.240:53 --domain g.tz1.qd.je --resolver 169.255.187.58 --protocol udp --local 127.0.0.1:1080 --verbose
```

## 3) badvpn-udpgw helper (when raw UDP is blocked)

Run UDPGW locally and route DNSTT through it
```bash
badvpn-udpgw --listen-addr 127.0.0.1:7300 --max-clients 256 --log-level 1
```
Then configure dnstt client to use upstream UDP via this socket.

## 4) Systemd service (recommended)

`/etc/systemd/system/dnstt-server.service`:
```ini
[Unit]
Description=DNSTT server
After=network.target

[Service]
ExecStart=/usr/local/bin/dnstt-server --listen 0.0.0.0:53 --nameserver g.tz1.qd.je --resolver 169.255.187.58 --protocol udp --mtu 1232
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

`/etc/systemd/system/badvpn-udpgw.service`:
```ini
[Unit]
Description=badvpn UDP gateway
After=network.target

[Service]
ExecStart=/usr/local/bin/badvpn-udpgw --listen-addr 127.0.0.1:7300 --max-clients 256
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

## 5) DNS resolver tuning
- Keep upstream resolver close to server location and low latency.
- Use `nscd` or `systemd-resolved` cache on local server.
- Avoid recursive cross-cloud hop chains.

## 6) Validation
```bash
dig @150.136.168.240 g.tz1.qd.je +short
dig @127.0.0.1 -p 1080 example.com
```