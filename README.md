# Hijri Holiday CLI

`hijri-holiday` is a simple command line to crawl `time.ir` and extract holidays for current year.

## Usage

```shell
hijri-holiday -o <json_path>
```
This will create a json file specified by `json_path` argument.

## Sample Output

```
{
  "scrap_time": "2024-07-25T15:08:13.684666968+03:30",
  "year": "1403",
  "holidays": []
}
```