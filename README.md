leetgogo-cli
=========================

A leetcode cli tool implemented in go, inspired by project [leetcode-cli](https://github.com/skygragon/leetcode-cli). 

Tips:
- only support `leetcode.cn`.
- code templates initialization only supports go.

# Quick Start

```shell
# prepare cookie session and csrf token
leetgogo config session <LEETCODE_SESSION>
leetgogo config csrftoken <token>

# init current workdir, to generate code templates after picking a problem
leetgogo init

# init specified workdir
leetgogo init <path-to-workdir> && cd <path-to-workdir>

# pick a random problem
leetgogo pick

# pick problem by problem slug
leetgogo pick <problem-slug>
leetgogo pick https://leetcode.cn/problems/<problem-slug>

# favorite/unfavorite problem to/from default `Favorite`
leetgogo favorite <problem-slug>
leetgogo unfavorite <problem-slug>

# favorite/unfavorite problem to/from specified favorite name
leetgogo favorite <problem-slug> <favorite-name>
leetgogo unfavorite <problem-slug> <favorite-name>

# execute problem solution with one or more test cases
leetgogo exec <problem-slug> <solution-filename> <test-cases>

# submit problem solution
leetgogo submit <problem-slug> <solution-filename>
```