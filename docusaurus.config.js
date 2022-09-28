// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Bucketeer Docs',
  tagline: 'Feature Flag and A/B Testing Managment platform',
  url: 'https://docs.bucketeer.io',
  baseUrl: '/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.png',
  organizationName: 'bucketeer-io', // Usually your GitHub org/user name.
  projectName: 'bucketeer', // Usually your repo name.

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en', 'ja'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          routeBasePath: '/',
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl:
            'https://github.com/bucketeer-io/bucketeer/tree/master',
          showLastUpdateTime: true,
        },
        blog: false,
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      colorMode: {
        defaultMode: 'light',
        disableSwitch: true,
        respectPrefersColorScheme: false,
      },
      navbar: {
        title: '',
        logo: {
          alt: 'Feature Flag and A/B Testing Managment platform',
          src: 'img/bucketeer-logo-white.png',
        },
        items: [
          {
            type: 'localeDropdown',
            position: 'right',
          },
          {
            href: 'https://github.com/bucketeer-io/bucketeer',
            // label: 'GitHub',
            position: 'right',
            className: 'header-github-link',
            'aria-label': 'Bucketeer - Join us on Github',
          },
          {
            href: 'https://twitter.com/bucketeer_io',
            // label: 'Twitter',
            position: 'right',
            className: 'header-twitter-link',
            'aria-label': 'Bucketeer - Follow us on Twitter',
          },
          {
            href: 'https://app.slack.com/client/T08PSQ7BQ/C043026BME1',
            // label: 'Slack',
            position: 'right',
            className: 'header-slack-link',
            'aria-label': 'Bucketeer Slack - Join the conversation',
          },
        ],
      },
      prism: {
        // theme: darkCodeTheme,
        theme: lightCodeTheme,
        additionalLanguages: [
          'groovy',
          'kotlin',
          'dart',
          'javascript',
          'swift',
          'go',
        ],
      },
    }),
};

module.exports = config;
