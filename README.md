This repo contains a GitHub Action that deploys the requested application to either dev or prod.

Here is the startup script that got base64ed (obviously with two different values for `ENV`). A potential future improvement would be to have it stored in the repo and base64ed at runtime:
```
#!/bin/bash
export ENV="test"
export GOCACHE="/home/admin/.cache/go-build"
apt update
apt install git -y
git clone https://github.com/smistephen/opslevel-deploy.git
cd opslevel-deploy/
apt install golang -y
TOKEN=`curl -X PUT "http://169.254.169.254/latest/api/token" -H "X-aws-ec2-metadata-token-ttl-seconds: 21600"` && HOSTNAME=`curl -H "X-aws-ec2-metadata-token: $TOKEN" http://169.254.169.254/latest/meta-data/public-hostname`
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout keyfile.key -out crtfile.crt -subj "/C=CA/ST=NS/L=Halifax/O=SteveLLC/OU=Testing/CN=$HOSTNAME"
go build -v main.go
./main
```
