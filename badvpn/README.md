# badvpn (UDPGW)

This folder provides a straightforward setup for building and running `badvpn` with UDPGW support.

## What is badvpn/UDPGW?

`badvpn` is a suite of tunneling tools. `badvpn-udpgw` exposes a UDP proxy over a TCP or Unix socket, commonly used for route UDP over SSH tunnels or VPNs.

## Quick install (build from source)

```bash
cd /workspaces/Tester/badvpn

# clone upstream source
git clone https://github.com/ambrop72/badvpn.git src
cd src
mkdir -p build && cd build
cmake .. -DBUILD_NOTHING_BY_DEFAULT=1 -DBUILD_UDPGW=1
make -j$(nproc)

# optional install
sudo make install
```

The binary appears at `src/build/badvpn-udpgw`.

## Run UDPGW

```bash
# listen TCP for local UDP socket to upstream DNS/UDP services
./src/build/badvpn-udpgw --listen-addr 127.0.0.1:7300 --max-clients 100
```

Then configure client to use `127.0.0.1:7300` over tunnel.

## Notes

- You can change `--listen-addr` to `127.0.0.1:7300` or any available local port.
- For production, use a systemd service unit to supervise.
