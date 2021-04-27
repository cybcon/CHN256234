# CHN256234
CHN256234 - Programming request from a customer for a new project for [bridgingIT GmbH](https://www.bridging-it.de/).

## Abstract

This small example in golang will process integer intervals.
This intervals needs to be in a json format as described below.


read in integer intervals, given in a json format via command line or input file, will be parsed and 

## Json input format

The input needs to be in a json format, consists of an array representation in depth of 2.
The main array collect the different intervals:
```json
[
      <interval1>,
      <interval2>,
      <interval3>,
      ...
      <intervalN>
]
```

Every interval is also an array of 2 integer values. The lower- and the upper-end of the interval:

```
   [<lower>, <upper>]
```

### Example

```json
[
      [25,30],
      [2,19],
      [14,23],
      [4,8]
]
```

or without spaces and newline characters:

```json
[[25,30],[2,19],[14,23],[4,8]]
```

## Build the tool from source

```
c:\go\bin\go.exe build -a -installsuffix cgo -o interval_merger.exe .\src\interval_merger
```

## Execution

### usage

```
interval_merger.exe -h
Usage of interval_merger.exe:
  -file string
        The json file that contains the intervals
  -i string
        A json string that contains the intervals
```

### example

```
interval_merger.exe -i "[[25,30],[2,19],[14,23],[4,8]]"
[[2,23][25,30]]
```

## See also

- [Coding Task](./doc/Coding-Task.md)

