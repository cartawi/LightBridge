# LightBridge

LightBridge 是基于 Sub2API 改造的 AI API Gateway 二进制部署版本，包含 LightBridge 品牌、Release 发布包，以及支持安装、升级、回退、从 Sub2API 自动迁移到 LightBridge 的一键脚本。

## 快速安装

```bash
curl -fsSL https://raw.githubusercontent.com/WilliamWang1721/LightBridge/main/deploy/install.sh | sudo bash
```

服务启动后，在浏览器打开初始化向导：

```text
http://服务器IP:8080
```

## 升级 LightBridge

将已有 LightBridge 二进制 systemd 部署升级到最新 Release：

```bash
curl -fsSL https://raw.githubusercontent.com/WilliamWang1721/LightBridge/main/deploy/install.sh | sudo bash -s -- upgrade
```

升级或回退到指定版本：

```bash
curl -fsSL https://raw.githubusercontent.com/WilliamWang1721/LightBridge/main/deploy/install.sh | sudo bash -s -- upgrade -v v0.0.1
```

## 从 Sub2API 迁移到 LightBridge

如果服务器上仍是旧的 Sub2API 二进制 systemd 部署，使用：

```bash
curl -fsSL https://raw.githubusercontent.com/WilliamWang1721/LightBridge/main/deploy/install.sh | sudo bash -s -- migrate -v v0.0.1
```

`migrate` 会自动检测 `sub2api.service`、`/opt/sub2api/sub2api`、`/opt/sub2api/LightBridge`、`/usr/local/bin/sub2api`，备份旧部署，复制旧配置和运行数据到 LightBridge 目录，安装 `LightBridge.service`，禁用旧 `sub2api.service`，并启动 LightBridge。

迁移备份目录：

```text
/opt/LightBridge-migration-backups/<timestamp>
```

## 常用命令

```bash
sudo systemctl status LightBridge
sudo journalctl -u LightBridge -f
sudo systemctl restart LightBridge
```

更完整的部署说明见 [deploy/README.md](deploy/README.md)。

## 友情链接

- [LinuxDO](https://linux.do/) - 一个友好的开发者社区。
