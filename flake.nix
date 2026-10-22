{
  inputs = {
    flake-utils = {
      url = "github:numtide/flake-utils";
    };
    nixpkgs = {
      url = "github:NixOS/nixpkgs/nixos-26.05-small";
    };
  };

  outputs =
    {
      self,
      flake-utils,
      nixpkgs,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
        version = pkgs.lib.fileContents ./internal/index/version.txt;
        # Go 1.27.3 rejects chains through the compromised eMudhra
        # intermediates (CVE-2026-111943). The default pkgs.go is an
        # older 1.26 release without the fix.
        go = pkgs.go_1_27;
        # NSS 3.131.1 removes the emSign ECC Root CA - G3 root that
        # the compromised eMudhra intermediates chain to. Until nixpkgs
        # ships it, drop the root ourselves so git and other OpenSSL
        # clients reject those chains too. (The bundle already excludes
        # AAA Certificate Services, the root of the Sectigo cross-signs,
        # from TLS trust.)
        cacert =
          if pkgs.lib.versionOlder pkgs.cacert.version "3.131.1" then
            pkgs.cacert.override { blacklist = [ "emSign ECC Root CA - G3" ]; }
          else
            pkgs.cacert;
      in
      {
        packages = {
          scip-go = (pkgs.buildGoModule.override { inherit go; }) {
            pname = "scip-go";
            inherit version;
            src = ./.;
            vendorHash = "sha256-dwcZEiNSi9NsfbR7LfudptvS32TrRcm51ny16rTyE+A=";
            subPackages = [ "cmd/scip-go" ];
            env.CGO_ENABLED = 0;
            checkPhase = "go test ./...";
            # Use proxyVendor so deps go into the module cache instead
            # of a vendor/ directory that conflicts with test sub-modules
            # having their own go.mod.
            proxyVendor = true;
            nativeCheckInputs = [ pkgs.git ];
          };
          default = self.packages.${system}.scip-go;
        }
        // pkgs.lib.optionalAttrs pkgs.stdenv.isLinux {
          docker = pkgs.dockerTools.buildLayeredImage {
            name = "scip-go";
            tag = version;
            contents = [
              self.packages.${system}.scip-go
              go
              pkgs.git
              (pkgs.dockerTools.override { inherit cacert; }).caCertificates
              pkgs.dockerTools.usrBinEnv
              pkgs.dockerTools.fakeNss
              pkgs.busybox
            ];
            fakeRootCommands = ''
              mkdir -p /work /tmp
            '';
            enableFakechroot = true;
            config = {
              Cmd = [ "scip-go" ];
              WorkingDir = "/work";
              Env = [
                "GOTOOLCHAIN=auto"
                "HOME=/tmp"
                "TMPDIR=/tmp"
                "GOCACHE=/tmp/go-build"
                "GOMODCACHE=/tmp/go-mod"
              ];
              Labels = {
                "org.opencontainers.image.source" = "https://github.com/scip-code/scip-go";
              };
            };
          };
        };

        checks = {
          nixfmt = pkgs.runCommand "check-nixfmt" { } ''
            ${pkgs.nixfmt}/bin/nixfmt --check ${./flake.nix}
            touch $out
          '';
          goimports = pkgs.runCommand "check-goimports" { } ''
            cd ${./.}
            bad=$(
              # Snapshot outputs are generated with modified
              # indentation for alignment with annotations.
              find . -name '*.go' \
                -not -path '*/testdata/snapshots/output/*' \
                -exec ${pkgs.gotools}/bin/goimports -l {} +
            )
            if [ -n "$bad" ]; then
              echo "goimports check failed on:"
              echo "$bad"
              exit 1
            fi
            touch $out
          '';
          renovate = pkgs.runCommand "check-renovate" { } ''
            LOG_LEVEL=warn ${pkgs.renovate}/bin/renovate-config-validator \
              ${./.github/renovate.json}
            touch $out
          '';
        };

        devShells = {
          default = pkgs.mkShellNoCC {
            buildInputs = [
              go
              pkgs.gotools
              pkgs.nixfmt
            ];
          };
        };
      }
    );
}
