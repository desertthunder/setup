# Neovim Configuration Reference

Quick reference for keybindings, plugins, and commands. See [GUIDE.md](GUIDE.md) for detailed explanations and tutorials.

## Directory Structure

```sh
config/nvim/
├── init.lua              # Entry point
├── lua/
│   ├── config/           # Core configuration (options, keymaps, autocmds)
│   └── plugins/          # Plugin specifications
```

## Essential Vim Keybindings

### Normal Mode

**Navigation:**

- `h/j/k/l` - Left/Down/Up/Right
- `w/b` - Next/previous word
- `0/$` - Start/end of line
- `gg/G` - Start/end of file
- `{/}` - Previous/next paragraph
- `Ctrl-u/d` - Scroll half page up/down
- `Ctrl-b/f` - Scroll full page up/down
- `%` - Jump to matching bracket

**Editing:**

- `i/a` - Insert before/after cursor
- `I/A` - Insert at start/end of line
- `o/O` - New line below/above
- `x/X` - Delete char under/before cursor
- `dd` - Delete line
- `yy` - Yank (copy) line
- `p/P` - Paste after/before cursor
- `u` - Undo
- `Ctrl-r` - Redo
- `.` - Repeat last command

**Visual Mode:**

- `v` - Visual character mode
- `V` - Visual line mode
- `Ctrl-v` - Visual block mode

**Search:**

- `/pattern` - Search forward
- `?pattern` - Search backward
- `n/N` - Next/previous match
- `*/#` - Search word under cursor forward/backward

**Files:**

- `:w` - Save
- `:q` - Quit
- `:wq` or `:x` - Save and quit
- `:q!` - Quit without saving

## Custom Keybindings

### General

- `<Space>` - Leader key
- `<leader>?` - Toggle cheatsheet (shows all keymaps)
- `<Esc>` - Clear search highlights
- `<leader>q` - Open diagnostic quickfix list
- `<Esc><Esc>` - Exit terminal mode (in terminal)

### Window Navigation

- `<C-h/j/k/l>` - Move to left/down/up/right window

### LSP (when active)

- `grn` - Rename symbol
- `gra` - Code action
- `grr` - Find references
- `gri` - Go to implementation
- `grd` - Go to definition
- `grD` - Go to declaration
- `gO` - Document symbols
- `gW` - Workspace symbols
- `grt` - Type definition
- `<leader>th` - Toggle inlay hints

### Telescope (Fuzzy Finder)

- `<leader>sh` - Search help
- `<leader>sk` - Search keymaps
- `<leader>sf` - Search files
- `<leader>ss` - Search Telescope pickers
- `<leader>sw` - Search current word
- `<leader>sg` - Live grep
- `<leader>sd` - Search diagnostics
- `<leader>sr` - Resume last search
- `<leader>s.` - Recent files
- `<leader><leader>` - Find buffers
- `<leader>/` - Fuzzy search in current buffer
- `<leader>s/` - Live grep in open files
- `<leader>sn` - Search Neovim config files

**In Telescope picker:**

- `<C-/>` (insert mode) or `?` (normal mode) - Show help

### Completion (Blink.cmp)

- `<c-space>` - Open completion menu / show docs
- `<c-y>` - Accept completion
- `<c-n>/<c-p>` or `<up>/<down>` - Select next/previous
- `<c-e>` - Hide menu
- `<c-k>` - Toggle signature help
- `<tab>/<s-tab>` - Navigate snippet placeholders

### Formatting

- `<leader>f` - Format buffer

### Mini.nvim Textobjects

- `va)` - Visual select around paren
- `yinq` - Yank inside next quote
- `ci'` - Change inside quote

### Mini.nvim Surround

- `saiw)` - Surround add inner word with paren
- `sd'` - Surround delete quotes
- `sr)'` - Surround replace paren with quote

### Git (Gitsigns)

- `<leader>h` group - Git hunk operations (shown by which-key)

## Installed Plugins

| Plugin                 | Description                                            |
| ---------------------- | ------------------------------------------------------ |
| **lazy.nvim**          | Plugin manager                                         |
| **telescope.nvim**     | Fuzzy finder for files, buffers, LSP, etc.             |
| **which-key.nvim**     | Shows pending keybindings                              |
| **nvim-lspconfig**     | LSP configuration                                      |
| **mason.nvim**         | LSP/DAP/linter installer                               |
| **blink.cmp**          | Autocompletion engine                                  |
| **LuaSnip**            | Snippet engine                                         |
| **conform.nvim**       | Code formatting (format on save)                       |
| **nvim-treesitter**    | Syntax highlighting and code understanding             |
| **gitsigns.nvim**      | Git decorations and commands                           |
| **mini.nvim**          | Collection of small plugins (statusline, surround, ai) |
| **tokyonight.nvim**    | Colorscheme                                            |
| **todo-comments.nvim** | Highlight TODO/FIXME/etc in comments                   |
| **lazydev.nvim**       | Lua LSP for Neovim config                              |
| **guess-indent.nvim**  | Auto-detect indentation                                |
| **cheatsheet.nvim**    | Custom plugin - displays keymaps in floating window    |

## Quick Commands

### Plugin Management

- `:Lazy` - Open plugin manager
- `:Lazy update` - Update all plugins
- `:Lazy sync` - Install missing and update plugins
- `:Lazy clean` - Remove unused plugins

### LSP

- `:LspInfo` - Show LSP client info
- `:Mason` - Open Mason package manager
- `:ConformInfo` - Check formatter status

### Treesitter

- `:TSUpdate` - Update parsers
- `:TSInstall <language>` - Install language parser

### Diagnostics

- `:checkhealth` - Check Neovim health
- `:checkhealth lazy` - Check plugin manager
- `:checkhealth lsp` - Check LSP

### Help

- `:help` - Open help
- `:help <topic>` - Search help for topic
- `:Tutor` - Interactive Neovim tutorial

## Configuration

### Add Language Server

Edit `lua/plugins/lspconfig.lua`, add to `servers` table:

```lua
local servers = {
  lua_ls = { ... },
  gopls = {},     -- Add Go
  pyright = {},   -- Add Python
}
```

### Add Formatter

Edit `lua/plugins/conform.lua`, add to `formatters_by_ft`:

```lua
formatters_by_ft = {
  lua = { 'stylua' },
  python = { 'black' },
  go = { 'gofmt' },
},
```

### Change Colorscheme

Edit `lua/plugins/colorscheme.lua`:

```lua
vim.cmd.colorscheme 'tokyonight-storm'  -- or 'tokyonight-moon', 'tokyonight-day'
```

### Add Plugin

Create `lua/plugins/<name>.lua`:

```lua
return {
  'author/plugin-name',
  opts = {},
}
```

## Troubleshooting

| Issue                  | Solution                                            |
| ---------------------- | --------------------------------------------------- |
| Plugins not loading    | `:Lazy sync` then `:checkhealth lazy`               |
| LSP not working        | `:LspInfo` then `:Mason` then `:checkhealth lsp`    |
| Formatting not working | `:ConformInfo` then check `lua/plugins/conform.lua` |
| Treesitter errors      | `:TSUpdate` then restart Neovim                     |

## Resources

- `:help` - Built-in Neovim documentation
- `:help lua-guide` - Lua in Neovim guide
- [GUIDE.md](GUIDE.md) - Detailed configuration guide
- [Kickstart.nvim](https://github.com/nvim-lua/kickstart.nvim) - Original project
