{
  pkgs ? (
    import <nixpkgs> {
      config.allowUnfree = true;
    }
  ),
  ...
}:
let
  browsers = pkgs.playwright-driver.browsers;
  headlessShellDir = builtins.head (
    builtins.filter (e: pkgs.lib.hasPrefix "chromium_headless_shell-" e) (
      builtins.attrNames (builtins.readDir browsers)
    )
  );
in
pkgs.mkShell {
  buildInputs = [
    browsers
    pkgs.go
  ];
  CHROME_PATH = "${browsers}/${headlessShellDir}/chrome-headless-shell-linux64/chrome-headless-shell";
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
