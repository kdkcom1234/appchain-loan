# loan

**loan** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project.

For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install

To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/loan@latest! | sudo bash
```

`username/loan` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

---

## Link

```
https://docs.ignite.com/guide/loan/intro
```

## Scaffolding a Chain

```shell
ignite scaffold chain loan --no-module  --address-prefix loan
```

## Set Default Denom

```yml
validation: sovereign
version: 1
accounts:
  - name: alice
    coins:
      - 200000000uloan
    mnemonic: umbrella mushroom luggage metal traffic slight faculty volcano mansion desert unaware sniff label artwork ozone final thunder crowd all suit wool protect ring describe
  - name: bob
    coins:
      - 100000000uloan
    mnemonic: loud donkey use home resist clinic base feature figure erode birth slogan draw relief joke weasel rough evolve mind feature hat basket march another
faucet:
  name: bob
  coins:
    - 100000uloan
client:
  typescript:
    path: ts-client
  openapi:
    path: docs/static/openapi.yml
validators:
  - name: alice
    bonded: 100000000uloan
genesis:
  app_state:
    staking:
      params:
        bond_denom: uloan
    mint:
      params:
        mint_denom: uloan
```

### Connect Keplr

```ts
// set denom
await client.useKeplr({
  stakeCurrency: {
    coinDenom: "LOAN",
    coinMinimalDenom: "uloan",
    coinDecimals: 6,
  },
  currencies: [
    { coinDenom: "LOAN", coinMinimalDenom: "uloan", coinDecimals: 6 },
  ],
  feeCurrencies: [
    { coinDenom: "LOAN", coinMinimalDenom: "uloan", coinDecimals: 6 },
  ],
});
```

### Send Coin

```shell
loand tx bank send alice loan1hgetjlxqkv5dssgv6z4xzxtqfpfu0gm779xyry 1000000uloan --chain-id loan
```
