-- project
set_project("grpc-web-template")

-- set xmake minimum version
set_xmakever("2.6.1")

-- set warning all as error
set_warnings("all", "error")

add_subdirs("projects/core")
add_subdirs("projects/website")
