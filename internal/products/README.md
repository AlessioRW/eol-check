# Package - Products

Products are designed in a way that every check for all products work universally, this is to isolate the version checking logic which is expected to be expanded upon from the rest of the code.

## Design

This implementation relies on a `Product` and `Method` interface per product. 