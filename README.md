# Gymshark Middleweight Technical Test

## Introduction

I was briefed with creating an API in Go that can handle the processing of X number of units into Y amount of packs following some set rules.

1. Only whole packs can be sent.
2. Send out no more items than necessary.
3. Send as few packs as possible.

Optional Tasks

1. Flexible application as to enable change, addition and removal of packs.
2. Create a UI to interact with the API.

## Solution

**I'd like to create this application while completing all of the optional tasks where possible. With a relatively simple backend HTTP server written in Go that handles required HTTP requests, I'm initially thinking of serving a single HTML file with internal vanilla Javascript and internal CSS for the client. I could create a frameworked front-end, but why the bloat? KISS (Keep It Simple, Stupid)!**

With a currently limited but increasing understanding of Go, I'm hoping to create an application that utilises best implementation (Correctness over simplicity), using Go's pointers to store pack sizes in memory without increasing memory heap size where possible when changing pack sizes. A simple HTML front-end should allow the manipulation of _number of units_ and _pack sizes_ to start, where I'll add the complexity of removing and adding new pack sizes later.

### Client

- Input box to register the amount of units needing to be sent. On change of input, resets the 2nd column of packages required.
- Send button to tell the server to perform pack calculations. ✅
- A 2x5 table where column 1 is the pack size and column 2 is the required number of packs for the relevant pack size. ✅

| Pack Size | Number of Packs |
| :-------: | :-------------: |
|    250    |        1        |
|    500    |        4        |
|   1000    |        2        |
|   2000    |        1        |
|   5000    |        0        |

Optional Tasks

- Column 1 pack sizes able to be changed.
- Button to add a new row (Pack size) to the table.
- Button to remove a row (Pack size) from the table.

### API

| Method | URL    | Description                                                                                                     |
| ------ | ------ | --------------------------------------------------------------------------------------------------------------- |
| GET    | /      | Retrieves HTML static homepage                                                                                  |
| POST   | /units | Send units JSON `{ value: number }` and return num of packages `{Packages: [n, n, n, n, n]}` in ascending order |

- HTTP server that serves the HTML file from local folder. ✅
- HandleFunc GET '/' route. Serve HTML file ✅
- HandleFunc POST '/units' route. Send number of units required, handle pack calculation and return pack values. ✅

Optional Tasks

- HandleFunc PUT '/pack' route. Send the relevant pack size to change the size of a pack when editing a column 1 cell.
- HandleFunc DELETE '/pack' route. Delete the relevant pack.
- HandleFunc POST '/pack' route. Add a new pack size.

## Notes, Thoughts and Process

- Some suggest using Gorilla Mux to route an API, but in the interest of learning and simplicity, I'd like to forego Gorilla Mux and instead use 'vanilla Go' where possible
- Starting without optional task due to limited Go knowledge, would like to attempt to use slices to give min/max package options when adding/removing packages to conserve memory.
- Deployment issues through Elastic Beanstalk and Heroku. Decided to go with Fly.io and use Docker to solve the issues. Recreated the routes using Mux Gorilla and served the HTML file using Go 1.16's 'embed' which seemed to solve the issues.
