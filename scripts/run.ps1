# Put atton.exe with script
# run with `\run.ps1` Powershells script
# see basic statistic

while ($true)
{
    write-host 'fetching data ...'
    ./atton.exe parse
    Start-Sleep -Seconds (Get-Random -Minimum 60 -Maximum 120)
}