# apko

```sh
melange keygen
```

```sh
dagger -m ./apko --source . --apko-file hello.apko.yaml --image ghcr.io/developer-guy/bash --keyring-append melange.rsa export --path apko.tar

docker load < apko.tar
```