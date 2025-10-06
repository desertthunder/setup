local M = {}

local state = {
  win = nil,
  buf = nil,
}

function M.setup(opts)
  local config = require 'cheatsheet.config'
  config.setup(opts)
end

function M.toggle()
  if state.win and vim.api.nvim_win_is_valid(state.win) then
    M.close()
  else
    M.open()
  end
end

function M.open()
  local ui = require 'cheatsheet.ui'
  local keymaps = require 'cheatsheet.keymaps'
  local grouped_keymaps = keymaps.extract()
  state.buf, state.win = ui.create_window(grouped_keymaps)
end

function M.close()
  if state.win and vim.api.nvim_win_is_valid(state.win) then
    vim.api.nvim_win_close(state.win, true)
    state.win = nil
  end

  if state.buf and vim.api.nvim_buf_is_valid(state.buf) then
    vim.api.nvim_buf_delete(state.buf, { force = true })
    state.buf = nil
  end
end

return M
