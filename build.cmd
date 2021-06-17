@ECHO OFF
IF [%1]==[all] CALL :ALL
IF [%1]==[clean] CALL :CLEAN
IF [%1]==[zip] CALL :ZIP
IF [%1]==[windows] CALL :WINDOWS
IF [%1]==[linux] CALL :LINUX
EXIT /B 1

:ALL
    CALL :CLEAN
    CALL :WINDOWS
    CALL :LINUX
    CALL :ZIP
EXIT /B 0

:CLEAN
    CALL RMDIR "./dist" /q /s
EXIT /B 0

:WINDOWS
    SET GOOS=windows
    SET GOARCH=amd64
    CALL go build -trimpath -ldflags "-s -w" -o ./dist/windows/xmltransform.exe
    COPY "inputdata.xml" "./dist/windows/inputdata.xml"
    COPY "template.tmpl" "./dist/windows/template.tmpl"
    ECHO xmltransform.exe -i inputdata.xml -o outputFile.csv -t template.tmpl > ./dist/windows/testme.bat
    ECHO "WINDOWS ready"
EXIT /B 0

:LINUX
    SET GOOS=linux
    SET GOARCH=amd64
    CALL go build -trimpath -ldflags "-s -w" -o ./dist/linux/xmltransform
    COPY "inputdata.xml" "./dist/linux/inputdata.xml"
    COPY "template.tmpl" "./dist/linux/template.tmpl"
    ECHO "LINUX ready"
EXIT /B 0

:ZIP
    CALL 7z.exe a -tzip ./dist/dist.zip ./dist
    ECHO "ZIP ready"
EXIT /B 0