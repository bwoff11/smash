# Define target platforms
$platforms = @("windows/amd64", "windows/386", "darwin/amd64", "linux/amd64")

# Define output directory and ensure it exists
$outputDir = ".\builds"
if (-Not (Test-Path $outputDir)) {
    New-Item -ItemType Directory -Path $outputDir
}

# Loop through each platform and build
foreach ($platform in $platforms) {
    $split = $platform -split "/"
    $GOOS = $split[0]
    $GOARCH = $split[1]
    $outputName = "smash-$GOOS-$GOARCH"

    if ($GOOS -eq "windows") {
        $outputName += ".exe"
    }

    Write-Host "Building for $GOOS $GOARCH"

    # Execute go build
    $env:GOOS = $GOOS
    $env:GOARCH = $GOARCH
    go build -o "$outputDir\$outputName"

    # Check if build was successful
    if ($LASTEXITCODE -ne 0) {
        Write-Host "An error has occurred! Aborting the script execution..."
        exit 1
    }
}
