Proves that you can import individual packages of a monorepo as dependencies.

The thing about doing this, is you need to explicitly check the Gopkg.toml to ensure that the dependency constraint is there for the correct version.
Then you can run `dep ensure` to sync everything up