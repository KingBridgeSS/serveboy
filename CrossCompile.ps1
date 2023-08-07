$env:GOOS = 'windows'; go build  -ldflags "-s -w" -o sb.exe;
$env:GOOS = 'linux'; go build  -ldflags "-s -w" -o sb;
$env:GOOS = '';