<#
.Synopsis
Makefile script in PowerShell that contains commands useful
during development of Carcassonne-Visualiser.

.Description
Available commands:
    run     Run application. Requires to pass log file name as param.
    update  Get latest versions of Carcassonne-Engine and raylib-go.

.Parameter Command
Command to execute. See Cmdlet's description for more information.

#>

[Diagnostics.CodeAnalysis.SuppressMessageAttribute(
    "PSReviewUnusedParameter",
    "",
    Justification = "Parameter is automatically provided by PowerShell which we have no control over."
)]
[CmdletBinding()]
param (
    [parameter(Mandatory = $true)]
    [ValidateSet("run", "update")]
    [String]
    $command,
    [parameter(Mandatory = $false)]
    [String]
    $filename,
    [switch]
    $help = $false
)

function Exit-On-Fail([int]$exitCode) {
    if ($exitCode) {
        Exit $exitCode
    }
}

function run() {
    if (-not ([string]::IsNullOrEmpty($filename))) {
        Write-Output "Running application..."
        & go run -tags sdl . $filename
        Exit-ON-Fail $LASTEXITCODE
    } else {
        Write-Output "Log file name not specified"
    }
}

function update() {
    Write-Output "Updating Carcassonne-Engine..."
    & go get -u "github.com/YetAnotherSpieskowcy/Carcassonne-Engine@main"
    Write-Output "Updating raylib-go"
    & go get -v -u "github.com/gen2brain/raylib-go/raylib"
    Exit-ON-Fail $LASTEXITCODE
}

$script:availableCommands = @(
    "run",
    "update"
)

if ($help) {
    Get-Help $MyInvocation.InvocationName
    exit
}

switch ($command) {
    {$script:availableCommands -contains $_} {
        & $command @Args
        break
    }
    default {
        Write-Output (
            """$command"" is not a valid command.",
            "To see available commands, type: ""$($MyInvocation.InvocationName) -help"""
        )
        break
    }
}
