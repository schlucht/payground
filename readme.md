# Übungsrepo

Diverse Übungen 

- docker run -it -v ${pwd}/workspaces/payground/:/root/.config nvim bash

docker run -w /root -it --rm alpine:edge sh -uelic '
  apk add git lazygit neovim ripgrep alpine-sdk --update
  git clone https://github.com/LazyVim/starter ~/.config/nvim
  cd ~/.config/nvim
  nvim
'
### on Linux and Mac
- git clone https://github.com/nvim-lua/kickstart.nvim.git "${XDG_CONFIG_HOME:-$HOME/.config}"/nvim
- nvim --headless "+Lazy! sync" +qa