# tumelo_task

## To run the task:

Run:

```
go run main.go
```

From the root directory

You will be asked to provide the path to the CSV file containing input recommendations data.

The CLI will provide information on the process

## Notes on the task:

- Overall it was really interesting, I definitely found it hard, it got me to learn a bunch more Go
- Please note that there are a few comments in this project which are for you guys to help understand my thinking, under normal circumstances almost none of these would be in the code, the only ones I would leave are docstrings for functions (which are builtin Go functionality)
- I believe commenting code in general is a poor practice and with proper naming and splitting out functions it should pretty much never be necessary

### Testing

- There are a lot more unit tests I would write for this before I would consider it production ready.
- I have included example tests for some areas like for csv_reader, recommendation.Validate() recommendation data cleaning, so you can see my approach
- Particularly the matching and indexing parts need a lot of testing attention

### Performance

- In many places where a large slice is being passed around I've used pointers, this is because otherwise we'd be copying that large amount of data when passed by value, in places many times.
