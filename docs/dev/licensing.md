The project is licensed under MIT license. This makes it incompatible with GPL licensed dependencies.

In order to honor all the licenses of the dependencies [go-licenses](https://github.com/google/go-licenses) is used. **In addition to that, all direct dependencies are manually inspected**.


Before pushing code run the following command to check for incompatiable licenses
```
go-licenses check --ignore "barman-exporter" ./... --disallowed_types=restricted,forbidden
```


You can also get a report for manual inspection by running
```
go-licenses report --ignore "barman-exporter" ./...
```

To get all the artifacts needed for distributing the binary the following command can be used:

```
go-licenses save --ignore "barman-exporter" ./... --save_path="3rd_parties"
```

The 3rd_parties output directory must be zipped and distributed with the binary.