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
  '';

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

}
