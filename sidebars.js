/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

// @ts-check

/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */

const sidebars = {
  docs: [
    'home',
    {
      type: 'category',
      label: 'Getting Started',
      link: {
        type: 'doc',
        id: 'getting-started/index'
      },
      items: [
        'getting-started/create-bucketeer-account',
        'getting-started/bucketeer-dashboard',
        {
          type: 'category',
          label: 'Quickstart',
          link: {
            type: 'doc',
            id: 'getting-started/quickstart/index'
          },
          items: [
            'getting-started/quickstart/create-an-api-key',
            'getting-started/quickstart/create-your-first-feature-flag',
            'getting-started/quickstart/integrate-bucketeer'
          ]
        },
        'getting-started/bucketeer-vocabulary'
      ],
    },
    {
      type: 'category',
      label: 'Using Feature Flags',
      link: {
        type: 'doc',
        id: 'using-feature-flags/index'
      },
      items: [
        {
          type: 'category',
          label: 'Creating feature flags',
          items: [
            'using-feature-flags/creating-feature-flags/create-feature-flag',
            'using-feature-flags/creating-feature-flags/targeting',
            'using-feature-flags/creating-feature-flags/manage-variations',
            'using-feature-flags/creating-feature-flags/auto-operation',
            'using-feature-flags/creating-feature-flags/evaluate-results',
            'using-feature-flags/creating-feature-flags/other-flag-settings'
          ],
        },
        'using-feature-flags/api-keys',
        {
          type: 'category',
          label: 'Testing with flags',
          link: {
            type: 'doc',
            id: 'using-feature-flags/testing-with-flags/index'
          },
          items: [
            'using-feature-flags/testing-with-flags/goals',
            'using-feature-flags/testing-with-flags/experiments',
            'using-feature-flags/testing-with-flags/using-experiments',
          ],
        },
        'using-feature-flags/integration'
      ],
    },
    {
      type: 'category',
      label: 'SDKs',
      link: {
        type: 'doc',
        id: 'sdk/index',
      },
      items: [
        {
          type: 'category',
          label: 'Client-side',
          items: [
            'sdk/client-side/android/index',
            'sdk/client-side/ios/index',
            'sdk/client-side/javascript/index',
            'sdk/client-side/flutter/index'
          ],
        },
        {
          type: 'category',
          label: 'Server-side',
          items: [
            'sdk/server-side/go/index',
            'sdk/server-side/node-js/index'
          ],
        },
      ],
    },
    {
      type: 'category',
      label: 'Best Practices',
      link: {
        type: 'doc',
        id: 'best-practices/index',
      },
      items: [
          'best-practices/account-types',
          'best-practices/feature-flags-life-cycle',
          'best-practices/optimize-bucketeer-with-tags'
      ],
    },
    {
      type: 'category',
      label: 'Contribution Guide',
      link: {
        type: 'doc',
        id: 'contribution-guide/index',
      },
      items: [
        'contribution-guide/contributing',
        {
          type: 'category',
          label: 'Documentation style guide',
          link: {
            type: 'doc',
            id: 'contribution-guide/documentation-style-guide/index',
          },
          items: [
            'contribution-guide/documentation-style-guide/about-this-guide',
            'contribution-guide/documentation-style-guide/voice-and-tone',
            'contribution-guide/documentation-style-guide/formatting-and-organization',
            'contribution-guide/documentation-style-guide/language-and-grammar',
            'contribution-guide/documentation-style-guide/punctuation',
            'contribution-guide/documentation-style-guide/ui-elements-and-interaction',
            'contribution-guide/documentation-style-guide/links',
            'contribution-guide/documentation-style-guide/code-elements',
            'contribution-guide/documentation-style-guide/command-line-syntax',
          ],
        }
      ],
    },
  ],
}

  module.exports = sidebars;
