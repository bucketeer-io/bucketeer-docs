---
title: Code elements
# sidebar_position: 
slug: /documentation-style-guide/code-elements
tags: ['contribuition','documentation']
---


In this section, you will find guidelines for using standard formatting codes correctly. Follow these recommendations and avoid common errors and make your content assertive.

## Code in text

When writing regular text, use a code font to indicate anything related to coding rather than exclusively using it in code samples.

In HTML, use the `<code>` element.
In Markdown, use backticks <code>`</code>.

Below are some examples of specific items that you can write in code font:

- Attribute names and values
- Command output (for example, `ping 192.160.100.2`)
- Data types
- Defined (constant) values for an element or attribute
- DNS record types
- Enum (enumerator) names
- Environment variable names
- Element names (XML and HTML)

    -Place angle brackets (`<>`) around the element name; you might have to escape the angle brackets to make them appear in the document.

- Filenames, filename extensions (if used), and paths
- Folders and directories
- HTTP verbs
- HTTP status codes
- HTTP content-type values
- IAM role names (for example, `functions/certain.admin`)
- Language keywords
- Method and function names
- Namespace aliases
- Placeholder variables
- Query parameter names and values
- Strings (such as URLs or domain names)

## Code samples

Here, you will find basic guidelines on how to format code samples:

- Don't use tabs to indent code; use spaces only.

- Wrap lines at 80 characters.

  - Wrap code lines at a smaller character count for easier reading on narrow browsers or printed documents.

- Mark code blocks as preformatted text. In HTML, use a `<pre>` element; in Markdown, indent every line of the code block by four spaces.

- Indicate omitted code using three dots and no spaces (...). Do not use the ellipsis character (…). If the omission is one or more lines long, place the three dots on their own line. Do not format a sample with omitted code as a click-to-copy code block.

- When including code samples, it's helpful to introduce them with a sentence or paragraph.

  *Option 01. Use a colon if the sample follows the intro, and try to keep the same intro standard, for example:*

  - The following code sample shows how to use the `get` method:

    ```curl
    curl https://docs.bucketeer.io/
    ```

    For information about other methods, see [link].

  *Option 02. Use a period if there's more material (such as a note) between the introduction and the sample, for example:*

  - The following code sample shows how to use the `get` method. For information about other methods, see [link].

    ```curl
    curl https://docs.bucketeer.io/
    ```

> **Note**
> If you have questions about a particular programming language formatting, we recommend check a [coding-style guide](https://google.github.io/styleguide/).
