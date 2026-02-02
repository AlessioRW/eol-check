# Am I EOL

## Overview

This is a Golang project that works on-top of the [endolflife.date](https://endoflife.date/) API. Checks are implemented per product available through the API, allowing for multiple methods of checking.

## Usesage Guide

The project takes a `.yaml` file, which the path to is passed as an argument when calling. The file is structured as such:

#### Example config
```yaml
config:
  - id: "go-main" # id of the check
    product: "golang" # what product is being checked
    path: "./go.mod" # path where to do check
    method: "file" # method of checking
  - id: "go-random"
    product: "golang"
    method: "command"
```

> Note: This is specifically for only running checks in a Golang project, though multiple products **can** be listed in the same config file, and a product may have different fields provided. Please check the specific README files for each product in the [products directory](./internal/products/)


**ID** - 
ID to store check results against, and can be any unique value. This is required for all products

**Product** - 
What product, language, serivce, etc, is being checked in the EOL API. See a list of available [here](https://endoflife.date/). This is required for all products

**Path** - 
Path to perform the check in.

**Method** - 
Method of checking for specific product. This is required for all products. Refer to product README for available methods and required configs.


#### Output

Once run, a file named `eol-output.csv` will be produced as such. Here is one generated of the example config above:

#### Example Output
```csv
id,product,version,is_eol,eol_date,maintained,lts
go-main,golang,1.24,false,N/A,true,false
go-random,golang,1.24,false,N/A,true,false
```

**id** - 
Id of check specified in config.

**product** - 
Product being checked.

**version** - 
Verison of product in your project.

**is_eol** - 
Boolean value if the product version being used is past it's end-of-life.

**eol_date** - 
Date that the product version being used will reach it's end-of-life. 

**maintained** - 
Boolean value if the product version being used is still being maintained.

**lts** - 
Boolean value if the product version being used is receiving long-term support# eol-check
