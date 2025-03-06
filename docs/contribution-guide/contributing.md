---
id: contributing
title: Contributing
slug: /contribution-guide/contributing
---

# Contribution Guide

We would ❤️ for you to contribute to Bucketeer and help make it better! Anyone can use, improve, and enjoy it!<br />
If you don't find your answer here, feel free to [contact us](https://app.slack.com/client/T08PSQ7BQ/C043026BME1).

## Creating an issue

Please create an issue in the [bucketeer-io/bucketeer](https://github.com/bucketeer-io/bucketeer/issues) repository if you find a problem.

## Creating a pull request

Fork the repository.

To find good issues for your first pull request, you can use our [help wanted issues](https://github.com/bucketeer-io/bucketeer/issues?q=is%3Aissue+is%3Aopen+label%3A"help+wanted") or our [good first issues](https://github.com/bucketeer-io/bucketeer/issues?q=is%3Aissue+is%3Aopen+label%3A"good+first+issue") as a reference.

:::info

Please do not force push to your PR branch after asking for a review. It will force us to review the whole PR again because we don't know what has changed.

:::

## Commit message format

We are following the [Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0) message rules.<br />
This format leads to easier-to-read commit history and enables us to generate changelogs and follow [semantic versioning](https://semver.org).

:::tip

The commit message is used for our release changelog. Please write a message that is easier to understand from the user's point of view.

:::

### Types

Must be one of the following:

- **build:** Changes that affect the build system or external library dependencies
- **chore:** Other changes that don't modify src or test files
- **ci:** Changes to the CI configuration files and scripts
- **docs:** Documentation only changes
- **feat:** A new or a feature update
- **fix:** A bug fix
- **perf:** A code change that improves performance
- **refactor:** A code change that neither fixes a bug nor adds a feature
- **revert:** Reverts a previous commit
- **test:** Adding missing or correcting existing tests

:::info

For BREAKING CHANGES commits, you must append a `!` after the type, which introduces a breaking API change (correlating with MAJOR in Semantic Versioning).<br />
E.g. `feat!: new API interface 2.0`

:::

### Subject

The subject contains the description of the change:

- Use the imperative, present tense: "change" not "changed" nor "changes"
- Don't capitalize the first letter
- No dot (.) at the end

### Example

```
feat: add webhook feature
^--^  ^-----------------^
|     |
|     +-> Subject in present tense. Not capitalized.
|
+-------> Type: build|chore|ci|docs|feat|fix|perf|refactor|revert|test
```

## Code reviews

All submissions, including submissions by project members, require review. We use GitHub pull requests for this purpose. Consult [GitHub Help](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests) for more information on using pull requests.
