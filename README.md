# DOCUMENTS (CNPJ AND CPF) VALIDATOR

Project to validate CNPJs/CPFs from a CSV file.

## Getting Started

Save a CSV file with CNPJs/CPFs to validate and run.

### CSV file format

You can insert CNPJs and CPFs at same CSV file.

```
00600158178664
00170610038389
00512008250801
00512084132765
00500112750504
54896550544
54898510159
```

### Constants

If necessary to change default values, you need to config constants and rebuild project. 

```
const (
	printValidDocuments   = false
	printInvalidDocuments = true
	filePathName          = "documents.csv"
)
```


