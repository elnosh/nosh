# nosh

nostr command-line tool for generating events

## Documentation

```
USAGE:
   nosh [global options] command [command options] [arguments...]

COMMANDS:
   key      generate, set private key
   event    generate nostr event
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## How to use

* generate key
```
$ nosh key --gen
private key: 4ecf33b88e4bf1f9da31dbca11248833ac90c48ffe6f92d57d97d4a837caab04
```

* set key
```
$ nosh key --set 4ecf33b88e4bf1f9da31dbca11248833ac90c48ffe6f92d57d97d4a837caab04
```

* generate event
   ```
   $ nosh event -c "hi nostr" -pk 4ecf33b88e4bf1f9da31dbca11248833ac90c48ffe6f92d57d97d4a837caab04 
   {"id":"e21a6a1f4ef383277f7ce0ad419d8034b86f5ef4483df73f6ab02a62d12ae644","pubkey":"6ca4366e27fc89c8d69a758aa1cc617e8d2333cd7ed9364132f165f9099e91a8","created_at":1684416218,"kind":1,"tags":[],"content":"hi nostr","sig":"e2bdcb7b1eb6b089773b24ffc4090dd23fc262db4368a475d47e3f8dcb9ee7cedfe0d250b9fb8e68f076be1b7c7fc8fe12776816a1ec9681a45d6df233e38c66"}
   ```

   * if key is set, event can be generated without the "pk" flag
   ```
   $ nosh event -c "test message"
   {"id":"9d67074fc29ecb875b1c41011ca4e6c70a012297694c6e2e6f75c6058c834084","pubkey":"6ca4366e27fc89c8d69a758aa1cc617e8d2333cd7ed9364132f165f9099e91a8","created_at":1684416326,"kind":1,"tags":[],"content":"test message","sig":"96a19acf43d599cc1badc5e670931f6840eb68b4f08a0017313eb7e917c90786e6e273399e8d786e4bbe1e1b9a9a8251f388635e29b310becffc32993cc6a6fc"}
   ```

* wrap event ["EVENT", < json event >]
```
$ nosh event -c "nostr is cool" --wrap
["EVENT", {"id":"5622d96bd48b2b227c091967321aa893bd71866c8c657ab713ddd6f1b55b896f","pubkey":"6ca4366e27fc89c8d69a758aa1cc617e8d2333cd7ed9364132f165f9099e91a8","created_at":1684416626,"kind":1,"tags":[],"content":"nostr is cool","sig":"91b4f6f520c7c91c365aaa7b4a6ab9a7b129fbf93e910f2b44167a62909fc8e7f5bea5d6e290affe76e6fc3fb433fda79ae8a3eee9aa8e0bc878e508d1430dfe"}]
```

[nak](https://github.com/fiatjaf/nak) and [nostril](https://github.com/jb55/nostril) are better and done by cooler people.