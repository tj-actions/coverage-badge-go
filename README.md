## coverage-badge-go

<img align="right" width="250px" src="https://user-images.githubusercontent.com/17484350/138557170-d8079b94-a517-4366-ade8-8d473e3f3f1d.jpg">

Generate a coverage badge like this one for your go projects.

                  👇

[![CI](https://github.com/tj-actions/coverage-badge-go/workflows/CI/badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3ACI)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![Update release version.](https://github.com/tj-actions/coverage-badge-go/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3A%22Update+release+version.%22)

[![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?logo=ubuntu\&logoColor=white)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Mac OS](https://img.shields.io/badge/mac%20os-000000?logo=macos\&logoColor=F0F0F0)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Windows](https://img.shields.io/badge/Windows-0078D6?logo=windows\&logoColor=white)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Public workflows that use this action.](https://img.shields.io/endpoint?url=https%3A%2F%2Fapi-tj-actions1.vercel.app%2Fapi%2Fgithub-actions%2Fused-by%3Faction%3Dtj-actions%2Fcoverage-badge-go%26badge%3Dtrue)](https://github.com/search?o=desc\&q=tj-actions+coverage-badge-go+path%3A.github%2Fworkflows+language%3AYAML\&s=\&type=Code)

## Usage

```yaml
on:
  pull_request:
    branches:
      - main

jobs:
  test:
    runs-on: ubuntu-latest
    name: Update coverage badge
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

      - name: Run Test
        run:
          go test -v ./... -covermode=count -coverprofile=coverage.out
          go tool cover -func=coverage.out -o=coverage.out

      - name: Go Coverage Badge  # Pass the `coverage.out` output to this action
        uses: tj-actions/coverage-badge-go@v1.1
        with:
          filename: coverage.out

      - name: Verify Changed files
        uses: tj-actions/verify-changed-files@v8.1
        id: verify-changed-files
        with:
          files: README.md

      - name: Push changes
        if: steps.changed_files.outputs.files_changed == 'true'
        uses: ad-m/github-push-action@master
        with:
          github_token: ${{ github.token }}
          branch: ${{ github.head_ref }}
```

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|  INPUT   |  TYPE  | REQUIRED |     DEFAULT      |                                     DESCRIPTION                                      |
|----------|--------|----------|------------------|--------------------------------------------------------------------------------------|
| color    | string | false    |                  | Color of the badge -<br>green/yellow/red                                             |
| filename | string | true     | `"coverage.out"` | File containing the tests output<br>                                                 |
| green    | string | false    |                  | At what percentage does the<br>badge become green instead of<br>yellow (default: 70) |
| target   | string | false    |                  | Target file (default "README.md")                                                    |
| text     | string | false    |                  | Text on the left side<br>of the badge (default: "Coverage")<br>                      |
| value    | string | false    |                  | Text on the right side<br>of the badge                                               |
| yellow   | string | false    |                  | At what percentage does the<br>badge become yellow instead of<br>red (default 30)    |

<!-- AUTO-DOC-INPUT:END -->

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
