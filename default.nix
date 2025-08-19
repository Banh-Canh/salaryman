{
  pkgs ? import <nixpkgs> { },
}:
let
  cleanExamplesFilter =
    name: type:
    let
      parentDir = baseNameOf (dirOf name);
    in
    !(parentDir == "examples" && (pkgs.lib.hasSuffix ".pdf" name));

  cleanDocsFilter =
    name: type:
    let
      parentDir = baseNameOf (dirOf name);
    in
    !(parentDir == "docs" && (pkgs.lib.hasSuffix ".md" name));

  cleanSource =
    src:
    pkgs.lib.cleanSourceWith {
      filter = cleanExamplesFilter;
      src = pkgs.lib.cleanSourceWith {
        filter = cleanDocsFilter;
        src = pkgs.lib.cleanSource src;
      };
    };
  build = pkgs.buildGoModule {
    pname = "salaryman";
    version = "nix";

    src = cleanSource ./.;
    ldflags = [
      "-s"
      "-w"
      "-X github.com/Banh-Canh/salaryman/cmd.version=nix"
    ];

    vendorHash = "sha256-jck1S7M3IwQorOsIFx94F0rSI+BEy+aEy/ZyNSYH5SY=";

    subPackages = [ "." ];
  };
in
build
