git-zoo
===

🐶🐱🐭🐹🐰🦊🐻🐼🐨🐯🦁🐮🐷🐸🐵

This tool adds the emoji of animals in your git commit message. Let's make commit log a zoo.

## Feature
- Supports both `git commit` and `git commit -m` commands.
- Works with IDEs such as VSCode and Atom.

## Install

The binary has not been distributed yet. Install with `go get` command.

```bash
$ go get github.com/koirand/git-zoo
```

## Usage
Execute `git zoo install` on your git repository. This will create a git hook symbolic link in `.git/hooks/prepare-commit-msg`.
If you have already created git hooks, you can add the command as follows:

```
#!/bin/sh

...

git-zoo $1
```

## Example

```bash
$ echo "foo" >> README.md
$ git add --all
$ git commit -m "first commit"
[master (root-commit) f543f06] 🦁 first commit
 1 file changed, 1 insertion(+)
 create mode 100644 README.md

$ echo "bar" >> README.md
$ git commit -am "second commit"
[master e89092b] 🐸 second commit
 1 file changed, 1 insertion(+)

$ echo "baz" >> README.md
$ git commit -am "third commit"
[master ccbcd8a] 🐮 third commit
 1 file changed, 1 insertion(+)
```

