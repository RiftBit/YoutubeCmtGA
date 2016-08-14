@echo Starting YT Comment GiveAway
@setx -m GODEBUG "cgocheck=0"
@set GODEBUG=cgocheck=0
call .\YTCommentGiveAway_windows_x64.exe
@echo Finished
pause