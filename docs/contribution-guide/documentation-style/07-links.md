---
title: Links
slug: /contribution-guide/documentation-style/links
tags: ['contribution', 'documentation']
---

When writing link text, it's best to use brief yet descriptive phrases that offer an idea of the content you're linking to.

## Formats to structure link text

To write practical link text, choose one of these formats:

- Ensure that the text of the link matches precisely with the title or heading you are referring to, for example:
  - :+1: To check information about grammar, see section [Language and grammar](04-language-and-grammar.md).

- Write a description of the destination page to use as the link text, capitalized as if it's part of the sentence, for example:
  - :+1: You can Control your features and Make better decisions with data using [Bucketeer solution](https://bucketeer.io/).

## Guidelines to write link text

When writing link text, it is important to adhere to the following guidelines:

- When you write a complete sentence that refers the reader to another topic, introduce the link with the phrase For more information, see or For more information about..., see. Check examples:
  - :+1: For more information, see [UI elements and interaction](06-ui-elements-and-interaction.md).
  - :+1: For more information about referencing UI elements, see  [UI elements and interaction](06-ui-elements-and-interaction.md).

- Write unique, descriptive link text that makes sense without the surrounding text. Don't use phrases such as this document, this article, or click here.

    :+1: **Recommended:**

    ```html
    For more information, see <a href="/tasks">How to open and manage tasks</a>
    ```

    :x: **Not recommended:**

    ```html
    Know more? <a href="/tasks">Click here.</a>.
    ```

    :x: **Not recommended:**

    ```html
    For more information, see <a href="/tasks">this article</a>.
    ```

- Don't use a URL as link text. Instead, use the page title or a description of the page.

    :+1: **Recommended:**

    ```html
    For more information about protocols, see <a href="http://www.w3.org/Protocols/rfc2616/rfc2616.html" class="external">HTTP/1.1 RFC</a>.
    ```

    :x: **Not recommended:**

    ```html
    See the HTTP/1.1 RFC at <a href="http://www.w3.org/Protocols/rfc2616/rfc2616.html">http://www.w3.org/Protocols/rfc2616/rfc2616.html</a>.
    ```

- To ensure that readers can quickly determine if a link is relevant, use descriptive link text that accurately describes the target page. Follow these tips for improving scanning content:
  - Whenever possible, use short link text. Avoid using long link text, like a sentence or paragraph.
  - Put important words at the beginning of your link text to make it more effective.
  - If a link downloads a file, write link text that indicates this action and the file type.

    :+1: **Recommended:**

    ```html
    <a href="/readme.txt">download the README.txt file</a>
    ```
