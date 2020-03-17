# nc-covid-tracking
Some scripts and data to help gather and process data from the NC DHHS about COVID-19.

## Snapshot today's numbers

```
./fetch > data/$(date --rfc-3339=date).json
./fetch | ./tocsv > data/$(date --rfc-3339=date).csv
```
