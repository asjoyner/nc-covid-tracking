# nc-covid-tracking
Some scripts and data to help gather and process data from the NC DHHS about COVID-19.

For some visualizations of this data, see this [Google Sheet](https://docs.google.com/spreadsheets/d/1vICoIsna_KlC_vonRJoc5UoN85pkQcslAU_E0bj54iY/preview).

## bash scripts

There are a few `bash` scripts in this repository which simplify fetching data from the DHHS website.

1 `fetch` returns the JSON representation of the DHHS data.  It just wraps the long ArcGIS URL in curl.
1 `tocsv` makes a handy CSV representation of that by county, for tinkering with in Google Sheets, etc.
1 `daily-fetch` is the script I run to populate the `data/` directory.

## data/ directory

I also capture the daily output of those scripts in the data/ directory, to establish a historical record of what DHHS published on what day.  `daily-fetch` is a slightly optimized and more consistent way to do that, but it basically does this:

```
./fetch > data/$(date --rfc-3339=date).json
./fetch | ./tocsv > data/$(date --rfc-3339=date).csv
```

# Go horizontal CSV

To make a stacked graph of this data by County over time in Google Sheets I needed to group the data by county "horizontally".  Each county needs to be a column, with the number of cases in colums below it.  That's tedious to do in Bash, so main.go contains some code to pull the JSON and output it in a horizontal representation.
