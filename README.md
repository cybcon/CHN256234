# CHN256234
CHN256234 - Programming request from a customer for a new project for [bridgingIT GmbH](https://www.bridging-it.de/).

see also: [Coding Task](./doc/Coding-Task.md)

## Abstract

This small example in golang will process integer intervals.
This intervals needs to be in a json format as described below.

## Json input format

The input needs to be in a json format, consists of an array representation in depth of 2.
The main array collect the different intervals:

```
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

## Build and runtime alternatives

### Docker

You will have a [`Dockerfile`](./Dockerfile) inside this repo to build a docker image out of the golang source.

Example:
```
git clone https://github.com/cybcon/CHN256234.git
docker build -t interval_merger:latest ./CHN256234/
```

After building the image, you can run the image
example:

```
# docker run --rm interval_merger -i "[[25,30],[2,19],[14,23],[4,8]]"
[[2,23][25,30]]
```


A ready to use docker image can be pulled from Docker Hub: https://hub.docker.com/r/oitc/chn256234

```
docker run --rm oitc/chn256234:latest -i "[[25,30],[2,19],[14,23],[4,8]]"
```

### Azure DevOps CI/CD pipeline script

Inside this repo, there is also an Azure DevOps CI/CD build pipeline definition: [`azure-pipelines.yml`](azure-pipelines.yml).
This can also be used to build the docker container and automatically push to a docker repository.


## FAQ

- Raw coding time of this solution: ~4h
- Resilience:
  - Map input parameters to a struct, invalid values will be ignored
  - Large input parameters over command line will not be processed due to bash input limitations
  - Large input via input file: **TODO**
- Runtime metrics:
  - Azure VM parameters:
    - Region: West Europe
    - Image: Ubuntu Server 18.04 LTS - Gen 1
    - Size: Standard_D4as_v4 - 4 vcpus, 16 GiB memory
    - OS disk type: Premium SSD
  - runtime measurement with the Linux command: `time`


| intervals |  runtime | memory consumption |
|----------:|---------:|-------------------:|
|         5 | 0m0.002s |                    |
|        50 | 0m0.002s |                    |
|       100 | 0m0.002s |                    |
|       500 | 0m0.003s |                    |
|     1.000 | 0m0.003s |                    |
|     5.000 | 0m0.008s |                    |
|    10.000 | 0m0.014s |                    |
|    50.000 | 0m0.065s |                    |
|   100.000 | 0m0.131s |                    |
|   500.000 | 0m0.651s |                    |
| 1.000.000 | 0m1.297s |                    |
