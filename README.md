# column-bling

Use column-bling to colorize column output.

![Screenshot 2021-05-25 at 12 51 19](https://user-images.githubusercontent.com/5672001/119485935-f713b380-bd57-11eb-861d-c6557e3a60a2.png)


## Install

```sh
git clone https://github.com/reegnz/column-bling
cd column-bling
go build
```

## Usage

### With [column](https://man7.org/linux/man-pages/man1/column.1.html) and csv files

```sh
column -s, -t data.csv | ./column-bling
```

### With [kubectl](https://kubernetes.io/docs/reference/kubectl/kubectl/) column output

```sh
kubectl get pods -A | ./column-bling
```
