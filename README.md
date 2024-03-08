# daggerverse


## grype

```sh
dagger call -m ./grype scan --image-ref alpine:latest
```

## melange

```sh
dagger call -m ./melange build --melange-file=./melange/hello.melange.yaml --workspace-dir=./
```

## apko

```sh
melange keygen # gen key

dagger -m ./apko --source . --apko-file hello.apko.yaml --image ghcr.io/developer-guy/bash --keyring-append melange.rsa export --path apko.tar

docker load < apko.tar
```

## License

[MIT](./LICENSE)