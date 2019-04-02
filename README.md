git-zoo
===

This tool adds animals emoji in git commit message. Let's make commit log a zoo.

## Feature
- Supports both `git commit` and `git commit -m` commands.
- Works with IDEs such as VSCode and Atom.

## Usage

The binary has not been distributed yet. Install with `go get` command.

```bash
go get github.com/koirand/git-zoo
```

Then, Make symbolic link of githook in your git repository.

```bash
ln -s $GOPATH/bin/git-zoo {your git repository}/.git/hooks/prepare-commit-msg
```

## Example

```bash
$ echo "foo" >> README.md

$ git commit -am "first commit"
🦁 first commit

[master (root-commit) f543f06] 🦁 first commit
 1 file changed, 1 insertion(+)
 create mode 100644 README.md

$ echo "bar" >> README.md

$ git commit -am "second commit"
🐸 second commit

[master e89092b] 🐸 second commit
 1 file changed, 1 insertion(+)

$ echo "baz" >> README.md

$ git commit -am "third commit"
🐮 third commit

[master ccbcd8a] 🐮 third commit
 1 file changed, 1 insertion(+)
```

