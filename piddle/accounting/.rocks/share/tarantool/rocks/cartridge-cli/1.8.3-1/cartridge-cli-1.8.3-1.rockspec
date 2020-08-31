package = 'cartridge-cli'
version = '1.8.3-1'
source = {
    url = 'git+https://github.com/tarantool/cartridge-cli.git',
    tag = '1.8.3',
}

dependencies = {
    'lua >= 5.1',
}

external_dependencies = {
    TARANTOOL = {
        header = 'tarantool/module.h',
    },
}

build = {
    type = 'cmake',

    variables = {
        TARANTOOL_DIR = '$(TARANTOOL_DIR)',
        TARANTOOL_INSTALL_LUADIR = '$(LUADIR)',
        TARANTOOL_INSTALL_BINDIR = '$(BINDIR)',
        LUAROCKS = 'true',
    }
}
