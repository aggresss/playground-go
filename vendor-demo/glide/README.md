## glide

下载安装：

```
go get -u -v github.com/Masterminds/glide
```

特征文件：`glide.lock` `glide.yaml`

操作：

```
glide init
glide install
```

说明：

```
NAME:
   glide -Vendor Package Management for your Go projects.

   Each projectshould have a 'glide.yaml' file in the project directory. Files
   looksomething like this:

       package:github.com/Masterminds/glide
       imports:
       -package: github.com/Masterminds/cookoo
        version: 1.1.0
       -package: github.com/kylelemons/go-gypsy
        subpackages:
         - yaml

   For moredetails on the 'glide.yaml' files see the documentation at
  https://glide.sh/docs/glide.yaml


USAGE:
   glide [globaloptions] command [command options] [arguments...]

VERSION:
   0.13.0-dev

COMMANDS:
     create,init       Initialize a new project,creating a glide.yaml file
    config-wizard, cw  Wizard thatmakes optional suggestions to improve config in a glide.yaml file.
     get                Install one or more packagesinto `vendor/` and add dependency to glide.yaml.
     remove,rm         Remove a package from theglide.yaml file, and regenerate the lock file.
     import             Import files from other dependencymanagement systems.
     name               Print the name of this project.
     novendor,nv       List all non-vendor paths in adirectory.
    rebuild            Rebuild ('gobuild') the dependencies
     install,i         Install a project'sdependencies
     update,up         Update a project'sdependencies
     tree               (Deprecated) Tree prints thedependencies of this project as a tree.
     list               List prints all dependenciesthat the present code references.
     info               Info prints information aboutthis project
    cache-clear, cc    Clears theGlide cache.
     about              Learn about Glide
     mirror             Manage mirrors
     help,h            Shows a list of commands orhelp for one command

GLOBAL OPTIONS:
   --yaml value,-y value  Set a YAML configuration file.(default: "glide.yaml")
   --quiet,-q             Quiet (no info or debugmessages)
   --debug                 Print debug verboseinformational messages
   --homevalue            The location of Glidefiles (default: "/home/users/qiangmzsx/.glide") [$GLIDE_HOME]
   --tmpvalue             The temp directory touse. Defaults to systems temp [$GLIDE_TMP]
  --no-color              Turn offcolored output for log messages
   --help,-h              show help
   --version,-v           print the version
```

