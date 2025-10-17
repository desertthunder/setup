local M = {}

--- Generate highlight groups based on colors and config
--- @param colors table Color palette
--- @param config table Configuration
--- @return table highlights All highlight groups
function M.setup(colors, config)
  local highlights = {}

  -- Helper function to create highlight group
  local function hl(group, opts) highlights[group] = opts end

  hl('Normal', { fg = colors.fg, bg = colors.bg })
  hl('NormalFloat', { fg = colors.fg, bg = colors.bg_float })
  hl('FloatBorder', { fg = colors.blue, bg = colors.bg_float })
  hl('NormalNC', { fg = colors.fg, bg = colors.bg })
  hl('ColorColumn', { bg = colors.bg_highlight })
  hl('Cursor', { fg = colors.bg, bg = colors.fg })
  hl('CursorLine', { bg = colors.bg_highlight })
  hl('CursorColumn', { bg = colors.bg_highlight })
  hl('LineNr', { fg = colors.fg_gutter })
  hl('CursorLineNr', { fg = colors.purple, bold = true })
  hl('SignColumn', { fg = colors.fg_gutter, bg = colors.bg })
  hl('Folded', { fg = colors.blue, bg = colors.bg_highlight })
  hl('FoldColumn', { fg = colors.fg_gutter, bg = colors.bg })

  hl('VertSplit', { fg = colors.grey })
  hl('WinSeparator', { fg = colors.grey })
  hl('StatusLine', { fg = colors.fg, bg = colors.bg_statusline })
  hl('StatusLineNC', { fg = colors.fg_dark, bg = colors.bg_statusline })
  hl('Pmenu', { fg = colors.fg, bg = colors.bg_popup })
  hl('PmenuSel', { bg = colors.bg_highlight })
  hl('PmenuSbar', { bg = colors.bg_highlight })
  hl('PmenuThumb', { bg = colors.fg_gutter })
  hl('TabLine', { fg = colors.fg_dark, bg = colors.bg_statusline })
  hl('TabLineFill', { bg = colors.bg_statusline })
  hl('TabLineSel', { fg = colors.fg, bg = colors.blue })

  hl('Visual', { bg = colors.bg_highlight })
  hl('VisualNOS', { bg = colors.bg_highlight })
  hl('Search', { fg = colors.bg, bg = colors.yellow })
  hl('IncSearch', { fg = colors.bg, bg = colors.orange })
  hl('Substitute', { fg = colors.bg, bg = colors.red })
  hl('MatchParen', { fg = colors.orange, bold = true })

  hl('ModeMsg', { fg = colors.fg, bold = true })
  hl('MoreMsg', { fg = colors.green })
  hl('Question', { fg = colors.blue })
  hl('ErrorMsg', { fg = colors.error })
  hl('WarningMsg', { fg = colors.warning })
  hl('Title', { fg = colors.purple, bold = true })
  hl('Directory', { fg = colors.blue })

  hl('DiffAdd', { bg = colors.git_add, fg = colors.bg })
  hl('DiffChange', { bg = colors.git_change, fg = colors.bg })
  hl('DiffDelete', { bg = colors.git_delete, fg = colors.bg })
  hl('DiffText', { bg = colors.diff_text, fg = colors.bg })

  hl('SpellBad', { sp = colors.error, undercurl = true })
  hl('SpellCap', { sp = colors.warning, undercurl = true })
  hl('SpellLocal', { sp = colors.info, undercurl = true })
  hl('SpellRare', { sp = colors.hint, undercurl = true })

  hl('Comment', vim.tbl_extend('force', { fg = colors.comment }, config.styles.comments))
  hl('Constant', { fg = colors.violet })
  hl('String', { fg = colors.orange })
  hl('Character', { fg = colors.green })
  hl('Number', { fg = colors.purple })
  hl('Boolean', { fg = colors.purple })
  hl('Float', { fg = colors.purple })

  hl('Identifier', vim.tbl_extend('force', { fg = colors.fg }, config.styles.variables))
  hl('Function', vim.tbl_extend('force', { fg = colors.green1 }, config.styles.functions))

  hl('Statement', vim.tbl_extend('force', { fg = colors.pink }, config.styles.keywords))
  hl('Conditional', vim.tbl_extend('force', { fg = colors.pink }, config.styles.keywords))
  hl('Repeat', vim.tbl_extend('force', { fg = colors.pink }, config.styles.keywords))
  hl('Label', { fg = colors.pink })
  hl('Operator', { fg = colors.pink })
  hl('Keyword', vim.tbl_extend('force', { fg = colors.pink }, config.styles.keywords))
  hl('Exception', { fg = colors.pink })

  hl('PreProc', { fg = colors.pink })
  hl('Include', { fg = colors.pink })
  hl('Define', { fg = colors.pink })
  hl('Macro', { fg = colors.pink })
  hl('PreCondit', { fg = colors.pink })

  hl('Type', { fg = colors.cyan })
  hl('StorageClass', { fg = colors.cyan })
  hl('Structure', { fg = colors.cyan })
  hl('Typedef', { fg = colors.cyan })

  hl('Special', { fg = colors.yellow })
  hl('SpecialChar', { fg = colors.yellow })
  hl('Tag', { fg = colors.pink })
  hl('Delimiter', { fg = colors.fg })
  hl('SpecialComment', { fg = colors.comment })
  hl('Debug', { fg = colors.orange })

  hl('Underlined', { underline = true })
  hl('Ignore', { fg = colors.grey })
  hl('Error', { fg = colors.error })
  hl('Todo', { fg = colors.purple, bold = true })

  if config.plugins.treesitter then
    hl('@variable', vim.tbl_extend('force', { fg = colors.fg }, config.styles.variables))
    hl('@variable.builtin', { fg = colors.purple })
    hl('@variable.parameter', { fg = colors.orange })
    hl('@variable.member', { fg = colors.cyan })

    hl('@constant', { fg = colors.violet })
    hl('@constant.builtin', { fg = colors.purple })
    hl('@constant.macro', { fg = colors.purple })

    hl('@function', vim.tbl_extend('force', { fg = colors.green1 }, config.styles.functions))
    hl('@function.builtin', { fg = colors.green1 })
    hl('@function.call', { fg = colors.green1 })
    hl('@function.macro', { fg = colors.pink })
    hl('@function.method', { fg = colors.green1 })
    hl('@function.method.call', { fg = colors.green1 })

    hl('@keyword', vim.tbl_extend('force', { fg = colors.pink }, config.styles.keywords))
    hl('@keyword.function', { fg = colors.pink })
    hl('@keyword.operator', { fg = colors.pink })
    hl('@keyword.return', { fg = colors.pink })
    hl('@keyword.conditional', { fg = colors.pink })
    hl('@keyword.repeat', { fg = colors.pink })
    hl('@keyword.import', { fg = colors.pink })

    hl('@string', { fg = colors.orange })
    hl('@string.escape', { fg = colors.green })
    hl('@string.regex', { fg = colors.yellow })
    hl('@string.special', { fg = colors.yellow })

    hl('@type', { fg = colors.cyan })
    hl('@type.builtin', { fg = colors.cyan })
    hl('@type.definition', { fg = colors.cyan })

    hl('@operator', { fg = colors.pink })
    hl('@punctuation.delimiter', { fg = colors.fg })
    hl('@punctuation.bracket', { fg = colors.fg })
    hl('@punctuation.special', { fg = colors.yellow })

    -- Markup (Markdown, etc.)
    hl('@markup.heading', { fg = colors.purple, bold = true })
    hl('@markup.strong', { bold = true })
    hl('@markup.italic', { italic = true })
    hl('@markup.underline', { underline = true })
    hl('@markup.link', { fg = colors.blue })
    hl('@markup.link.url', { fg = colors.cyan, underline = true })
    hl('@markup.raw', { fg = colors.orange })
    hl('@markup.list', { fg = colors.pink })

    -- Tags (HTML, JSX, etc.)
    hl('@tag', { fg = colors.pink })
    hl('@tag.attribute', { fg = colors.green1 })
    hl('@tag.delimiter', { fg = colors.fg })
  end

  if config.plugins.lsp then
    hl('DiagnosticError', { fg = colors.error })
    hl('DiagnosticWarn', { fg = colors.warning })
    hl('DiagnosticInfo', { fg = colors.info })
    hl('DiagnosticHint', { fg = colors.hint })

    hl('DiagnosticUnderlineError', { sp = colors.error, undercurl = true })
    hl('DiagnosticUnderlineWarn', { sp = colors.warning, undercurl = true })
    hl('DiagnosticUnderlineInfo', { sp = colors.info, undercurl = true })
    hl('DiagnosticUnderlineHint', { sp = colors.hint, undercurl = true })

    hl('DiagnosticVirtualTextError', { fg = colors.error })
    hl('DiagnosticVirtualTextWarn', { fg = colors.warning })
    hl('DiagnosticVirtualTextInfo', { fg = colors.info })
    hl('DiagnosticVirtualTextHint', { fg = colors.hint })

    hl('LspReferenceText', { bg = colors.bg_highlight })
    hl('LspReferenceRead', { bg = colors.bg_highlight })
    hl('LspReferenceWrite', { bg = colors.bg_highlight })
    hl('LspSignatureActiveParameter', { fg = colors.orange, bold = true })
    hl('LspCodeLens', { fg = colors.comment })
  end

  if config.plugins.telescope then
    hl('TelescopeBorder', { fg = colors.blue, bg = colors.bg_float })
    hl('TelescopeNormal', { fg = colors.fg, bg = colors.bg_float })
    hl('TelescopePromptBorder', { fg = colors.purple, bg = colors.bg_float })
    hl('TelescopePromptTitle', { fg = colors.purple, bold = true })
    hl('TelescopeResultsTitle', { fg = colors.blue })
    hl('TelescopePreviewTitle', { fg = colors.green })
    hl('TelescopeSelection', { fg = colors.fg, bg = colors.bg_highlight })
    hl('TelescopeSelectionCaret', { fg = colors.purple })
    hl('TelescopeMatching', { fg = colors.green1, bold = true })
  end

  if config.plugins.neotree then
    hl('NeoTreeNormal', { fg = colors.fg, bg = colors.bg_sidebar })
    hl('NeoTreeNormalNC', { fg = colors.fg, bg = colors.bg_sidebar })
    hl('NeoTreeRootName', { fg = colors.purple, bold = true })
    hl('NeoTreeDirectoryName', { fg = colors.blue })
    hl('NeoTreeDirectoryIcon', { fg = colors.blue })
    hl('NeoTreeFileName', { fg = colors.fg })
    hl('NeoTreeFileIcon', { fg = colors.fg })
    hl('NeoTreeGitAdded', { fg = colors.git_add })
    hl('NeoTreeGitModified', { fg = colors.git_change })
    hl('NeoTreeGitDeleted', { fg = colors.git_delete })
    hl('NeoTreeGitIgnored', { fg = colors.git_ignore })
    hl('NeoTreeIndentMarker', { fg = colors.grey })
  end

  if config.plugins.bufferline then
    hl('BufferLineFill', { bg = colors.bg_statusline })
    hl('BufferLineBackground', { fg = colors.fg_dark, bg = colors.bg_statusline })
    hl('BufferLineBufferSelected', { fg = colors.fg, bg = colors.bg, bold = true })
    hl('BufferLineBufferVisible', { fg = colors.fg, bg = colors.bg_statusline })
    hl('BufferLineModified', { fg = colors.orange, bg = colors.bg_statusline })
    hl('BufferLineModifiedSelected', { fg = colors.orange, bg = colors.bg })
    hl('BufferLineModifiedVisible', { fg = colors.orange, bg = colors.bg_statusline })
    hl('BufferLineSeparator', { fg = colors.grey, bg = colors.bg_statusline })
    hl('BufferLineSeparatorSelected', { fg = colors.bg, bg = colors.bg })
    hl('BufferLineSeparatorVisible', { fg = colors.grey, bg = colors.bg_statusline })
    hl('BufferLineIndicatorSelected', { fg = colors.purple, bg = colors.bg })
    hl('BufferLineCloseButton', { fg = colors.fg_dark, bg = colors.bg_statusline })
    hl('BufferLineCloseButtonSelected', { fg = colors.red, bg = colors.bg })
    hl('BufferLineCloseButtonVisible', { fg = colors.fg_dark, bg = colors.bg_statusline })
  end

  if config.plugins.gitsigns then
    hl('GitSignsAdd', { fg = colors.git_add })
    hl('GitSignsChange', { fg = colors.git_change })
    hl('GitSignsDelete', { fg = colors.git_delete })
  end

  if config.plugins.alpha then
    hl('AlphaHeader', { fg = colors.purple })
    hl('AlphaButtons', { fg = colors.green1 })
    hl('AlphaShortcut', { fg = colors.orange })
    hl('AlphaFooter', { fg = colors.cyan })
  end

  -- Allow user to override highlights
  if config.on_highlights then highlights = config.on_highlights(highlights, colors) or highlights end

  return highlights
end

--- Apply highlights to Neovim
--- @param highlights table Highlight groups
function M.apply(highlights)
  for group, opts in pairs(highlights) do
    vim.api.nvim_set_hl(0, group, opts)
  end
end

return M
