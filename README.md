# eol-check

## Overview

This is a config-driven Golang project that works on-top of the [endolflife.date](https://endoflife.date/) API. Checks are implemented per product available through the API, and allows for multiple implementations of checking called 'methods'.

## Usesage Guide

The project takes a `.yaml` file, which the path to is passed as an argument when calling. The file is structured as such:

#### Example config
```yaml
config:
  - id: "golang file check" # id of the check
    product: "golang" # what product is being checked
    method: "file" # method of checking
    args: # array of arguments to pass to check, specific for each method
      - "./go.mod"
  - id: "aws glue"
    product: "aws_glue"
    method: "sdk"
    args: 
      - "glue-job-name" 
```

> Note: This is specifically for only running checks in a Golang project, though multiple products **can** be listed in the same config file, and a product may have different fields provided. Please check the specific README files for each product in the [product's directory](./internal/products/)


**ID** - 
ID to store check results against, and can be any unique value. This is required for all products

**Product** - 
What product, language, serivce, etc, is being checked in the EOL API. See a list of available [here](https://endoflife.date/). This is required for all products

**Method** - 
Method of checking for specific product. This is required for all products. Refer to product README for available methods and required configs.

**Args** - 
Array of arguments to pass into check. These are specific to each method and are documented in the product's directory [PRODUCT].md file

#### Output

Once run, a file named `eol-output.csv` will be produced as such. Here is one generated of the example config above:

#### Example Output
```csv
id,product,version,is_eol,eol_date,maintained,lts
go-file,golang,1.11,true,2019-09-03,false,false
glue-sdk,aws_glue,5.0,false,N/A,true,false
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
