# hxu - Utility for Helix

This is an open source utility for more easily consuming resources from Q2's Helix product.

## Commands
`hxu file` - processes stdin stream as one of the fixed-width files emitted by Helix. Will emit a jsonlines-formatted stream to stdout. Auto detects file content based on file header line and parses appropriately.  Options exist to override this default behavior. Pass the `--help` flag to see all of them.

### Options
running `hxu help file` or `hxu file --help`:
```
Converts native fixed-width Helix files to different formats

Usage:
  hxu file [inputFile] [outputFile] [flags]

Flags:
      --format string   format [jsonlines,tsv,csv,verifyonly] (default "jsonlines")
      --help            display help content
      --noheader        suppress emitting the header line
      --preserve        preserve exact original content sans insignificant whitespace; do not apply conversions
      --strict          error if unexpected bytes appear in content lines
      --verifyonly      only verifies the file; does not emit content. synonym for --format=verifyonly
      --windows         use windows-style line endings (\r\n) instead of unix-style (\n)
```

The `--preserve` flag is to disable the auto-conversion of fields from string to their appropriate types and format.  For example, in the account balance file the `accountBalance` field as it appears in the file is `000000000001234`. By default, this is emitted by `hxu` as `12.34`, which is its semantic meaning.  To disable this conversion, specify the `--preserve` flag and it will emit from `hxu` as the original value, `000000000001234`. 

The `file` command does some implicit validation to ensure the file is valid:
- Ensures all bytes for all fields on every line exist (proper ignores extra fields)
- Compares line count reported by header with actual line count
If you want to perform this validation without converting and emitting file content, pass the `--format=verifyonly` flag

`hxu help` - displays helpful information for using non-default behavior of hxu. e.g. arguments and flags for each individual command.

### Caveats
jsonlines-formatted files are, by definition, a single json object per line.  json does not enforce any kind of ordering of properties on an object when serializing, so json properties may be in different order from one line to the next.  Fields with no significant characters (e.g. empty string or all whitespace) will be omitted to conserve file size.

The intent of this conversion utility is not necessarily to have smaller sized files; it is to create files that are more easily consumed. However, jsonlines-formatted files are on average about 1/2 the original file size.  CSV and TSV files are roughly about 1/6 the original file size.

jsonlines-formatted files retain the original header information as its first line, and as such, has a different schema than all remaining lines in the file.  CSV- and TSV-formatted files do not retain original header information; it is lost in the conversion. The first line of CSV- and TSV-formatted files are the individual field names.  The header line can be suppressed for any output format using the `--noheader` flag.