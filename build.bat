@echo off
REM ============================================================
REM  Focus — build script for Windows
REM  Run this from the project root in a normal Command Prompt.
REM  Go 1.21+ must be installed: https://go.dev/dl/
REM ============================================================

echo [1/3] Downloading dependencies...
go mod tidy
if errorlevel 1 (
    echo ERROR: go mod tidy failed. Is Go installed?
    pause
    exit /b 1
)

echo [2/3] Building focus.exe...
if not exist dist mkdir dist

REM -H windowsgui hides the console window
REM -s -w strips debug info for a smaller binary
go build -ldflags="-H windowsgui -s -w" -o dist\focus.exe .\cmd\focus
if errorlevel 1 (
    echo ERROR: Build failed. See output above.
    pause
    exit /b 1
)

echo [3/3] Done!
echo.
echo  Output: dist\focus.exe
echo.
echo  To run: right-click dist\focus.exe ^> Run as administrator
echo  Or add dist\ to your PATH and type: focus
echo.
pause
