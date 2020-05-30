# MONOREPO

## install protoc

```sh
VERSION=3.12.2
PROTOC_ZIP=protoc-$VERSION-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$VERSION/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP
```
