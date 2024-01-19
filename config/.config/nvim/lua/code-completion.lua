-- complete opt used to manage code suggestion format

vim.opt.completeopt = {'menuone', 'noselect', 'noinsert', 'preview' }
vim.opt.shortmess = vim.opt.shortmess + { c = true }

local cmp = require ('cmp')

cmp.setup({
        --Configurations
        --sources are installed that can be used for code suggestions
        sources = {
                { name = 'path' },
                { name = 'nvim_lsp', keyword_length = 3 },
                { name = 'nvim_lsp_signature_help', keyword_length=2 },
                { name = 'nvim_lua', keyword_length = 2},
                { name = 'buffer', keyword_length = 2 },
                { name = 'vsnip', keyword_length = 2 },
        },
        -- Window for styling
        window = {
            completion = cmp.config.window.bordered(),
            documentation = cmp.config.window.bordered(),
        },

         -- Mapping are keyboard shortcuts to execute commands
         mapping = {
            -- SHIFT+TAB go to previous item
            ['<S-Tab>'] = cmp.mapping.select_prev_item(),
            -- TAB go to next item
            ['<TAB>'] = cmp.mapping.select_next_item(),
            -- CTRL-SPACE to bring up the code completion
            ['<C-Space>'] = cmp.mapping.complete(),
            -- CTRL+e close the tree
            ['<C-e>'] = cmp.mapping.close(),
            -- Pressing ENTER or RETURN will confirm an apply
            ['<CR>'] = cmp.mapping.confirm({
                    behavior = cmp.ConfirmBehavior.Insert,
                    select = true,
            }),
    }
})

