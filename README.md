# Dotfiles 2.0 + CLI (`thunderize`)

A personal dotfiles and development environment management [tool](#cli) (`thunderize`) designed for macOS and Arch[^1] btw.

## Goals

1. Persist my stuff
2. Act as a learning tool for myself & anyone wanting to configure their environment. We should
[read the manual](https://www.youtube.com/watch?v=rysgxl35EGc) but irl examples are nice too.

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

Common rsync patterns used in this [project](./main.go):

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
CONFIG_NAME=hypr

# Deploy: repo -> XDG_CONFIG_HOME
rsync -av --delete ~/Projects/Personal/setup/config/$CONFIG_NAME/ ${XDG_CONFIG_HOME:-~/.config}/$CONFIG_NAME/

# Backup: XDG_CONFIG_HOME -> repo
rsync -av --delete ${XDG_CONFIG_HOME:-~/.config}/$CONFIG_NAME/ ~/Projects/Personal/setup/config/$CONFIG_NAME/
```

#### Flags

`-r`: recursive
`-v`: verbose
`-h`: human readable
`-a`: archive mode (preserves permissions, timestamps, symlinks, etc.)

## Inspiration

- Folke's [Dot](https://github.com/folke/dot)files
- [Omarchy](https://github.com/basecamp/omarchy) (DHH's opinionated arch installation)
- My favorite [subreddit](https://reddit.com/r/unixporn)

## Components

| Program   | Configuration                                  |
| --------- | ---------------------------------------------- |
| Zsh       | [zshrc](config/zshrc) & [omp](config/omp.json) |
| Alacritty | [link](config/alacritty/)                      |
| Hyprland  | [link](config/hypr/)                           |
| Waybar    | [link](config/waybar/)                         |
| Rofi      | [link](config/rofi/)                           |
| Zathura   | [link](config/zathura/zathurarc)               |
| asdf      | [link](config/tool-versions)                   |
| Neovim    | [details](#neovim) & [link](config/nvim/)      |

### CLI

- `thunderize install pacman` - Install official repo packages
- `thunderize install aur` - Install AUR packages
- `thunderize install dev` - Install development tools via asdf
- `thunderize install all` - Install all packages
- `thunderize config deploy [name|all]` - Deploy configurations to system
- `thunderize config backup [name|all]` - Backup configurations from system
- `thunderize config list` - List available configurations
- `thunderize config validate` - Validate configuration files
- `thunderize setup` - Run full system setup
- `thunderize check` - Run system checks
- `thunderize secrets init` - Initialize secrets from template

Built with [urfave/cli](https://github.com/urfave/cli) and [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss)

### Neovim

#### Structure

```sh
config/nvim/
├─ lua/
│  ├── plugins/
│  │   └─ init.lua        # Plugin entry point with inlined shorter & local definitions
│  ├─ config/             # "Core" configuration (keymaps, options, etc.)
│  ├─ cheatsheet/         # Cheatsheet plugin
│  ├─ iced-lightning/     # Custom theme (Iceberg port)
│  └─ dracula-recharged/  # Custom Dracula variant
└─ colors/                # Color scheme files
```

#### Plugins

- Dashboard
  - `goolord/alpha-nvim` with custom startup screen
- Themes
  - `tokyonight.nvim`
  - `nightfox.nvim`
  - [`dracula-recharged`](config/nvim/lua/dracula-recharged/README.md) (custom)
  - [`iced-lightning`](config/nvim/lua/iced-lightning/README.md) (custom Iceberg port with light/dark variants)
- UI
  - `akinsho/bufferline.nvim` - Buffer tabs
  - `echasnovski/mini.statusline` - Status line
  - `lukas-reineke/indent-blankline.nvim` - Indentation guides
- Editing
  - `windwp/nvim-autopairs` - Auto-close brackets
  - `echasnovski/mini.surround` - Surround operations
  - `echasnovski/mini.ai` - Extended text objects
  - `NMAC427/guess-indent.nvim` - Auto-detect indentation
- Development
  - `folke/lazydev.nvim` - Lua LSP for Neovim config
  - `folke/todo-comments.nvim` - TODO/FIXME highlighting
  - `catgoose/nvim-colorizer.lua` - Color code preview

### Wallpaper

See hyprpaper [conf](config/hypr/hyprpaper.conf)

![wallpaper](config/default-paper.png)

## TODO

1. Thinner, lighter borders

### Waybar

1. Bluetooth Controller
2. Icons

[^1]: <https://endeavouros.com/>
