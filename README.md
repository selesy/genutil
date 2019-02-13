# go-genutil

Code generation utilities to make it easier to write Go generators.

## Rationale

A typical Go generator takes a list of packages to process, and possibly
command line flags and/or parameters, and finds references to code that
should be generated along with generator-specific "directives".

Go doesn't have annotations for program meta-data and seems to prefer
generation over run-time reflection.  This library provides the following
operations that are common to Go generators:

- Selection of one or more packages to be processed
- Filtering of elements to be processed
- Creation of a specification from the discovered directives
