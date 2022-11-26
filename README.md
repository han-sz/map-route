# map-route
Generates a geo-plot map of NSW train, bus and ferry stops

![](generated-plot.png)


## Build

```
make build
chmod +x build/map-route
```

## Run

```
./build/map-route
```

or with inputs

```
map-route -w 1920 -h 1080 -i data/routes.csv -o generated-plot.png
```