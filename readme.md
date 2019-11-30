# How much time has passed?

You can see how much time has passed with start date in year, month, day, hour, minute, second format.

All you have to do is update the date.json file and then run the project.


## Installation

### Clone

- Clone this repo to your local machine using `https://github.com/erenakpinar/go-how-much-time-passed`

### Setup
- Create a new dates.json file.

### Example

- We left a sample item for you.
```
[
  {
    "Subject": "My Birthday",
    "StartDate": "2019-11-30"
  },
  {
    "Subject": "My Birthday",
    "StartDate": "2019-11-30",
    "EndDate": "2019-12-01"
  }
]
```
#### Output
```
Example 1: 0 days; 0 years, 0 months, 0 days, 16 hours, 21 mins and 1 seconds.
Example 2: 1 days; 0 years, 0 months, 1 days, 0 hours, 0 mins and 0 seconds.
```

### Running
```
$ go run main.go
```

