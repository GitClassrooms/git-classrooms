{
  description = "Development enviroment for git classrooms";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    mockery.url = "github:nixos/nixpkgs/92d295f588631b0db2da509f381b4fb1e74173c5";
    pre-commit-hooks.url = "github:cachix/git-hooks.nix";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    ...
  } @ inputs: (flake-utils.lib.eachDefaultSystem
    (system: let
      pkgs = nixpkgs.legacyPackages.${system};
      pre-commit-check = inputs.pre-commit-hooks.lib.${system}.run {
        src = ./.;
        hooks = {
          golangci-lint.enable = false;
          gofmt.enable = true;
          gotest.enable = false;
        };
      };

      concur = pkgs.buildGoModule {
        pname = "concur";
        version = "v0.1.0";

        src = pkgs.fetchFromGitHub {
          owner = "akatranlp";
          repo = "concur";
          rev = "7ebd166c4d2b8642a53efa336b9192079bd8d67c";
          sha256 = "sha256-A56wANrOzU16eb1th1TnalzcVGeLjz5KOS9Xv2+xZfQ=";
        };
        vendorHash = "sha256-InM4nDMCBNGLy4INGGBR896YZgCchqF1/pBGQMK1ruc=";
      };
    in {
      devShells.default = pkgs.mkShell {
        nativeBuildInputs = with pkgs; [
          # Frontend
          nodejs_22
          yarn

          # Backend
          inputs.mockery.legacyPackages.${system}.go-mockery
          go_1_23
          air
          go-swag
          golangci-lint
          goose
          delve
          concur

          yq-go
          docker
          docker-compose
          gnumake
          git
        ];

        shellHook = ''
          cd ./frontend && yarn install

          cd ..
          go mod download
          ${pre-commit-check.shellHook}
        '';
      };
    }));
}
