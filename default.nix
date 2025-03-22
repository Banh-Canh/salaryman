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
  cleanFileFilter =
    name: type:
    !(
      (type == "file")
      && ((baseNameOf (dirOf name)) == "shell.nix")
      && ((baseNameOf (dirOf name)) == "README.md")
      && ((baseNameOf (dirOf name)) == "LICENSE")
    );

  # Clean cached Node.js artifacts, and apply Nix's default clean routine.
  cleanSource =
    src:
    pkgs.lib.cleanSourceWith {
      filter = cleanExamplesFilter;
      src = pkgs.lib.cleanSourceWith {
        filter = cleanFileFilter;
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
      "-X github.com/Banh-Canh/salaryman/cmd/version.version=nix"
    ];

    vendorHash = "sha256-nsZB/JVJ0GwMG7Ok8XwA/O/X//PI9n0qASRNOip3fhI=";

    subPackages = [ "." ];
  };
in
build
