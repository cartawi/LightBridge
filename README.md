# LightBridge

LightBridge is a systemd-friendly binary deployment of the Sub2API-derived AI API gateway, with LightBridge branding, release assets, and an installer that supports installation, upgrade, rollback, and Sub2API-to-LightBridge migration.

## Quick Install

```bash
curl -fsSL https://raw.githubusercontent.com/WilliamWang1721/LightBridge/main/deploy/install.sh | sudo bash
```

After the service starts, open the setup wizard:

```text
http://YOUR_SERVER_IP:8080
```

## Upgrade LightBridge

Upgrade an existing LightBridge binary deployment to the latest release:

```bash
curl -fsSL https://raw.githubusercontent.com/WilliamWang1721/LightBridge/main/deploy/install.sh | sudo bash -s -- upgrade
```

Install or roll back to a specific release:

```bash
curl -fsSL https://raw.githubusercontent.com/WilliamWang1721/LightBridge/main/deploy/install.sh | sudo bash -s -- upgrade -v v0.0.1
```

## Migrate From Sub2API

For servers that still have the legacy Sub2API binary systemd deployment, run:

```bash
curl -fsSL https://raw.githubusercontent.com/WilliamWang1721/LightBridge/main/deploy/install.sh | sudo bash -s -- migrate -v v0.0.1
```

The migration command detects `sub2api.service`, `/opt/sub2api/sub2api`, `/opt/sub2api/LightBridge`, and `/usr/local/bin/sub2api`. It backs up the legacy deployment, copies existing config/runtime files into the LightBridge layout, installs `LightBridge.service`, disables the old `sub2api.service`, and starts LightBridge.

Migration backups are written to:

```text
/opt/LightBridge-migration-backups/<timestamp>
```

## Useful Commands

```bash
sudo systemctl status LightBridge
sudo journalctl -u LightBridge -f
sudo systemctl restart LightBridge
```

See [deploy/README.md](deploy/README.md) for detailed deployment notes.

## Links

- [LinuxDO](https://linux.do/) - A friendly developer community.
