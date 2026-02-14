@echo off

net session >nul 2>&1
if %errorlevel% neq 0 (
    powershell -Command "Start-Process '%~f0' -ArgumentList '%*' -Verb RunAs"
    exit
)

node "D:\Projects\Web\focus\focus.js" %*
