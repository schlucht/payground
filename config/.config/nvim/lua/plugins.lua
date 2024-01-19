-- Install your plugins here
return require('packer').startup(function(use)
	use ("wbthomason/packer.nvim") -- Have packer manage itself	
    use ('williamboman/mason.nvim')
    use ('williamboman/mason-lspconfig.nvim')
    use ('neovim/nvim-lspconfig')
    
    -- Plugins for code completion
    use ('hrsh7th/nvim-cmp')
    use ('hrsh7th/cmp-nvim-lsp')
    use ('hrsh7th/cmp-nvim-lua')
    use ('hrsh7th/cmp-nvim-lsp-signature-help')
    use ('hrsh7th/cmp-vsnip')
    use ('hrsh7th/cmp-path')
    use ('hrsh7th/cmp-buffer')
    use ('hrsh7th/cmp-vsnip')
    
end)
