param (
    [string]$portWeb = "8080",
    [string]$portBack = "8081",
    [switch]$f
 )

 # build front container
cd ..\..
docker build . -f APP/Dockerfile -t ogree-app
$assetsDir = "${PWD}\APP\assets\custom"
$file = "${assetsDir}\.env"
(Get-Content $file) -replace '8081', $portBack | Set-Content $file

# run container
$basename = "ogree-superadmin"
$containername = $basename
$index = 1
$result = @(docker ps --all --format "{{json .}}" --filter "name=$containername")
While ($result)
{
    if ($result -Match "failed") {
        Write-Host "Unable to check running containers! Try default name only"
        break
    }
    Write-Host "Container $containername already exists"
    if ($f.IsPresent) {
        Write-Host "Stopping it if running"
        docker stop $containername
    }
    $containername = "$basename-$index"
    $result = @(docker ps --all --format "{{json .}}" --filter "name=$containername")
    $index++
}

Write-Host "Launch $containername container"
docker run --restart always --name $containername -p ${portWeb}:80 -v ${assetsDir}:/usr/share/nginx/html/assets/assets/custom -d ogree-app:latest
if ($LASTEXITCODE -ne 0) {
    Write-Host "UNABLE TO LAUNCH WEBAPP CONTAINER, CHECK ERROR ABOVE" -ForegroundColor red
}

# compile and run back
cd BACK\app
docker run --rm -v ${PWD}:/workdir -w /workdir -e GOOS=windows golang go build -o ogree_app_backend.exe
.\ogree_app_backend.exe -port $portBack