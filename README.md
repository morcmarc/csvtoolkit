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

```
$ csvtoolkit query -q 'keys()[0]' inputfile.csv
```

## Roadmap

- Be able to turn off type inference during convert
- Refactor type inference and convert packages
- Cover `convert` with tests