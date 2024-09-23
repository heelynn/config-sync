@echo off

echo config-sync starting...

:: Get the directory of the current script
set CURRENT_DIR=%~dp0
set CURRENT_DIR=%CURRENT_DIR:~0,-1%

:: Remove the last directory name to get the parent directory
for %%i in ("%CURRENT_DIR%") do set BASE_DIR=%%~dpi
set BASE_DIR=%BASE_DIR:~0,-1%

:: App file path
set APP=%BASE_DIR%\bin\config-sync.exe

:: App arguments
set APP_ARGS=-config "%BASE_DIR%\conf\application.yaml" -config-path "%BASE_DIR%\conf"

:: Command to execute the app
set COMMAND=%APP% %APP_ARGS%
:: Execute the command
%COMMAND%

pause "Press Enter to continue..."
:: If you want to see the command before executing, uncomment the line below:
:: echo %APP% %APP_ARGS%