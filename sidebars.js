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
    'getting-started',
    {
      type: 'html',
      value: '<span class="sidebar-title">SDKs</span>', // The HTML to be rendered
      defaultStyle: true, // Use the default menu item styling
    },
    // {
      // type: 'category',
      // label: 'SDKs',
      // link: {
      //   type: 'doc',
      //   id: 'sdk/index',
      // },
      // items: [
        {
          type: 'category',
          label: 'Client-side',
          className: "sidebar-client-side",
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
          className: "sidebar-server-side",
          items: [
            'sdk/server-side/go/index',
            'sdk/server-side/node-js/index'
          ],
        },
    //   ],
    // },
    {
      type: 'html',
      value: '<span class="sidebar-title">Guides</span>', // The HTML to be rendered
      defaultStyle: true, // Use the default menu item styling
    },
    {
      type: 'category',
      label: 'Contribution Guide',
      className: "sidebar-contribution-guide",
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
