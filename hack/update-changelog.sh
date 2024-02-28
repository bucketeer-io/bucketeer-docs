#!/bin/bash

# This script is triggered from other repositories in the Bucketeer Organization
# to update the CHANGELOG SDKs pages.
# In case of changes in the template, we also might need to update the trigger side.

# Envs
CHANGELOG_TEMP_FILE=changelog_temp.md

echo "DOC_FILENAME = $DOC_FILENAME"
echo "DOC_FILEPATH = $DOC_FILEPATH"
echo "DOC_TITLE = $DOC_TITLE"
echo "DOC_SLUG = $DOC_SLUG"
echo "CHANGELOG_URL = $CHANGELOG_URL"
echo "PWD = $PWD"

# Download the CHANGELOG.md
curl -f -o $CHANGELOG_TEMP_FILE $CHANGELOG_URL

# Remove the first line (The H1 # CHANGELOG)
sed -i'' -e '1d' $CHANGELOG_TEMP_FILE

# Reformat the data from (YYYY-MM-DD) to (YYYY/MM/DD)
sed -i'' -e 's/\(\([0-9]\{4\}\)-\([0-9]\{2\}\)-\([0-9]\{2\}\)\)/\2\/\3\/\4/g' $CHANGELOG_TEMP_FILE

# Write the new CHANGELOG.md
file=$(cat << EOF
---
title: ${DOC_TITLE}
slug: ${DOC_SLUG}
toc_max_heading_level: 2
---

<style>
  {\`
    h2:not(:first-of-type) {
      border-top: 2px solid #ddd7e9;
      padding-top: 40px;
    }
  \`}
</style>

$(cat $CHANGELOG_TEMP_FILE)

EOF
)

# Override the current file
# The file path MUST exist or it will fail.
echo "$file" > ./${DOC_FILEPATH}/${DOC_FILENAME}
cat ./${DOC_FILEPATH}/${DOC_FILENAME}

# Delete the temp file
rm $CHANGELOG_TEMP_FILE

echo "File ./${DOC_FILENAME} updated."
