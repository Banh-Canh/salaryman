{
  pkgs ? (
    import <nixpkgs> {
      config.allowUnfree = true;
    }
  ),
  ...
}:
pkgs.mkShell {
  buildInputs = [
    pkgs.google-chrome
    pkgs.go
  ];
  packages = [
    (pkgs.writeShellScriptBin "salaryman" ''
      #!/bin/bash
        $(nix-build .)/bin/salaryman "$@"
    '')
    (pkgs.writeShellScriptBin "generateExamples" ''
      #!/bin/bash
      rm -f examples/*.pdf
      salaryman local -f examples/example.json -o examples/classic.pdf -t classic
      salaryman local -f examples/example.json -o examples/basic.pdf -t basic
      salaryman local -f examples/example.json -o examples/simple.pdf -t simple
      salaryman local -f examples/example.json -o examples/oldman.pdf -t oldman
      salaryman local -f examples/example.json -o examples/stackoverflow.pdf -t stackoverflow
      salaryman local -f examples/example.json -o examples/modern.pdf -t modern
      go run ./scripts/gen_doc.go
    '')
  ];
}
