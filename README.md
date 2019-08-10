# auspost-postcode

## Install

```
go get github.com/rhomber/auspost-postcode
```

## Usage

```
client := auspost.NewDefaultClient("193b1ab5-d4cf-4143-af9d-bd4ad4bd4e52")
res, err := client.PostcodeSearch("Melton", "", false)
```