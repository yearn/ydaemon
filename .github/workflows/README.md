# testing

run workflows in a local container using https://github.com/nektos/act

eg to run the deploy workflow in your devenv, configure both a .env and .secrets file then run
```
act -j deploy --env-file .env
```
