# Run tests
## Important
Be careful with the following tests, as they don't act on the database level which could lead to changes of your Domino installation. Therefore, it would be a wise decision to only run those tests in a virtual environment with a separate HCL Domino installation.
- notesadministrationprocess
- notesregistration
- notesreplica

## Preparation
To be able to run the tests, a local database file called *"GoInterface.nsf"* ist required. The database file must contain the following objects.
- A form called "TestForm" which contains a text-field with the name "testField".
- An agent called "TestAgent".

Additionally, the *-Default-* user needs to have *Manager* access rights with all permissions enabled.

## Execution
When the tests run, a copy of the local database is created to make sure it doesn't get changed. To run a test, switch into one of the directories within the *tests*-directory and run *go test -v*. **IMPORTANT:** If you HCL Domino version is 32 bit, make sure the enviroment variable *GOARCH=386* is set.
