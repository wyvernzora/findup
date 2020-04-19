# findup
Finds the first file matching glob pattern, in the nearest ancestor directory. Similar to [sindresorhus/find-up](https://github.com/sindresorhus/find-up).

Alternative is [h2non/findup](https://github.com/h2non/findup), which always starts at current work directory.
Test suite adapted from [h2non/findup](https://github.com/h2non/findup).

## Installation
```
go get github.com/wyvernzora/findup
```

## Usage
```
import "github.com/wyvernzora/findup"

path, err := findup.Find("package.json")
fmt.Println(path)

path, err = findup.FindInDir("/some/directory/somewhere", "*.json")
fmt.Println(path)

```

## License
MIT
