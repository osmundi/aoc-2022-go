
### Init folders
for i in {1..24}; do mkdir day_${i}; done

### Generate base files for swifter development
for i in {1..24}; do cp template.go.copy day_${i}/main.go; done

### Fetch input data
for i in {1..24}; do curl -v  "https://adventofcode.com/2022/day/${i}/input" -H 'Cookie: session=<sessionid>' > day_${i}/data.csv ; done

