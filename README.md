Ideally, the GitHub Action in this repo would be a one-click deploy of the application.

But there seems to be some odd discrepancy between the UI and CLI, because providing the following `user-data` via the UI works, but not via the CLI.

My best examination of the error logs shows that it receives an authentication error when it attempts to clone the git repo. Reviewing the script shows it seems to have been communicated to the server intact, so I'm unsure where the discrepancy is.

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
