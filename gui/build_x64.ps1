
# 获取系统时间
$origin_date=Get-Date
$build_time=$origin_date.ToString('yyyy-MM-dd hh:mm:ss')
$commit_id=git rev-parse --short HEAD

$last_tag_commit_id=git rev-list --tags --max-count=1
$last_tag=git describe --tags $last_tag_commit_id
Write-Host "BuildTime = $build_time"
Write-Host "CommitID = $commit_id"
Write-Host "last_tag = $last_tag"

rsrc -manifest ebookdownloader_gui.manifest -ico ebookdownloader.ico -arch amd64 -o rsrc_x64.syso

go build -ldflags "-H windowsgui -w -s main.Commit=$commit_id -X 'main.BuildTime=$build_time' -X main.Version=$last_tag -linkmode external -extldflags '-static'" -o ebookdownloader_gui.exe
Copy-Item ebookdownloader_gui.exe ..\
Remove-Item  rsrc_x64.syso