# CNPJs VALIDATOR

Project to validate CNPJs from a CSV file.

## Getting Started

Save a CSV file with CNPJs to validate an run.

### CSV file format

```
00600158178664
00170610038389
00512008250801
00512084132765
00500112750504
```

### Constants

If necessary to change default values, you need to config constants and build project. 

```
const (
	printCnpjsValidos   = false
	printCnpjsInvalidos = true
	filePathName        = "cnpjs.csv"
)
```


