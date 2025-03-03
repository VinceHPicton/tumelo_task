# tumelo_task

## To run the task:

Run:

```
go run main.go
```

From the root directory

You will be asked to provide the path to the CSV file containing input recommendations data.

- First use ExampleRecommendationsOriginal.csv (enter: ./ExampleRecommendationsOriginal.csv into the CLI to select it)
  This will demonstrate how the program handles invalid data
- Then use ExampleRecommendationsClean.csv (enter: ./ExampleRecommendationsClean.csv into the CLI to select it)
  This will demonstrate the program actually finding and matching data correctly.

The CLI will provide information on the process

## Notes on the task:

- Overall it was really interesting and enjoyable, I definitely found it hard too, it got me to learn a bunch more Go
- There are many TODOs thoughout the code which should give a lot of info on what I'd do with more time.
- Please note that there are a few comments in this project which are for you guys to help understand my thinking, under normal circumstances almost none of these would be in the code, the only ones I would leave are docstrings for functions (which are builtin Go functionality)
- I believe commenting code in general is a poor practice and with proper naming and splitting out functions it should pretty much never be necessary

### Testing

- There are a lot more unit tests I would write for this before I would consider it production ready.
- I have included example tests for some areas like for csv_reader, recommendation.Validate() recommendation data cleaning, so you can see my approach

### Performance

- In many places where a large slice is being passed around I've used pointers, this is because otherwise we'd be copying that large amount of data when passed by value, in places many times.

### Some other bits I'd like to do

- Use a .env file and pass API keys and other global data, in Go you need a dependency for this so I have left it out.
