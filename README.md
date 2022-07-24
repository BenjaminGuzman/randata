# randata
Randomly generate data and export it to various formats (csv, json, xlsx...). 

Useful in testing to generate random test datasets

## Usage

### Options

Brief description of available options:

- `--count`: Number of random records (rows) to generate

- `--format`: Output format. If not given it'll be inferred from the output file extension (as long as it's valid).
  Valid formats are: csv, json

- `--out`: Output file

- `--mode`: Decide how to update the output file. Valid values are: overwrite, append

- `--fields`: Comma separated fields to project.
 
Fields are ignored in append mode. "ALL" is a special value, and it is equivalent to selecting all fields.

If you need a field that is not one of the pre-defined, you can provide its name followed by its format
(separated by a colon, e.g. `customField:### ??**`).
To specify the format you can use the following wildcards:
- `?` for a random letter
- `#` for a random digit
- `*` for a random ASCII char (between 32 (space) and 126 (tilde))

You can query the list of pre-defined columns in [generator.go](generator/generator.go)

You're also free to change the code as you please.

### Examples

Generate a CSV with 1000 records, each one containing a last name, password, zip (consisting of 5 digits), 
and a "sample" field. 

```shell
randata --out data.csv --count 1000 --fields "lastName,password,zip:#####,sample:letter? ascii* number#"
```

Generated output (data.csv):

```text
lastName,password,zip,sample
Reilly,{dkfndo{gbmhd|,89329,letterq asciiK number9
Ruecker,bd}rnv|vv,39090,letterk asciiU number5
McClure,fshy{kor,09436,letterc ascii< number1
Weimann,ejelxlauly,51831,lettera ascii0 number8
Murazik,b~ftnbiqk{xkjbxu,73276,lettero asciio number9
```

## License

![GPLv3](gplv3.png)