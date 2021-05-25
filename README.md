# column-bling

Use column-bling to colorize column output.

## Install

```sh
git clone https://github.com/reegnz/column-bling
cd column-bling
go build
```

## Usage

### With `column` and csv files

```sh
column -s, -t data.csv | ./column-bling
```

### With `kubectl` column output

```sh
kubectl get pods -A | ./column-bling
```
