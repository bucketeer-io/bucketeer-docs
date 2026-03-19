---
name: update-documentation
description: How to add or update documentation pages
---

When the user asks you to modify the documentation, carefully determine whether you are **Adding a New Document** or **Updating an Existing Document**, and follow the corresponding steps to ensure the Docusaurus site structure remains intact.

> [!IMPORTANT]
> When generating or updating documentation, **do not use icons or emojis (emotes)** in the text content. Keep the text professional and clean.

### Scenario A: Adding a New Document

(e.g., adding a new OpenFeature provider, a new SDK, or a tutorial)

1. **Create the Markdown File**
   - Create a `.md` or `.mdx` file in the appropriate subdirectory within `docs/` or `changelog/`.
   - Ensure the file has Docusaurus Front Matter with the `title` and `slug` correctly defined.
   - Example:
     ```markdown
     ---
     title: Swift (iOS)
     slug: /open-feature/swift
     ---
     ```

2. **Update the Sidebar (`sidebars.js`)**
   - Locate `sidebars.js` in the project root.
   - Find the correct array (`docs`, `changelog`) and the appropriate label/category (e.g., `Client & Mobile`, `Server & Edge`, or `OpenFeature Providers`).
   - Append the new file path (relative to the `docs/` folder, omitting the file extension) to the `items` array.
   - Example addition: `'open-feature/swift/index'`

3. **Update Index and Landing Pages**
   - **`docs/bucketeer-docs.mdx`**: If it's a new SDK or Provider, add an HTML anchor element pointing to the new slug (e.g., `<a href="/open-feature/swift">...</a>`) in the appropriate flex/grid list. Ensure you include the correct brand icon.
   - **`docs/sdk/index.mdx`**: For Bucketeer SDKs, update the features checkmark table. For OpenFeature SDKs, update the HTML link list at the bottom.
   - **`docs/open-feature/overview.mdx`**: If it's an OpenFeature provider, append it to the internal document list.

4. **Verify the Build Locally**
   - Ensure you did not introduce broken markdown links or syntax errors.
   - // turbo
   - Run: `npm run build`
   - If the build fails with broken links, fix the links in the markdown files or sidebars.js.

### Scenario B: Updating an Existing Document

(e.g., modifying an existing SDK guide, adding a new parameter to the API reference, or fixing a typo)

1. **Locate the Document**
   - Use `grep_search` or `find_by_name` to locate the target `.md` or `.mdx` file within the `docs/` or `changelog/` directories.
   - Use `view_file` to read the existing content and understand its structure.

2. **Make the Edit Safely**
   - Apply your changes to the markdown file.
   - **Crucial Warning:** If you modify a Markdown Heading (e.g., changing `## Setup Context` to `## Initialize Context`), you MUST use `grep_search` to find and update any other documentation pages that link to that specific anchor (`#setup-context`). Docusaurus will fail the build if an anchor link is broken.

3. **Verify the Build Locally**
   - Ensure your updates did not introduce syntax errors or broken links.
   - // turbo
   - Run: `npm run build`
   - If the build fails (especially with 'broken markdown link' errors), immediately locate the source of the broken link and fix it.
