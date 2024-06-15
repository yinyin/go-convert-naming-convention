# go-convert-naming-convention

[![GoDoc](https://pkg.go.dev/badge/github.com/yinyin/go-convert-naming-convention)][API]

[Go][] library and CLI program for converting the naming convention of given name string.

Support convertions for *lowerCamelCase*, *UpperCamelCase*, *kebeb-case* and *snake_case*.

## Usage

### Library

```
import (
	namingconv "github.com/yinyin/go-convert-naming-convention"
)

opts := namingconv.NewDefaultOptions()

varNameString := namingconv.ToLowerCamelCase("myVarName", opts)
varNameString := namingconv.ToUpperCamelCase("myVarName", opts)
varNameString := namingconv.ToSnakeCase("myVarName", opts)
varNameString := namingconv.ToKebebCase("myVarName", opts)
```

The default options will use common initialisms from [golint](https://golang.org/x/lint/golint).

### Command line program

The CLI program can be build with:

```sh
go build github.com/yinyin/go-convert-naming-convention/cmd/convert-naming-convention
```

The result is formatted as JSON. Here is the sample output of basic usage:

```text
$ ./convert-naming-convention myVariableName

{
  "Original": "myVariableName",
  "Splited": [
    "My",
    "Variable",
    "Name"
  ],
  "LowerCamelCase": "myVariableName",
  "UpperCamelCase": "MyVariableName",
  "SnakeCase": "my_variable_name",
  "KebebCase": "my-variable-name"
}
```

Common initialism and word segmentation override is supported:

*Without customization*:

```text
$ ./convert-naming-convention -- ftpReq nameGen2x typeG3dec typeG4Dec
{
  "Original": "ftpReq",
  "Splited": [
    "Ftp",
    "Req"
  ],
  "LowerCamelCase": "ftpReq",
  "UpperCamelCase": "FtpReq",
  "SnakeCase": "ftp_req",
  "KebebCase": "ftp-req"
}
{
  "Original": "nameGen2x",
  "Splited": [
    "Name",
    "Gen",
    "2x"
  ],
  "LowerCamelCase": "nameGen2x",
  "UpperCamelCase": "NameGen2x",
  "SnakeCase": "name_gen_2x",
  "KebebCase": "name-gen-2x"
}
{
  "Original": "typeG3dec",
  "Splited": [
    "Type",
    "G3dec"
  ],
  "LowerCamelCase": "typeG3dec",
  "UpperCamelCase": "TypeG3dec",
  "SnakeCase": "type_g3dec",
  "KebebCase": "type-g3dec"
}
{
  "Original": "typeG4Dec",
  "Splited": [
    "Type",
    "G4Dec"
  ],
  "LowerCamelCase": "typeG4Dec",
  "UpperCamelCase": "TypeG4Dec",
  "SnakeCase": "type_g4dec",
  "KebebCase": "type-g4dec"
}
```

*With customization*:

```text
$ ./convert-naming-convention --initial FTP --except Gen2x=Gen2-X -- \
    ftpReq nameGen2x typeG3dec typeG4Dec

{
  "Original": "ftpReq",
  "Splited": [
    "FTP",
    "Req"
  ],
  "LowerCamelCase": "ftpReq",
  "UpperCamelCase": "FTPReq",
  "SnakeCase": "ftp_req",
  "KebebCase": "ftp-req"
}
{
  "Original": "nameGen2x",
  "Splited": [
    "Name",
    "Gen2",
    "X"
  ],
  "LowerCamelCase": "nameGen2X",
  "UpperCamelCase": "NameGen2X",
  "SnakeCase": "name_gen2_x",
  "KebebCase": "name-gen2-x"
}
{
  "Original": "typeG3dec",
  "Splited": [
    "Type",
    "G3dec"
  ],
  "LowerCamelCase": "typeG3dec",
  "UpperCamelCase": "TypeG3dec",
  "SnakeCase": "type_g3dec",
  "KebebCase": "type-g3dec"
}
{
  "Original": "typeG4Dec",
  "Splited": [
    "Type",
    "G4Dec"
  ],
  "LowerCamelCase": "typeG4Dec",
  "UpperCamelCase": "TypeG4Dec",
  "SnakeCase": "type_g4dec",
  "KebebCase": "type-g4dec"
}
```

## License

Default common initialisms and tokenization code are heavily based on the `lintName` function of [golint](https://golang.org/x/lint/golint). These parts are licensed under [a BSD-style license](https://developers.google.com/open-source/licenses/bsd).

The rest part of code are licensed under [MIT License](LICENSE.md).

[API]: https://pkg.go.dev/github.com/yinyin/go-convert-naming-convention
[Go]: https://golang.org
