@echo off

echo conf-sync starting...

:: Get the drive and path of the current batch file
set "CURRENT_DIR=%~dp0"

:: Remove the trailing backslash from the path if it exists
set "CURRENT_DIR=%CURRENT_DIR:~0,-1%"

:: Remove the last directory name to get the parent directory
for %%i in ("%CURRENT_DIR%") do set "BASE_DIR=%%~dpi"

set "BASE_DIR=%BASE_DIR:~0,-1%"


:: App file path
set APP=%BASE_DIR%\bin\conf-sync.exe

:: App arguments
set APP_ARGS=-config %BASE_DIR%\conf\application.yaml

:: Command to execute the app
set COMMAND=%APP% %APP_ARGS%

%COMMAND%