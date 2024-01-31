# ShakeSearch Challenge

Welcome to the Pulley Shakesearch Challenge! This repository contains a simple web app for searching text in the complete works of Shakespeare.

## Explanation of Changes
#### As of Jan 30 2024, Tue.
Changes made to the repo included:
1. Adding logic to ensure case-insensitive* queries.
    - *case types covered are lowercase, uppercase, and start case.
2. Adding support for multi-word queries.
3. Adding a default limit to the number of results returned on each initial query (20 as defined by the tests).
4. Adding functionality to the "Load More" button.
    - Additional results are loaded 20 at a time.
    - Offset pagination strategy since the following assumptions are being made:
        1. The offset would always be relatively small (e.g. 20 results at a time).
        2. The queried data will not change very much if at all.

View of all changes can be found [here](https://github.com/susannasouv/ShakeSearchChallenge/compare/d776f53f32df1db390f90dbe843cb0415229f903..master).

## Prerequisites

To run the tests, you need to have [Go](https://go.dev/doc/install) and [Docker](https://docs.docker.com/engine/install/) installed on your system.

## Your Task

Your task is to fix the underlying code to make the failing tests in the app pass. There are 3 frontend tests and 3 backend tests, with 2 of each currently failing. You should not modify the tests themselves, but rather improve the code to meet the test requirements. You can use the provided Dockerfile to run the tests or the app locally. The success criteria are to have all 6 tests passing.

## Instructions

<img width="404" alt="image" src="https://github.com/ProlificLabs/shakesearch/assets/98766735/9a5b96b5-0e44-42e1-8d6e-b7a9e08df9a1">

*** 

**Do not open a pull request or fork the repo**. Use these steps to create a hard copy.

1. Create a repository from this one using the "Use this template" button.
2. Fix the underlying code to make the tests pass
3. Include a short explanation of your changes in the readme or changelog file
4. Email us back with a link to your copy of the repo

## Running the App Locally


This command runs the app on your machine and will be available in browser at localhost:3001.

```bash
make run
```

## Running the Tests

This command runs backend and frontend tests.

Backend testing directly runs all Go tests.

Frontend testing run the app and mochajs tests inside docker, using internal port 3002.

```bash
make test
```

Good luck!
