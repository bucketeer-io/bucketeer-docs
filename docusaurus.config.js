// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {

  plugins: [
    [
      require.resolve("@cmfcmf/docusaurus-search-local"),
      {
        indexBlog: false,
      },
    ],
  ],

  title: 'Bucketeer Docs',
  tagline: 'Feature Flag and A/B Testing Managment platform',
  url: 'https://docs.bucketeer.io',
  baseUrl: '/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'bucketeer-io', // Usually your GitHub org/user name.
  projectName: 'bucketeer', // Usually your repo name.

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: [ 'en', 'ja' ],
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
            'https://github.com/bucketeer-io/bucketeer-docs/tree/master',
          showLastUpdateTime: true,
        },
        blog: false,
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
        gtag: {
          trackingID: 'G-WMMC2THNMZ',
          anonymizeIP: true,
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
          className: 'header-logo',
        },
        items: [
          {
            to: '/',
            label: 'Home',
            position: 'left',
            activeBaseRegex: "/$",
          },
          {
            to: 'getting-started',
            label: 'Getting Started',
            position: 'left',
            activeBaseRegex: "/getting-started",
          },
          {
            to: 'sdk',
            label: 'SDKs',
            position: 'left',
            activeBaseRegex: "/sdk",
          },
          {
            to: 'contribution-guide/contributing',
            label: 'Contributing',
            position: 'left',
            activeBasePath: "contribution-guide/",
          },
          {
            type: 'search',
            position: 'right',
          },
          // {
          //   type: 'localeDropdown',
          //   position: 'right',
          // },
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
        // theme: require("prism-react-renderer/themes/vsDark"),
        // theme: require("prism-react-renderer/themes/shadesOfPurple"),
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
