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
    {
      type: 'doc',
      id: 'bucketeer-docs',
      label: 'Bucketeer Docs',
      className: 'sidebar-bucketeer-docs',
    },
    {
      type: 'html',
      value: "<span class='sidebar-title'>Getting started</span>",
      defaultStyle: true,
    },
    {
      type: 'doc',
      id: 'getting-started/introduction',
      label: 'Introduction',
      className: 'sidebar-overview',
    },
    {
      type: 'doc',
      id: 'getting-started/create-an-account',
      label: 'Create an account',
      className: 'sidebar-account',
    },
    {
      type: 'category',
      label: "Quickstart",
      link: {
        type: 'doc',
        id: 'getting-started/quickstart/index',
      },
      className: 'sidebar-quickstart',
      items: [
        'getting-started/quickstart/create-an-api-key',
        'getting-started/quickstart/create-your-first-flag',
        'getting-started/quickstart/integrate-bucketeer'
      ],
    },
    {
      type: 'doc',
      id: 'getting-started/bucketeer-vocabulary',
      label: 'Bucketeer Vocabulary',
      className: 'sidebar-vocabulary',
    },
    // {
    //   type: 'doc',
    //   id: 'get-your-credentials',
    //   label: 'Get Your Credentials',
    //   className: 'sidebar-get-your-credentials',
    // },
    // {
    //   type: 'doc',
    //   id: 'bucketeer-dashboard',
    //   label: 'Bucketeer Dashboard',
    //   className: 'sidebar-bucketeer-dashboard',
    // },
    // {
    //   type: 'doc',
    //   id: 'choose-an-sdk',
    //   label: 'Choose An SDK',
    //   className: 'sidebar-choose-an-sdk',
    // },
    // {
    //   type: 'doc',
    //   id: 'integrate-bucketeers',
    //   label: 'Integrate Bucketeers',
    //   className: 'sidebar-integrate-bucketeers',
    // },
    // {
    //   type: 'doc',
    //   id: 'feature-flags',
    //   label: 'Feature Flags',
    //   className: 'sidebar-feature-flags',
    // },
    {
      type: 'html',
      value: "<span class='sidebar-title'>feature flags</span>",
      defaultStyle: true,
    },
    {
      type: 'doc',
      id: 'feature-flags/index',
      label: 'Overview',
      className: 'sidebar-overview',
    },
    {
      type: 'category',
      label: 'Creating Feature Flags',
      className: 'sidebar-creating-feature-flags',
      items: [
        'feature-flags/creating-feature-flags/create-feature-flag',
        'feature-flags/creating-feature-flags/targeting',
        'feature-flags/creating-feature-flags/manage-variations',
        'feature-flags/creating-feature-flags/auto-operation',
        'feature-flags/creating-feature-flags/trigger',
        'feature-flags/creating-feature-flags/evaluate-results',
        'feature-flags/creating-feature-flags/settings-and-history'
      ],
    },
    {
      type: 'doc',
      id: 'feature-flags/api-keys',
      label: 'API Keys',
      className: 'sidebar-api-keys',
    },
    {
      type: 'doc',
      id: 'feature-flags/audit-logs',
      label: 'Audit Logs',
      className: 'sidebar-audit-logs',
    },
    
    {
      type: 'html',
      value: "<span class='sidebar-title'>Experimentation</span>",
      defaultStyle: true,
    },
    {
      type: 'category',
      label: 'Testing With Flags',
      className: 'sidebar-testing-with-flags',
      link: {
        type: 'doc',
        id: 'experimentation/index',
      },
      items: [
        'experimentation/goals',
        'experimentation/experiments',
        'experimentation/using-experiments',
      ],
    },
    // {
    //   type: 'doc',
    //   id: 'feature-flags-integration',
    //   label: 'Feature Flags Integration',
    //   className: 'sidebar-feature-flags-integration',
    // },
    {
      type: 'html',
      value: "<span class='sidebar-title'>SDKS</span>",
      defaultStyle: true,
    },
    {
      type: 'doc',
      id: 'sdk/overview',
      label: 'Overview',
      className: 'sidebar-overview',
    },
    {
      type: 'category',
      label: 'Client',
      className: 'sidebar-client',
      items: [
        'sdk/client-side/android/index',
        'sdk/client-side/ios/index',
        'sdk/client-side/flutter/index',
        'sdk/client-side/javascript/index',
      ],
    },
    {
      type: 'category',
      label: 'Server',
      className: 'sidebar-server',
      items: [
        'sdk/server-side/go/index',
        'sdk/server-side/node-js/index',
      ],
    },
    {
      type: 'html',
      value: "<span class='sidebar-title'>Integration</span>",
      defaultStyle: true,
    },
    {
      type: 'category',
      label: 'Tools',
      link: {
        type: 'doc',
        id: 'integration/index',
      },
      className: 'sidebar-overview',      
      items: [
        'integration/pushes',
        'integration/notifications',
      ],
    },
    

    {
      type: 'html',
      value: "<span class='sidebar-title'>Best practices</span>",
      defaultStyle: true,
    },
    {
      type: 'doc',
      id: 'best-practices/optimize-bucketeer-with-tags',
      label: 'Optimize with tags',
      className: 'sidebar-contributing',
    },
    {
      type: 'html',
      value: "<span class='sidebar-title'>Contribution</span>",
      defaultStyle: true,
    },
    {
      type: 'doc',
      id: 'contribution-guide/contributing',
      label: 'Contributing',
      className: 'sidebar-contributing',
    },
    {
      type: 'category',
      label: 'Documentation Style',
      link: {
        type: 'doc',
        id: 'contribution-guide/documentation-style/index',
      },
      className: 'sidebar-documentation-style-guide',
      items: [
        'contribution-guide/documentation-style/consistency',
        'contribution-guide/documentation-style/voice-and-tone',
        'contribution-guide/documentation-style/formatting-and-organization',
        'contribution-guide/documentation-style/language-and-grammar',
        'contribution-guide/documentation-style/punctuation',
        'contribution-guide/documentation-style/ui-elements-and-interaction',
        'contribution-guide/documentation-style/links',
        'contribution-guide/documentation-style/code-elements',
        'contribution-guide/documentation-style/command-line-syntax',
      ],
    },
  ],
}

module.exports = sidebars;
