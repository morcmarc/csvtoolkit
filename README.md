# csvtoolkit

Toolkit for querying and converting csv.

## Usage

### `convert`

Converts the given csv file into json.

```
$ csvtoolkit convert -o outputfile.json inputfile.csv
```

### `query`

Runs the given query against a csv file. The syntax is based on [jq](http://stedolan.github.io/jq/)'s.

Example:

```
$ csvtoolkit query -q '.[0] | keys()' inputfile.csv
```

This will pick the first item in the file and feed it into the `keys` function
which will then return the name of the columns.

You can also use the `.` filter as a function, the query above can be written as:

```
$ csvtoolkit query -q '.(0) | keys' inputfile.csv
```

## Roadmap

- Be able to turn off type inference during convert
- Refactor type inference and convert packages
- Cover `convert` with tests