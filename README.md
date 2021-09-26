# archive

A utility written in Go which adds archiving functionality to your directories. Archiving in this context refers to archiving as found in email clients, etc.

## Motivation

Email clients and productivity software (e.g. to-do list or note taking applications) usually offer an 'archive' feature, which users can use to move items which are not relevant anymore to a separate area to reduce cluttering. This way, the user can focus on items which are relevant right now, while still keeping old items for reference. This utility aims to add this functionality to Unix-like systems.

## How does it work?

It is essentially a specialised interface to `tar`. Instead of the user having to worry about remembering and typing long tar commands, handling errors, keeping track of tar files, etc. this script abstracts away all that to make archiving as simple as archiving an email with the click of a button: `archive -a file.txt`.

## Usage

Each directory has its own archive file (called `.archive.tar`) that's created when first file is archived. That's where all archived files are stored. Add file to CWD's archive (without removing the file):

```sh
archive -a file.txt
```

To move file to archive (add to archive, remove it afterwards):

```sh
archive -a file.txt -d
```

To archive a gzipped version of the file (gzip the file, add gzipped version to archive, remove the gzipped file but keep the original unchanged):

```sh
archive -a file.txt -z
```

Pass the `-d` flag to remove the original, un-gzipped file after adding the gzipped version to archive:

```sh
# Gzip the file to file.txt.gz, add file.txt.gz to archive,
# then remove 'file.txt.gz' and remove 'file.txt' as well.
archive -a file.txt -zd
```

To unarchive a file (but still keep it in the archive):

```sh
archive -u file.txt
```

To unarchive a file and remove it from the archive:

```sh
archive -u file.txt -d
```

To list archive contents:

```sh
archive -l
```

## Options

```sh
-a, --add FILE          # Add file to archive of current directory
-u, --unarchive FILE    # Unarchive file from archive of current directory
-d, --delete            # Pass flag to -a, -u or -z to delete file in dir/archive after operation
-z, --gzip              # Used with -a to gzip the file/dir before archiving it. Original file is 
                        # not affected (i.e. not gzipped) but will be deleted if -d is passed.
-l, --list              # List the files in current directory archive
-t, --top-level         # List only top-level files and directories in current directory archive
-v, --version           # Print version and exit
-h, --help              # Print help and exit
```

## Development

- The utility is written in Go.
- The tests are written using Go's built-in [`testing` package](https://pkg.go.dev/testing). To run the tests:

```sh
go test ./test
```

- [gofmt](https://pkg.go.dev/cmd/gofmt) is used for automatically formatting the code to enforce a level of consistency, as well as simplify it. To run the formatter, run the `format.sh` script in the project root directory:

```sh
sh format.sh
```
