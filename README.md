# CHN256234
CHN256234 - Programming request

## Build

```
c:\go\bin\go.exe build -a -installsuffix cgo -o interval_merger.exe .\src\interval_merger
```

## Execute

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
interval_merger.exe -i "[[25,30],[2,19],[14, 23],[4,8]]"
```