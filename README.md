# go

lib for errors, gorm db, utils and validators


/app # go mod init github.com/uzrnem/golang/rdb
/app # go mod tidy

git tag "v1.0.0"
git push --tags


wget  https://go.dev/dl/go1.20.2.linux-amd64.tar.gz
sudo tar -xvf go1.20.2.linux-amd64.tar.gz
sudo mv go /usr/local
export GOROOT=/usr/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
