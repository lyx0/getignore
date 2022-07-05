# getignore
Getignore gets a `.gitignore` file from [https://github.com/github/gitignore](https://github.com/github/gitignore) and puts it in the current folder. 

If there is a `.gitignore` currently in the folder this tool will **overwrite** it so it is really only meant for the initial setup.

## Usage
Using `go install`:
```sh
$ go install .
```

Or build it and manually move the `getignore` executable somewhere in your path
```sh
$ go build .
```

Then run it with a `lang` command line flag. For example:
```sh
$ getignore --lang="rust"
$ getignore --lang="go"
$ getignore --lang="c++"
```


