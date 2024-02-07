set_project("website")

task("website:setup")
    set_category("action")
    on_run(function ()
        os.cd(os.scriptdir())
        os.exec("npm i")
    end)
    set_menu {
        usage = "xmake website:setup [options]",
        description = "",
        options = {}
    }

task("website:dev")
    set_category("action")
    on_run(function ()
        os.cd(os.scriptdir())
        os.exec("npm run dev")
    end)
    set_menu {
        usage = "xmake websitei:dev [options]",
        description = "run website in development",
        options = {}
    }

task("website:proto")
    set_category("action")
    on_run(function ()

        -- protobuf-ts approach
        os.cd(os.scriptdir())
        local core_proto_dir = "../core/proto"
        local out_dir = "src/apis/protobuf/core"
        local files, err = os.iorun("find " .. core_proto_dir .. " -path '*.proto'")
        local cmd = "npx protoc"
            .. " --experimental_allow_proto3_optional " -- flags
            .. " -I./../core/"
            .. " --ts_out " .. out_dir
            .." --ts_opt output_javascript"
            .. " --proto_path " .. core_proto_dir
            .. " " .. files

        --os.exec("npx protoc --ts_out src/apis/pb --proto_path ../core/proto/ user.proto")
        os.exec(cmd)

    end)
    set_menu {
        usage = "xmake website:proto [options]",
        description = "bring protobuf from services and compile them to web",
        options = {}
    }

target("website:dev")
    set_kind("phony")
    on_run(function (target)
        import("core.base.task")
        task.run("website:dev")
    end)