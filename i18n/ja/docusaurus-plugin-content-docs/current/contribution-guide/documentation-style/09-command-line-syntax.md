---
title: Command-line syntax
slug: /contribution-guide/documentation-style/command-line-syntax
tags: ['contribution', 'documentation']
---


On this section, you will understand how to document command-line commands and their respective arguments.

Here are some recommended best practices:

- Include an inline link for easy access to command reference, for example:

    To generate a new SSH Key, use the `bucketeer ssh-keygen` [command link]:

    ```command-line
    bucketeer ssh-keygen
    ```

- To do each task correctly, find the required arguments. Use as few optional arguments as possible to avoid extra documentation.

- Provide a click-to-copy command example that the reader doesn't need to edit after they copy it. If possible, include only runnable code and placeholder variables in the click-to-copy example.

- Some command examples contain optional arguments, mutually exclusive arguments, or repeated arguments that are indicated by square brackets ([]), pipes (|), braces ({}), and ellipses (...). These characters can break commands if they're not first removed. For that reason, avoid using these arguments in click-to-copy examples.

## Format a command

If you need to highlight a block of code, like a long command or a code example, you can use the following formatting:

- In HTML, use the `<pre>` element.
In Markdown, use a code fence <code>```</code>.

- To format a command with multiple elements, do the following:
If a line is longer than 80 characters, add a line break before certain characters, like a single hyphen, double hyphen, underscore, or quotation marks. After the first line, indent each line by four spaces to vertically align each line that follows a line break.

- To split a command line with a line break, ensure that every line, except the last one, ends with the command-continuation character. Commands that don't have the command-continuation character don't work.

  - Linux or Cloud Shell: A backslash typically preceded with a space ( \)

  - Windows: A caret preceded with a space ( ^)

## Command arguments

To show that an argument is optional:

1. Use square brackets.
If there are several optional arguments, put each in its own set of square brackets.
2. Avoid using optional arguments in click-to-copy code examples.

3. In the following example, `GROUP` is required, but `GLOBAL_FLAG` and `FILENAME` are optional:

```command-line
bucketeer dns GROUP [GLOBAL_FLAG] [FILENAME]
```
