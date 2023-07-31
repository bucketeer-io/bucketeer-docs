---
title: Formatting and Organization
slug: /contribution-guide/documentation-style/formatting-and-organization
tags: ['contribution', 'documentation']
---


This section covers a range of elements that contribute to your content's overall structure and presentation. By following these guidelines, you'll ensure your content is well-organized, visually appealing, and consistent.

## Date and time

Follow the standards below to ensure a consistent formatting convention for date and time.

### Date

- Spell out the names of months and days of the week in total. Only abbreviate if there is a space limitation. For example: January 19, 2017.

- If including the day of the week, add it before the month. For example: Tuesday, April 27, 2021. 

### Time

- Use the 12-hour clock.

- Use hyphens in time ranges. Don't add spaces before or after the hyphens, such as 5-10 minutes ago.

- Capitalize AM and PM, and leave one space between it and the time. For example: 3:45 PM.

- Remove the minutes from round hours, such as 3 PM.

### Dates and time together

When expressing the date and time together, it is best to mention the date first and then the time:

- 2017-04-15 at 3 PM
- May 4, 2009, at 6 PM

### Time zones

In cases where you need to use a time zone, use the following guidelines:

- Let the reader know if the time is local to their time.

  - 10 AM your local time.

- If using a specific time zone, spell out the region and include the [UTC or GMT label](https://www.worldtimeserver.com/learn/utc-vs-gmt/).

  - US and Canadian Pacific Standard Time (UTC-8)
  - US and Canadian Pacific Daylight Time (UTC-7)

## Images and graphs

Images and graphs are powerful resources to assist in understanding processes and technologies. Otherwise, observing some critical points is crucial, for these resources do not negatively compromise your content.

- *Choose simple and clear images* and remove anything unnecessary, and be careful with crops. If the image has important information, add [alt text](https://accessibility.huit.harvard.edu/describe-content-images).

- *Remember that your customer's screen may be small*, so keep your content easy to read.

- *Be detail-oriented when creating infographics*. Provide a way for viewers to zoom in on specific areas. Or, present individual sections of the infographic within the text, and provide a link to the full infographic.

- *Simplify charts and graphs* so readers can easily read the whole thing on a small screen.

- *Ensure it can be easy to edit when using text in graphics*. Automatic translation software cannot translate it, so it would be best to use captions or describe the graphic in the text if possible.

## Use of examples

To provide an example in a sentence, refer to the guidance below.

- At the end of a sentence, use a comma or an em dash.

| :x: &nbsp; Not recommended                               |
|----------------------------------------------------------|
| Insert a value for the key Id, for example, user-10.  |

 :x: &nbsp; Not recommended                               |
|----------------------------------------------------------|
| Insert a value for the key Id, for example; user-10.  |

| :+1: &nbsp; Recommended                                  |
|----------------------------------------------------------|
| Insert a value for the key Id, such as user-10.   |

| :+1: &nbsp; Recommended                                  |
|----------------------------------------------------------|
| Insert a value for the key Id. For example: user-10.   |

| :+1: &nbsp; Recommended                                  |
|----------------------------------------------------------|
| Insert a value for the key Id (e.g., user-10).   |


- In the middle of a sentence, use parentheses if the example is short. Otherwise, rewrite the sentence.

| :x: &nbsp; Not recommended                                                                                           |
|----------------------------------------------------------------------------------------------------------------------|
| Insert a value for the key Id (for example, a string including your user Id like user-10), and then click **Next**.  |

| :+1: &nbsp; Recommended                                 |
|---------------------------------------------------------|
| Insert a value for the key Id (for example, user-10), and then click **Next**  |

## Headings

Headings can provide structure and visual cues that make it easier for readers to scan and navigate. By dividing your text into logical sections, the extra spacing and distinct font sizes to headings will help readers find entry points and locate information efficiently.

Essential things to consider while writing headings:

- *Use brief headings*. It Is crucial to prioritize the most significant idea and place it at the beginning. Avoid exceeding 6-8 words in headings.

- *Be specific and detailed when creating headings*, especially second-level headings which should provide even more specific information.

- *Concentrate on what matters to customers*, and use customer-friendly language in headings. Avoid mentioning products or commands (unless it's the goal) and focus on what customers need to know or accomplish.

- *Format headings with sentence-style capitalization*. Capitalize the first word, proper nouns, and the first word after a colon (if there is one). Keep everything else in lowercase. Some examples below:

  - Control your features with Bucketeer
  - Speed up the deployment process
  - Find a solution
  - Use Bayesian probabilities on A/B tests
  - Bucketeer: Better decisions with data

## Notes and notices

Use a notice to provide useful or crucial information that is not part of the main text. However, has been noticed that readers often ignore elements such as notices that do not pertain to their interests or focus. After writing information in regular text, evaluate if it is necessary to highlight it as a notice.

- *Don't use many notices on a single page*. It can make them less noticeable. Find alternative ways to communicate information, especially if there are multiple notices in a row.

- *Avoid grouping multiple notices*. If necessary, reorganize the content for clarity.

### Notices types

When you decide to provide notice, try using these types:

- **Note**. It is a quick suggestion or tip. Offers helpful information that may be optional but can still be valuable to the reader. For example:

  > **Note**
  >
  > When making API calls, always include the appropriate authentication credentials in the request headers to ensure secure and authorized access to the API endpoints.

- **Warning**. It means "Don't do this" or that this step might be irreversible, such as leading to permanent data loss. Ignoring the warning could result in financial loss, loss of work, or a security breach for the reader. For example:

  > **Warning**
  >
  > Avoid including passwords as command line arguments when conducting A/B software tests with APIs. Storing sensitive information in this manner poses a significant security risk. Instead, utilize secure authentication methods and token-based access to protect confidential data.

## Numbers

To avoid confusion, always use consistent numbering. If you are referring to numbers displayed in the user interface (UI), replicate them exactly. For all other content, refer to the guidelines provided.

In body text, spell out whole numbers from zero through nine, and use numerals for 10 or greater. It's OK to use numerals for zero through nine when you have limited space, such as in tables and UI.

- 10 configuration steps
- eight repositories
- 3,210,000
- 1,000

If one item requires a numeral, use numerals for all the other items of that type. For example: One guide has 20 section, second has 5 sections, and the third has only 1 section.

When two numbers that refer to different things must appear together, use a numeral for one and spell out the other. For example: Ten 10-section guides.

Don't start a sentence with a numeral. Add a modifier before the number, or spell the number out if you can't rewrite the sentence. It's OK to start list items with numerals—use your judgment.

- More than 10 APIs have been launched.
- Eleven APIs have been launched.

Use commas in numbers that have four or more digits.

- 1,500
- 150,000
- 1,093 MB

## Tables

Tables simplify complex information by presenting it in a structured format. Typically, tables contain two or more rows and two or more columns. Avoid using tables for presenting a list of similar items; instead, use a list.

Keep in mind some things when structuring a table:

- *Make sure the table is clear*. You can use a title or brief intro.

- *Put identifying information in the leftmost column* of a table for easy understanding.

- *To ensure consistency, align the entries in a table*. This means using the same format for all items within a column, such as using only nouns or phrases that start with a verb.

- *Use sentence-style capitalization* for the table title and each column header. Use sentence-style capitalization for the text in cells unless there's a reason not to (for example, keywords that must be lowercase).

- *Don't leave cells blank or use em dashes*. Use "Not applicable" or "None" instead.

- *For better design, use responsive content*. Keep limited columns and one-line text per cell.

- *Balance row height* by increasing the width of text-heavy columns and reducing the width of columns with minimal text.

### Simple table example

| Command               | Description                                    |
|-----------------------|------------------------------------------------|
| `git status`          | List all new or modified files                 |
| `git diff`            | Show file differences that haven't been staged |
| `git power`           | Not applicable                                 |

## Text-formatting

This page offers a summary and quick reference for various general text-formatting conventions.

### Bold

- Use bold, `<b>` or `**`, for user interface elements and at the beginning of notices.

### Italic

- Use italics, `<i>` or `*`, to highlight words or phrases when defining terms or referring to words.

- Parameter names, data, and count. For example, when you refer to the parameters of a method like *keyAccount*.

- Mathematical variables and version variables. For example, *x + y = 3*, version 1.4.*x*.

- To indicate semantic emphasis in HTML, use the `<em>` element, which renders as italics in most contexts. To indicate emphasis in Markdown, use <code> * </code>, which render as italics; you can't do semantic tagging in Markdown.

### Code font

- Use `<code>` in HTML or <code>`</code> in Markdown to apply a monospace font and other styling to code in text, inline code, and user input.

- Use code blocks, `<pre>` or <code>```</code>, for code samples or other blocks of code.

- Use code font to mark up code, such as filenames, class names, method names, HTTP status codes, console output, and placeholders.
