#!/bin/bash

sudo rm -rf /usr/local/go
sudo tar -xzvf go1.17.11.linux-amd64.tar.gz -C /usr/local/

echo "GOPATH=$HOME/go" >> $HOME/.bashrc
echo "export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin" >> $HOME/.bashrc
echo "export GOROOT=/usr/local/go" >> $HOME/.bashrc
echo "export GOPROXY=https://goproxy.cn" >> $HOME/.bashrc
source $HOME/.bashrc

go version