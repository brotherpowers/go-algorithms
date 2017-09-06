# Contributing Guidelines

Want to help out with the Go-Algorithms? Great! While we don't have strict templates on the format of each contribution, we do have a few guidelines that should be kept in mind:

**Readability**

The `README` file is the cake, and the sample code is the cherry on top. Readablity is really important for us. A good contribution has succinct explanations supported by diagrams. Code is best introduced in chunks, weaved into the explanations where relevant. 

**Language Guidelines**

A good contribution abides to the [Effective Go](https://golang.org/doc/effective_go.html). We review the pull requests with this in mind.

### New Contributions

Before writing about something new:

1. Check the main page for existing implementations
2. Check the [pull requests](https://github.com/brotherpowers/go-algorithms/pulls) for "claimed" topics. More info on that below. 

If what you have in mind is a new addition, please follow this process when submitting your contribution:

1. Create a pull request to "claim" an algorithm or data structure. This is to avoid having multiple people working on the same thing.
2. Write an explanation of how the algorithm works. Include examples for readers to follow along. Take a look at [quicksort](Sort/quickSort/) to get an idea.
3. Include your name in the explanation, something like *Written by Your Name* at the end of the document. 
4. Add unit tests.

For the unit tests:

- The unit tests will be run on [Travis-CI](https://travis-ci.org/brotherpowers/go-algorithms). You don't need to make any changes as all the test cases will be automatically run.

**All contributions are most welcome!**
