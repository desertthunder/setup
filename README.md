# Dotfiles 2.0 + CLI (`thunderize`)

A personal dotfiles and development environment management tool designed for macOS and Arch btw

## Notes

Installing Python via asdf requires the following:

```sh
pacman -S --needed base-devel openssl zlib xz tk zstd
```

Install oh-my-zsh via the curl command on the website:

```sh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

### Rsync

Common rsync patterns used in this project:

```sh
# Directory sync (with --delete to mirror source)
rsync -av --delete ~/Projects/Personal/setup/config/hypr/ ~/.config/hypr/

rsync -av --delete ~/Projects/Personal/setup/config/nvim/ ~/.config/nvim/

rsync -av --delete ~/Projects/Personal/setup/config/alacritty/ ~/.config/alacritty/

# File sync (without --delete for single files)
rsync -av ~/Projects/Personal/setup/config/zshrc ~/.zshrc

rsync -av ~/Projects/Personal/setup/config/tool-versions ~/.tool-versions

rsync -av ~/Projects/Personal/setup/config/omp.json ~/.omp.json

# Reverse sync (backup from system to repo)
rsync -av --delete ~/.config/hypr/ ~/Projects/Personal/setup/config/hypr/

# With custom excludes
rsync -av --delete --exclude=node_modules --exclude=.cache source/ target/
```

Generic patterns using environment variables

```sh
# Deploy: repo -> XDG_CONFIG_HOME
CONFIG_NAME=hypr
rsync -av --delete ~/Projects/Personal/setup/config/$CONFIG_NAME/ \
  ${XDG_CONFIG_HOME:-~/.config}/$CONFIG_NAME/

# Backup: XDG_CONFIG_HOME -> repo
CONFIG_NAME=hypr
rsync -av --delete ${XDG_CONFIG_HOME:-~/.config}/$CONFIG_NAME/ \
  ~/Projects/Personal/setup/config/$CONFIG_NAME/
```

#### Flags

`-r`: recursive
`-v`: verbose
`-h`: human readable
`-a`: archive mode (preserves permissions, timestamps, symlinks, etc.)

## TODO

1. Thinner, lighter borders

### Waybar

1. Bluetooth Controller
2. Icons
