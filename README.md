[![CI](https://github.com/tj-actions/coverage-badge-go/workflows/CI/badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3ACI)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![Update release version.](https://github.com/tj-actions/coverage-badge-go/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3A%22Update+release+version.%22)
[![Public workflows that use this action.](https://img.shields.io/endpoint?url=https%3A%2F%2Fapi-tj-actions1.vercel.app%2Fapi%2Fgithub-actions%2Fused-by%3Faction%3Dtj-actions%2Fcoverage-badge-go%26badge%3Dtrue)](https://github.com/search?o=desc\&q=tj-actions+coverage-badge-go+path%3A.github%2Fworkflows+language%3AYAML\&s=\&type=Code)

                  👆

Generate a test coverage badge like this one for your go projects.

## coverage-badge-go

## Usage

```yaml
...
   steps:
      - name: Checkout
        uses: actions/checkout@v2
        with:
          persist-credentials: false # otherwise, the token used is the GITHUB_TOKEN, instead of your personal access token.
          fetch-depth: 0 # otherwise, there would be errors pushing refs to the destination repository.
      
      - name: Setup go
        uses: actions/setup-go@v2
        with:
          go-version: '1.14.4'

      - uses: actions/cache@v2
        with:
          path: ~/go/pkg/mod
          key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
          restore-keys: |
            ${{ runner.os }}-go-

      - name: Run Test  # Pass the `coverage.out` output to this action
        run:
          go test -v ./... -covermode=count -coverprofile=coverage.out
          go tool cover -func=coverage.out -o=coverage.out

      - name: Go Coverage Badge
        uses: tj-actions/coverage-badge-go@v1
        with:
          filename: coverage.out

      - name: Verify Changed files
        uses: tj-actions/verify-changed-files@v8.1
        id: verify-changed-files
        with:
          files: README.md

      - name: Push changes
        if: github.event == 'pull_request' && steps.changed_files.outputs.files_changed == 'true'
        uses: ad-m/github-push-action@master
        with:
          github_token: ${{ github.token }}
          branch: ${{ github.head_ref }}
```

## Features

*   Generate coverage badge from a coverage report.

## Inputs

|   Input       |    type    |  required     |  default                      |  description  |
|:-------------:|:-----------:|:-------------:|:----------------------------:|:-------------:|
| filename      |  `string`   |     `true`   | `coverage.out`               |  File containing the <br /> tests output <br />(default: "coverage.out") |

*   Free software: [MIT license](LICENSE)

If you feel generous and want to show some extra appreciation:

[![Buy me a coffee][buymeacoffee-shield]][buymeacoffee]

[buymeacoffee]: https://www.buymeacoffee.com/jackton1

[buymeacoffee-shield]: https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png

## Credits

This package was created with [Cookiecutter](https://github.com/cookiecutter/cookiecutter) using [cookiecutter-action](https://github.com/tj-actions/cookiecutter-action)

*   [gobadge](https://github.com/AlexBeauchemin/gobadge)
*   [gobinaries](https://github.com/tj/gobinaries)

## Report Bugs

Report bugs at https://github.com/tj-actions/coverage-badge-go/issues.

If you are reporting a bug, please include:

*   Your operating system name and version.
*   Any details about your workflow that might be helpful in troubleshooting.
*   Detailed steps to reproduce the bug.
