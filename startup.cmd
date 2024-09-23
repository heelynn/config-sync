@echo off
:: 执行 wmic 命令并捕获输出
for /f "tokens=1,2" %%a in ('wmic process where "name='config-sync.exe'" get processid^,commandline /value') do (
    set "%%a=%%b"
)

:: 检查是否成功捕获到进程 ID
if defined PROCESSID (
    echo config-sync.exe found with CommandLine "%COMMANDLINE%" PID %PROCESSID%.
    set /p ANSWER="Do you want start a new instance of config-sync? (y/n): "

    if /i "%ANSWER%"=="y" (
        :: 在这里放置继续执行的命令
    ) else if /i "%ANSWER%"=="n" (
        exit /b 0
        :: 在这里放置不继续执行的命令
    ) else (
        echo Invalid input, please type 'y' or 'n'.
    )
)


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