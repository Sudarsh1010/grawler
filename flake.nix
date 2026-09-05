{
  description = "";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      nixpkgs,
      flake-utils,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            # Go services
            go_1_27
            gopls
            golangci-lint
            gotools
            delve

            # # Contract-seam tests run against a real local Postgres cluster
            # (postgresql_17.withPackages (p: [ p.pgvector ]))

            # # Helm chart validation (helm lint / helm template)
            # kubernetes-helm

            curl
            pkg-config
          ];

          shellHook = "";
        };
      }
    );
}
