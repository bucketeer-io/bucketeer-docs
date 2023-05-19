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
        'getting-started/get-your-credentials',
        'getting-started/bucketeer-dashboard',
        'getting-started/choose-sdk',
        'getting-started/integrate-bucketeers'
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
      label: 'Contribution Guide',
      link: {
        type: 'doc',
        id: 'contribution-guide/index',
      },
      items: [
        'contribution-guide/contributing'
      ],
    },
  ],
}

  module.exports = sidebars;
