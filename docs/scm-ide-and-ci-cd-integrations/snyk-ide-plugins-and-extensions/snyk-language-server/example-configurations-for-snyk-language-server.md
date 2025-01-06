# Example configurations for Snyk Language Server

## Sublime Text를 위한 예제 구성

```
// 이 곳에 있는 설정은 "LSP/LSP.sublime-settings"의 설정을 덮어쓰게 됩니다.
{
  // 상태 표시줄에 Language Server 상태를 계속 표시합니다.
  "show_view_status": true,

  // Language Server 구성
  "clients": {
    "snyk": {
      // 이 구성을 활성화합니다
      "enabled": true,
      "command": [
        "/usr/local/bin/snyk-ls", // 다운로드한 이진 파일의 경로
        "-l", // 로그 레벨 정의
        "info", // info 레벨
        "-f", // 파일 로깅
        "/path/to/log/dir/snyk-ls-sublime.log" // 로그 파일
      ],
      // 이 Language Server가 연결되는 버퍼 유형을 선택하는 선택기입니다.
      "selector": "source", // 모든 파일 유형
      "initializationOptions": {
        "activateSnykCode":"true", // Snyk Code를 활성화합니다
        "token": "xxx" // 매번 인증할 필요 없이 Snyk 토큰을 설정합니다
      }
    },
  },
}
```

지원되는 파일을 열면 Sublime Text에서 Language Server가 시작되고 결과가 강조 표시됩니다.

![Sublime Text에서 표시된 Snyk Open Source 결과](https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252Fgit-blob-3a68fa1992e26cef384e0f1009dec2ea258fbe2e%252Fimage%2520%28166%29%2520%281%29%2520%281%29%2520%281%29%2520%281%29%2520%281%29%2520%281%29%2520%282%29.png%3Falt%3Dmedia\&width=768\&dpr=1\&quality=100\&sign=9a6b10b3\&sv=2)

![Sublime Text에서 표시된 Snyk Code 결과](<../../../.gitbook/assets/image (466) (1).png>)

## Neovim을 위한 예제 구성

{% hint style="info" %}
프로젝트 루트 디렉토리가 Git 정보에서 얻어지지 않는 경우, 이 스크립트는 루트 디렉토리를 `vim.loop.os_homedir()`로 설정하려고 시도합니다. 홈 디렉토리의 내용에 따라 메모리를 많이 사용할 수 있습니다. 이런 경우, 환경에 대체 루트 디렉토리를 설정하세요.
{% endhint %}

설정 방법은 다음과 같습니다:

```
mkdir -p ~/.config/nvim
touch init.lua
```

다음은 `init.lua` 스크립트의 예시입니다:

```
local on_windows = vim.loop.os_uname().version:match 'Windows'

local function join_paths(...)
    local path_sep = on_windows and '\\' or '/'
    local result = table.concat({ ... }, path_sep)
    return result
end

vim.cmd [[set runtimepath=$VIMRUNTIME]]

local temp_dir = vim.loop.os_getenv 'TEMP' or '/tmp'

vim.cmd('set packpath=' .. join_paths(temp_dir, 'nvim', 'site'))

local package_root = join_paths(temp_dir, 'nvim', 'site', 'pack')
local install_path = join_paths(package_root, 'packer', 'start', 'packer.nvim')
local compile_path = join_paths(install_path, 'plugin', 'packer_compiled.lua')

local function load_plugins()
    require('packer').startup {
        {
            'wbthomason/packer.nvim',
            'neovim/nvim-lspconfig',
        },
        config = {
            package_root = package_root,
            compile_path = compile_path,
        },
    }
end

_G.load_config = function()
    vim.lsp.set_log_level 'trace'
    if vim.fn.has 'nvim-0.5.1' == 1 then
        require('vim.lsp.log').set_format_func(vim.inspect)
    end
    local nvim_lsp = require 'lspconfig'
    local on_attach = function(_, bufnr)
        local function buf_set_option(...)
            vim.api.nvim_buf_set_option(bufnr, ...)
        end

        buf_set_option('omnifunc', 'v:lua.vim.lsp.omnifunc')

        -- Mappings.
        local opts = { buffer = bufnr, noremap = true, silent = true }
        vim.keymap.set('n', 'gD', vim.lsp.buf.declaration, opts)
        vim.keymap.set('n', 'gd', vim.lsp.buf.definition, opts)
        vim.keymap.set('n', 'K', vim.lsp.buf.hover, opts)
        vim.keymap.set('n', 'gi', vim.lsp.buf.implementation, opts)
        vim.keymap.set('n', '<C-k>', vim.lsp.buf.signature_help, opts)
        vim.keymap.set('n', '<space>wa', vim.lsp.buf.add_workspace_folder, opts)
        vim.keymap.set('n', '<space>wr', vim.lsp.buf.remove_workspace_folder, opts)
        vim.keymap.set('n', '<space>wl', function()
            print(vim.inspect(vim.lsp.buf.list_workspace_folders()))
        end, opts)
        vim.keymap.set('n', '<space>D', vim.lsp.buf.type_definition, opts)
        vim.keymap.set('n', '<space>rn', vim.lsp.buf.rename, opts)
        vim.keymap.set('n', 'gr', vim.lsp.buf.references, opts)
        vim.keymap.set('n', '<space>e', vim.diagnostic.open_float, opts)
        vim.keymap.set('n', '[d', vim.diagnostic.goto_prev, opts)
        vim.keymap.set('n', ']d', vim.diagnostic.goto_next, opts)
        vim.keymap.set('n', '<space>q', vim.diagnostic.setloclist, opts)
    end

    local lspconfig = require('lspconfig')
    local configs = require 'lspconfig.configs'

    if not configs.snyk then
        configs.snyk = {
            default_config = {
                cmd = {'/usr/local/bin/snyk-ls','-f','/path/to/log/snyk-ls-vim.log'},
                root_dir = function(name)
                    return lspconfig.util.find_git_ancestor(name) or vim.loop.os_homedir()
                end,
                init_options = {
                    activateSnykCode = "true",
                }
            };
        }
    end
    lspconfig.snyk.setup {
      on_attach = on_attach
    }

    print [[로그는 $HOME/.cache/nvim/lsp.log에서 찾을 수 있습니다. 이를 문제 템플릿에 설명된 대로 상세한 태그로 붙여넣으세요.]]
end

if vim.fn.isdirectory(install_path) == 0 then
    vim.fn.system { 'git', 'clone', 'https://github.com/wbthomason/packer.nvim', install_path }
    load_plugins()
    require('packer').sync()
    vim.cmd [[autocmd User PackerComplete ++once lua load_config()]]
else
    load_plugins()
    require('packer').sync()
    _G.load_config()
end
```

![Neovim에서 표시된 Snyk Code 결과](<../../../.gitbook/assets/image (219) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png>)
