FROM ubuntu:latest

RUN apt-get update && apt-get install -y wget git curl zsh unzip git golang-go cmake g++
RUN wget https://github.com/neovim/neovim/releases/download/stable/nvim-linux64.tar.gz \
    && tar zxfv nvim-linux64.tar.gz
RUN git clone --depth 1 https://github.com/wbthomason/packer.nvim\
     ~/.local/share/nvim/site/pack/packer/start/packer.nvim
RUN sh -c "$(wget https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" - \
    && sh install.sh
RUN wget -P ~/.local/share/fonts https://github.com/ryanoasis/nerd-fonts/releases/download/v3.0.2/JetBrainsMono.zip \
    && unzip ~/.local/share/fonts/JetBrainsMono.zip -d ~/.local/share/fonts \
    && rm ~/.local/share/fonts/JetBrainsMono.zip

COPY /config/ /root/



