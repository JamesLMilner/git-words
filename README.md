# Git Words

Ever wanted to know a little bit more about a projects commit history?

This project will alow you determine what the most common words are in your git commit history.

## Building

```
go get
go build
```
Then you could do something like the following on Linux:

```
sudo cp git-words /bin/
```

Or anywhere in your PATH if you wanted it available in your terminal.

## Usage

```
cd some-git-repo/
git-words
```
## Flags

`--min` : Set the minimum number of occurences to log (default is **2**)
`--case` : Passing will make git-words case insensitive, all words become lower case (default is **false**