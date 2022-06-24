# valorant-exporter

A valorant elo exporter written in golang.

## Available metrics

Name     | Description
---------|------------
valorant_elo | Elo of a player
valorant_games | Number of games played
valorant_tier | Tier of a player
valorant_wins | Number of games won

All metrics include the following labels:

* username
* tagline

## How to execute for development purposes

### Nix / NixOS

This repository contains a `flake.nix` file.

```sh
# run the package
nix run .#valorant_exporter

# build the package
nix build .#valorant_exporter
```

### General

Make sure [golang](https://go.dev) is installed.

```sh
# run application
go run .

# build application
go build
```

## How to install

### NixOS

1. Add this repository to your `flake.nix`:

```nix
{
  inputs.valorant-exporter = {
    url = "github:MayNiklas/valorant-exporter";
    inputs = { nixpkgs.follows = "nixpkgs"; };
  };
}
```

2. Enable the service & the prometheus scraper in your configuration:

```nix
{ valorant-exporter, ... }: {

  imports = [ valorant-exporter.nixosModules.default ];

  services.shelly-exporter = {
    enable = true;
    port = "1091";
    listen = "localhost";
    user = "valorant-exporter";
    group = "valorant-exporter";

    configure-prometheus = true;
    targets = [
      "mayniklas/niki"
    ];
  };
}
```

### Libaries used

* https://github.com/prometheus/client_golang

### Libary documentation

* https://pkg.go.dev/github.com/prometheus/client_golang/prometheus
