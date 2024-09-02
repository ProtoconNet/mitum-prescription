### mitum-prescription

*mitum-prescription* is a prescription contract model based on the second version of mitum(aka [mitum2](https://github.com/ProtoconNet/mitum2)).

#### Installation

Before you build `mitum-prescription`, make sure to run `docker run` for digest api.

```sh
$ git clone https://github.com/ProtoconNet/mitum-prescription

$ cd mitum-prescription

$ go build -o ./mitum-prescription
```

#### Run

```sh
$ ./mitum-prescription init --design=<config file> <genesis file>

$ ./mitum-prescription run <config file> --dev.allow-consensus
```

[standalong.yml](standalone.yml) is a sample of `config file`.

[genesis-design.yml](genesis-design.yml) is a sample of `genesis design file`.
