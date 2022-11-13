leetgogo-cli
=========================

A leetcode cli tool implemented in go, inspired by project [leetcode-cli](https://github.com/skygragon/leetcode-cli). 

Tips:
- only support `leetcode.cn`.
- code templates initialization only supports go.

# Quick Start

```shell
# prepare cookie session and csrf token
leetgogo config leetcode.session <SESSION>
leetgogo config leetcode.csrftoken <TOKEN>

# init current workdir, to generate code templates after picking a problem
leetgogo init

# init a specified workdir
leetgogo init <path-to-workdir> && cd <path-to-workdir>

# pick a random problem
leetgogo pick

# pick problem by problem slug
leetgogo pick <problem-slug>
leetgogo pick https://leetcode.cn/problems/<problem-slug>

# favorite/unfavorite one or more problems to/from default `Favorite`
leetgogo favorite <problem-slug> [...<problem-slug>]
leetgogo unfavorite <problem-slug> [...<problem-slug>]

# favorite/unfavorite problem to/from specified favorite name
leetgogo favorite --favorite-name <favoriate-name> <problem-slug> [...problem-slug]
leetgogo unfavorite --favorite-name <favoriate-name> <problem-slug> [...problem-slug]

# execute problem solution with one or more test cases
leetgogo  exec [--solution-filename filename] [--test-cases-filename filename] problem_slug

# submit problem solution
leetgogo submit [--solution-filename filename] problem_slug
```

# Commands

```shell
leetgogo is a command tool to work with leetcode nicely.

Usage:
  leetgogo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      manage leetgogo config
  exec        execute problem solution on remote leetcode servers
  exec        submit solution code to leetcode server
  favorite    favorite problem by slug
  help        Help about any command
  init        initialize your workdir, code templates will be generated in your workdir
  pick        pick a leetcode problem by slug or url with slug, default to a random one
  unfavorite  unfavorite problem by slug

Flags:
  -h, --help      help for leetgogo
  -v, --version   version for leetgogo

Use "leetgogo [command] --help" for more information about a command.
```