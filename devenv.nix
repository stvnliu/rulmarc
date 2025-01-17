{
  pkgs,
  lib,
  config,
  inputs,
  ...
}:

{
  name = "ai-game";
  # https://devenv.sh/basics/
  env = {
    GREET = "🛠️ Let's hack ";
  };

  # https://devenv.sh/scripts/
  scripts = {
    hello.exec = "echo $GREET";
    cat.exec = "bat $@";
    build.exec = "go build -o ./bin/";
    show = {
      # Prints scripts that have a description
      # Adapted from https://github.com/cachix/devenv/blob/ef61728d91ad5eb91f86cdbcc16070602e7afa16/examples/scripts/devenv.nix#L34
      exec = ''
        GREEN="\033[0;32m";
        YELLOW="\033[33m";
        NC="\033[0m";
        echo
        echo -e "✨ Helper scripts you can run to make your development richer:"
        echo
        ${pkgs.gnused}/bin/sed -e 's| |••|g' -e 's|=| |' <<EOF | ${pkgs.util-linuxMinimal}/bin/column -t | ${pkgs.gnused}/bin/sed -e "s|^\([^ ]*\)|$(printf "$GREEN")\1$(printf "$NC"):    |" -e "s|^|$(printf "$YELLOW*$NC") |" -e 's|••| |g'
        ${lib.generators.toKeyValue { } (
          lib.mapAttrs (name: value: value.description) (
            lib.filterAttrs (_: value: value.description != "") config.scripts
          )
        )}
        EOF
        echo
      '';
      description = "Print this message and exit.";
    };

  };

  # https://devenv.sh/packages/
  packages = with pkgs; [
    nixfmt-rfc-style
    bat
    jq
    tealdeer
    gopls
    air
    ncurses
  ];

  languages = {
    go.enable = true;

  };

  enterShell = ''
    hello
    show'';

  # https://devenv.sh/pre-commit-hooks/
  pre-commit.hooks = {
    nixfmt-rfc-style = {
      enable = true;
      excludes = [ ".devenv.flake.nix" ];
    };
    yamllint = {
      enable = true;
      settings.preset = "relaxed";
    };
  };

  # Make diffs fantastic
  difftastic.enable = true;

  # https://devenv.sh/integrations/dotenv/
  dotenv.enable = true;

  # https://devenv.sh/integrations/codespaces-devcontainer/
  devcontainer.enable = true;
}
