@echo off
echo =====================================
echo Removing Mark of the Web...
echo =====================================

powershell -ExecutionPolicy Bypass -Command ^
Unblock-File "%CD%\go-ftp-server-windows-amd64.exe"

echo.
echo =====================================
echo Adding Defender Exclusions...
echo =====================================

powershell -ExecutionPolicy Bypass -Command ^
Add-MpPreference -ExclusionProcess "%CD%\go-ftp-server-windows-amd64.exe"

powershell -ExecutionPolicy Bypass -Command ^
Add-MpPreference -ExclusionPath "%CD%"

echo.
echo Finished.
pause