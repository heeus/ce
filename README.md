[![codecov](https://codecov.io/gh/heeus/ce/branch/main/graph/badge.svg?token=FylcqUdTaR)](https://codecov.io/gh/heeus/ce)

# Basic usage

- go build -o ce.exe
- ./ce.exe -ihttp.Port 8888 server
- open http://localhost:8888/static/sys.monitor/site.hello/

# How it works

- cli/main parses args to params like `ihttp.CLIParams`, `ihttp.CLIParams` etc.
- Params are used to get `services` from `iservicesce.ProvideCEServices()`
- `services` then started using `iservicesctl.New().PrepareAndRun()`
