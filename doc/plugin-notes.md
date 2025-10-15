# Cheatsheet Plugin - Learning Notes

Custom Neovim plugin demonstrating core plugin development concepts.

```sh
lua/cheatsheet/
├── init.lua      - Main entry point, coordinates components
├── config.lua    - Configuration management
├── keymaps.lua   - Keymap extraction and organization
└── ui.lua        - Floating window and rendering
```

## Core Concepts

### Module System

Neovim loads Lua modules via `require()`. Modules return tables with functions (the `M` pattern). The main entry point (`init.lua`) coordinates between other modules, managing state and orchestrating API calls.

```lua
local M = {}
function M.setup(opts) end
return M
```

**State Management**: Track plugin state (like window/buffer handles) to enable toggling and proper cleanup. State typically stored in module-local tables.

### Configuration Management

Plugins provide `setup()` functions to merge user options with defaults. Standard pattern:

1. Define default configuration in the config module
2. Allow users to override via `setup()` options
3. Provide accessor to get merged config

```lua
M.current = vim.tbl_deep_extend('force', M.defaults, user_opts or {})
```

The `force` mode means user options override defaults. Use `vim.tbl_deep_extend()` for recursive merging of nested tables.

### Keymap Introspection

Extract keymaps from Neovim using `vim.api.nvim_get_keymap(mode)`. Returns array of keymap tables with fields: `lhs` (left-hand side key), `rhs` (right-hand side command), `desc` (description).

```lua
local raw_maps = vim.api.nvim_get_keymap(mode)
-- Returns: { lhs, rhs, desc, ... }
```

**Mode extraction**: Common modes are `n` (normal), `v` (visual), `i` (insert), `t` (terminal). The `desc` field is typically only set for user-defined keymaps, making it useful for filtering.

**Grouping strategies**:

- Common prefixes (`<leader>s*` → Search, `<leader>t*` → Toggle)
- Description keywords (`git`, `lsp`, `buffer`, `window`)
- Mode fallback (use mode name as category if no pattern matches)

**Filtering**: Exclude internal mappings by pattern matching (e.g., `<Plug>`, `<SNR>`) or empty descriptions.

### Buffer Management

Buffers store text content. Scratch buffers are temporary and won't prompt to save.

```lua
-- nvim_create_buf(listed, scratch)
-- listed: false = don't show in buffer list
-- scratch: true = deleted when hidden
local buf = vim.api.nvim_create_buf(false, true)
vim.bo[buf].buftype = 'nofile'    -- not associated with file
vim.bo[buf].filetype = 'cheatsheet'
vim.bo[buf].bufhidden = 'wipe'
vim.api.nvim_buf_set_lines(buf, 0, -1, false, lines)
```

**Writing to buffers**: `nvim_buf_set_lines(buffer, start, end, strict_indexing, lines)` where start=0, end=-1 replaces all lines. Set `modifiable=true` before writing, `false` after to make read-only.

### Window Management

Floating windows display buffers with custom positioning. Calculate size as percentage of editor dimensions, then center:

```lua
local opts = {
  relative = 'editor',  -- Position relative to editor
  width = win_width,
  height = win_height,
  row = row,           -- Top-left corner Y (calculated for centering)
  col = col,           -- Top-left corner X (calculated for centering)
  style = 'minimal',   -- No number column, sign column, etc.
  border = 'rounded',
  title = ' Title ',
  title_pos = 'center',
}
-- nvim_open_win(buffer, enter, config)
-- enter: true = move cursor to window
vim.api.nvim_open_win(buf, true, opts)
```

**Coordinate system**: (0,0) is top-left of editor. Centering formula: `(screen_dimension - window_dimension) / 2`

**Window lifecycle**: Toggle by tracking window handle and checking validity with `nvim_win_is_valid()`. Close with `nvim_win_close()` for proper cleanup.

### Syntax Highlighting

Apply highlights using namespaces to prevent conflicts between plugins:

```lua
local ns = vim.api.nvim_create_namespace('cheatsheet')
vim.api.nvim_buf_clear_namespace(buf, ns, 0, -1)  -- Clear old highlights
vim.api.nvim_buf_add_highlight(buf, ns, 'Title', line_idx, col_start, col_end)
```

Highlights are applied line-by-line after rendering. Use pattern matching on line content to determine which highlight group to apply. Set `col_end = -1` to highlight entire line from `col_start`.

### User Interaction

Buffer-local keymaps restrict bindings to specific buffer, preventing global conflicts:

```lua
vim.keymap.set('n', 'q', close_fn, { buffer = buf })
```

Autocommands with buffer scope enable automatic cleanup:

```lua
vim.api.nvim_create_autocmd('BufLeave', {
  buffer = buf,
  callback = close_fn,
  once = true,  -- Only trigger once, then remove
})
```

**Proper cleanup**: Always validate handles with `nvim_win_is_valid()` and `nvim_buf_is_valid()` before operations. Use `force = true` when deleting buffers to skip confirmation.

### Plugin Integration with Lazy.nvim

Plugin specification structure:

```lua
return {
  name = 'cheatsheet.nvim',
  dir = vim.fn.stdpath('config') .. '/lua/cheatsheet',  -- Local plugin
  keys = { { '<leader>?', desc = 'Toggle Cheatsheet' } },
  config = function()
    require('cheatsheet').setup({ ... })
  end,
}
```

For external plugins: `[1] = 'author/repo-name'`

## Development Phases

1. Plugin foundation - module structure, configuration
2. Data layer - keymap extraction and organization
3. UI layer - floating window creation
4. Rendering - grid layout and syntax highlighting
5. Interactivity - keymaps, autocommands, cleanup
6. Integration - user commands, documentation

## Extending the Plugin

### Add Custom Categories

Edit `keymaps.lua`, update `infer_category()`:

```lua
if key:match('^<leader>d') then
  return 'Debug'
end
```

### Add Custom Highlight Groups

Edit plugin spec, add to highlights table:

```lua
highlights = {
  header = 'Title',
  footer = 'SpecialComment',
}
```

## Relevant Neovim Help

- `:help api-buffer` - Buffer API
- `:help api-window` - Window API
- `:help nvim_create_buf` - Create buffers
- `:help nvim_open_win` - Open windows
- `:help nvim_buf_set_lines` - Write to buffers
- `:help nvim_get_keymap` - Extract keymaps
- `:help nvim_create_autocmd` - Create autocommands
- `:help lua-guide` - Lua in Neovim
