[![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge\&logo=ubuntu\&logoColor=white)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Mac OS](https://img.shields.io/badge/mac%20os-000000?style=for-the-badge\&logo=macos\&logoColor=F0F0F0)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge\&logo=windows\&logoColor=white)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Public workflows that use this action.](https://img.shields.io/endpoint?style=for-the-badge\&url=https%3A%2F%2Fused-by.vercel.app%2Fapi%2Fgithub-actions%2Fused-by%3Faction%3Dtj-actions%2Fcoverage-badge-go%26badge%3Dtrue%26package_id%3DUGFja2FnZS0yOTQ5NTE3MzA0)](https://github.com/search?o=desc\&q=tj-actions+coverage-badge-go+path%3A.github%2Fworkflows+language%3AYAML\&s=\&type=Code)

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->

[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)

<!-- ALL-CONTRIBUTORS-BADGE:END -->

## coverage-badge-go

<img align="right" width="250px" src="https://user-images.githubusercontent.com/17484350/138557170-d8079b94-a517-4366-ade8-8d473e3f3f1d.jpg">

Generate a coverage badge like this one for your Golang projects without uploading results to a third party.

                  👇

[![CI](https://github.com/tj-actions/coverage-badge-go/workflows/CI/badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3ACI)
[![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)](https://github.com/tj-actions/coverage-badge-go/actions/workflows/test.yml)
[![Update release version.](https://github.com/tj-actions/coverage-badge-go/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3A%22Update+release+version.%22)

## Usage

```yaml
name: Generate code coverage badge

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
        uses: actions/checkout@v4
        with:
          persist-credentials: false # otherwise, the token used is the GITHUB_TOKEN, instead of your personal access token.
          fetch-depth: 0 # otherwise, there would be errors pushing refs to the destination repository.
      
      - name: Setup go
        uses: actions/setup-go@v4
        with:
          go-version-file: 'go.mod'

      - name: Run Test
        run: |
          go test -v ./... -covermode=count -coverprofile=coverage.out
          go tool cover -func=coverage.out -o=coverage.out

      - name: Go Coverage Badge  # Pass the `coverage.out` output to this action
        uses: tj-actions/coverage-badge-go@v2
        with:
          filename: coverage.out

      - name: Verify Changed files
        uses: tj-actions/verify-changed-files@v16
        id: verify-changed-files
        with:
          files: README.md

      - name: Commit changes
        if: steps.verify-changed-files.outputs.files_changed == 'true'
        run: |
          git config --local user.email "action@github.com"
          git config --local user.name "GitHub Action"
          git add README.md
          git commit -m "chore: Updated coverage badge."

      - name: Push changes
        if: steps.verify-changed-files.outputs.files_changed == 'true'
        uses: ad-m/github-push-action@master
        with:
          github_token: ${{ github.token }}
          branch: ${{ github.head_ref }}
```

## Signed commits

In order to create signed commits see full guide [here](https://httgp.com/signing-commits-in-github-actions/)

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

```yaml
- uses: tj-actions/coverage-badge-go@v3
  id: coverage-badge-go
  with:
    # Color of the badge - 
    # green/yellow/red 
    # Type: string
    color: ''

    # File containing the tests output
    # Type: string
    # Default: "coverage.out"
    filename: ''

    # At what percentage does the 
    # badge become green instead of 
    # yellow (default: 70) 
    # Type: string
    green: ''

    # Optional URL when you click 
    # the badge 
    # Type: string
    link: ''

    # Target file (default "README.md")
    # Type: string
    target: ''

    # Text on the left side 
    # of the badge (default: "Coverage") 
    # Type: string
    text: ''

    # Text on the right side 
    # of the badge 
    # Type: string
    value: ''

    # At what percentage does the 
    # badge become yellow instead of 
    # red (default 30) 
    # Type: string
    yellow: ''

```

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

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->

<!-- prettier-ignore-start -->

<!-- markdownlint-disable -->

<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/mdonahue-godaddy"><img src="https://avatars.githubusercontent.com/u/81647737?v=4?s=100" width="100px;" alt="Michael Donahue"/><br /><sub><b>Michael Donahue</b></sub></a><br /><a href="https://github.com/tj-actions/coverage-badge-go/issues?q=author%3Amdonahue-godaddy" title="Bug reports">🐛</a> <a href="https://github.com/tj-actions/coverage-badge-go/commits?author=mdonahue-godaddy" title="Code">💻</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->

<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
