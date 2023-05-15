# nosh

nostr command-line tool

## Documentation

```
NAME:
   nosh - nostr CLI tool

USAGE:
   nosh [global options] command [command options] [arguments...]

COMMANDS:
   genkey   generate private key
   event    generate nostr events
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## How to use

* generate key
```
$ nosh genkey
private key: 23924aa0f9b1b21e8d725976794ef037c2afe9ffe56a186e7e79a97e9983413d
```

* generate event
```
$ nosh event -c "hi nostr" -k 1 -pk 23924aa0f9b1b21e8d725976794ef037c2afe9ffe56a186e7e79a97e9983413d

{"id":"d453c9931e161c1288dac57e7e0dccf222987a8e53ac9c7b17698e1025f387b5","pubkey":"87ddf25104407f89733fe58cbf1fd2f92a69462180012a2fcc4d6da77678cf53","created_at":1684120718,"kind":1,"tags":[],"content":"hi nostr","sig":"58a1d25e9b1ed02064585214bb9100ac7a193ec68bd714670e44671b22b9b0e5782fe3854ddde581120d5d9926a72b08254294babc16d1cc9a9315e2219418c6"}
```

* wrap event ["EVENT", < json event >]
```
$ nosh event -c "hi nostr" -k 1 -w -pk 23924aa0f9b1b21e8d725976794ef037c2afe9ffe56a186e7e79a97e9983413d

["EVENT", {"id":"2cee356b7ff42088b82dad86951d507985f8e5784b721f2c46a36a02aff6f6ac","pubkey":"87ddf25104407f89733fe58cbf1fd2f92a69462180012a2fcc4d6da77678cf53","created_at":1684120983,"kind":1,"tags":[],"content":"hi nostr","sig":"a5c6dc8258454fc27764de68848d5353a4aa9dea1f4d0a094ada93c7f047f7d76eb87604e38102c772f3eef0344011d45ef81ce35a05bff5d4e9842bc47bc101"}]

```

[nak](https://github.com/fiatjaf/nak) and [nostril](https://github.com/jb55/nostril) are better and done by cooler people.