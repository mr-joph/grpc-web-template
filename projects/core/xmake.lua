set_project("core")

task("core:dev")
    set_category("action")
    on_run(function ()
        os.cd(os.scriptdir())
        os.exec("go run pkg/main.go")
    end)
    set_menu {
        usage = "xmake core:dev [options]",
        description = "run core api in development",
        options = {}
    }

task("core:install-dep")
    set_category("action")
    on_run(function ()
        import("core.base.option")
        
        arg = table.concat(option.get("contents") or {}, " ")
        os.cd(os.scriptdir())
        os.exec("go get " .. arg)
    end)
    set_menu {
        usage = "xmake api:install [options]",
        description = "install deps",
        options = {
            {nil, "contents", "vs", nil, "The info contents." }
        }
    }

task("core:proto")
    set_category("action")
    on_run(function ()
        os.cd(os.scriptdir())

        local proto_dir = "proto/"
        local out_dir = "pkg/"
        -- Get all proto files
        local files, err = os.iorun("find " .. proto_dir .. " -path '*.proto'")
        print(files)
        -- build compile command
        local cmd = 'protoc'
            .. " --experimental_allow_proto3_optional " -- flags
            .. ' --go_out=' .. out_dir
            .. ' --go-grpc_out=' .. out_dir
            .. ' ' .. files

        os.exec(cmd)
    end)
        set_menu {
        usage = "xmake core:proto [options]",
        description = "compile protobuffers",
        options = {
        }
    }

task("core:cli")
    set_category("action")
    on_run(function ()
        os.cd(os.scriptdir())
        os.exec("evans -p 7070 -r repl")
    end)
        set_menu {
        usage = "xmake core:cli [options]",
        description = "runs evans, a client in command line to make operations on gRPC server",
        options = {
            {}
        }
    }

task("core:migrate")
    set_category("action")
    on_run(function ()
        import("core.base.option")
        os.cd(os.scriptdir())
        
        action = table.concat(option.get("action") or {}, " ")
        os.exec("go run migrate/main.go " .. action)
    end)
        set_menu {
        usage = "xmake core:migrate [options]",
        description = "run migrations",
        options = {
            {nil, "action", "vs", "migration task, like 'up' 'down' or 'new'"}
        }
    }
