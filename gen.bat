@echo off

@setlocal

set GENERATOR_PATH=generator/main.go

"go run" "%GENERATOR_PATH%" %*

@endlocal