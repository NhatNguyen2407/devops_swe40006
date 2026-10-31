# Retro Minicon Lab - SWE40006

This is my custom Minicon image exercise, based on the pixel/retro interface from my earlier AWS workshop. It is separate from the original ECS/Fargate setup.

## Components

- `index.html` - updated retro interface for this lab
- `main.go` - Go HTTP server that embeds the HTML and logo
- `AWS_logo.png` - reused from the repository root at build time
- `retro-server` - compiled static Linux ARM64 executable (generated)
- `bibii-retro.img` - self-contained Minicon image (generated)

## Build and run on Amazon Linux 2023 ARM64

Run the following after unpacking `entest-hai/minicon` into `~/minicon-main`:

```bash
cd ~/minicon-main
sudo dnf install -y golang
mkdir -p retro-app
curl -fL https://raw.githubusercontent.com/NhatNguyen2407/devops_swe40006/main/minicon-retro/main.go -o retro-app/main.go
curl -fL https://raw.githubusercontent.com/NhatNguyen2407/devops_swe40006/main/minicon-retro/index.html -o retro-app/index.html
curl -fL https://raw.githubusercontent.com/NhatNguyen2407/devops_swe40006/main/AWS_logo.png -o retro-app/AWS_logo.png
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o retro-app/retro-server ./retro-app
./minicon package --name bibii-retro:1.0 --bin ./retro-app/retro-server --entrypoint /retro-server --port 8081 --out bibii-retro.img
sudo ./minicon run --image bibii-retro.img --hostname bibii-retro
```

Open a second EC2 session and run:

```bash
curl -I http://127.0.0.1:8081
curl -s http://127.0.0.1:8081 | grep 'Minicon'
tar -tzf bibii-retro.img | head
tar -xOzf bibii-retro.img manifest.json
```

The server port is `8081` to avoid interfering with the lecturer's demo on `8080`. For browser access, allow inbound TCP/8081 from **My IP** only. The source remains available on GitHub and the custom image can be copied to another compatible ARM64 Linux host without Go or Docker.

## Evidence

- Source code and changed UI compared with the earlier AWS workshop
- Successful Go build and `file retro-app/retro-server`
- Successful `minicon package` and manifest for `bibii-retro:1.0`
- Successful `minicon run` and local HTTP check
- Browser rendering at EC2 public IP on port 8081

This teaching runtime uses Linux namespaces and chroot; it is not a hardened production container runtime.
