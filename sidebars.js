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
      className: "sidebar-bucketeer-docs",
    },
    {
      type: 'html',
      value: '<span class="sidebar-title">Getting started</span>',
      defaultStyle: true,
    },
        {
          type: 'category',
          label: "Create Bucketeer's Account",
          className: "sidebar-create-bucketeers-account",
          items: [
            'lorem/ipsum/index',
            {
              type: 'category',
              label: "Submenu test",
              items: [
                'lorem/ipsum/dolor/sit/amet/consectetur/adipiscing/index',
              ],
            },
          ],
        },
        {
          type: 'doc',
          id: 'get-your-credentials',
          label: 'Get Your Credentials',
          className: "sidebar-get-your-credentials",
        },
        {
          type: 'doc',
          id: 'bucketeer-dashboard',
          label: 'Bucketeer Dashboard',
          className: "sidebar-bucketeer-dashboard",
        },
        {
          type: 'doc',
          id: 'choose-an-sdk',
          label: 'Choose An SDK',
          className: "sidebar-choose-an-sdk",
        },
        {
          type: 'doc',
          id: 'integrate-bucketeers',
          label: 'Integrate Bucketeers',
          className: "sidebar-integrate-bucketeers",
        },
        {
          type: 'doc',
          id: 'feature-flags',
          label: 'Feature Flags',
          className: "sidebar-feature-flags",
        },
    {
      type: 'html',
      value: '<span class="sidebar-title">Using feature flag</span>',
      defaultStyle: true,
    },
        {
          type: 'category',
          label: "Creating Feature Flags",
          className: "sidebar-creating-feature-flags",
          items: [
            'lorem/ipsum/dolor/index',
          ],
        },
        {
          type: 'doc',
          id: 'api-keys',
          label: 'API Keys',
          className: "sidebar-api-keys",
        },
        {
          type: 'category',
          label: "Testing With Flags",
          className: "sidebar-testing-with-flags",
          items: [
            'lorem/ipsum/dolor/sit/index',
          ],
        },
        {
          type: 'doc',
          id: 'feature-flags-integration',
          label: 'Feature Flags Integration',
          className: "sidebar-feature-flags-integration",
        },
    {
      type: 'html',
      value: '<span class="sidebar-title">SDKS</span>',
      defaultStyle: true,
    },
        {
          type: 'category',
          label: "Client",
          className: "sidebar-client",
          items: [
            'lorem/ipsum/dolor/sit/amet/index',
          ],
        },
        {
          type: 'category',
          label: "Server",
          className: "sidebar-server",
          items: [
            'lorem/ipsum/dolor/sit/amet/consectetur/index',
          ],
        },
    {
      type: 'html',
      value: '<span class="sidebar-title">Guides</span>',
      defaultStyle: true,
    },
        {
          type: 'doc',
          id: 'contributing',
          label: 'Contributing',
          className: "sidebar-contributing",
        },
        {
          type: 'category',
          label: "Documentation Style Guide",
          className: "sidebar-documentation-style-guide",
          items: [
            'lorem/ipsum/dolor/sit/amet/consectetur/adipiscing/index',
          ],
        },
  ],
}

  module.exports = sidebars;
