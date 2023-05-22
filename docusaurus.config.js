// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Bucketeer - Feature Flag Management and A/B Testing platform',
  tagline: 'Feature Flag and A/B Testing Managment platform',
  url: 'https://bucketeer.io',
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
    locales: ['en', 'ja'],
  },
  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: false,
        blog: {
          routeBasePath: '/blog',
        },
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
      metadata: [
        {property: 'og:title', content: 'Bucketeer - Feature Flag Management and A/B Testing platform'},
        {property: 'og:description', content: 'Bucketeer is an open-source platform created to help teams make better decisions, reduce deployment lead time and release risk through feature flags.'},
        {property: 'og:image', content: 'https://bucketeer.io/img/bucketeer-logo.png'},
        {property: 'og:type', content: 'website'},
        {property: 'og:url', content: 'https://bucketeer.io'},
        {property: 'og:site_name', content: 'Bucketeer'},
        {itemprop: 'name', content: 'Bucketeer'},
        {itemprop: 'description', content: 'Bucketeer is an open-source platform created to help teams make better decisions, reduce deployment lead time and release risk through feature flags.'},
        {name: 'twitter:card', content: 'summary_large_image'},
        {name: 'twitter:site', content: '@bucketeer_io'},
        {name: 'twitter:title', content: 'Bucketeer'},
        {name: 'twitter:description', content: 'Bucketeer is an open-source platform created to help teams make better decisions, reduce deployment lead time and release risk through feature flags.'},
        {name: 'twitter:image', content: 'https://bucketeer.io/img/bucketeer-ogp-logo.png'},
        {name: 'twitter:image:alt', content: 'Bucketeer - Feature Flag Management and A/B Testing platform'}
      ],
      colorMode: {
        defaultMode: 'light',
        disableSwitch: true,
        respectPrefersColorScheme: false,
      },
      navbar: {
        title: '',
        logo: {
          alt: 'Feature Flag and A/B Testing Managment platform',
          src: 'img/bucketeer-logo.png',
        },
        items: [
          {
            to: 'https://docs.bucketeer.io',
            target: '_self',
            label: 'Documentation',
            position: 'right',
            // className: '',
            'aria-label': 'Bucketeer - Documentation',
          },
          {
            to: '#',
            target: '_self',
            label: 'Live Demo (coming soon)',
            position: 'right',
            className: 'link-disable',
            'aria-label': 'Bucketeer - Live Demo',
          },
          {
            to: 'blog',
            label: 'Blog',
            position: 'right',
            // className: '',
            'aria-label': 'Bucketeer - Blog',
          },
          {
            to: 'https://github.com/bucketeer-io/bucketeer',
            target: '_self',
            label: 'Github',
            position: 'right',
            // className: '',
            'aria-label': 'Bucketeer - Join us on Github',
          },
        ],
      },
      footer: {
        copyright: `©${new Date().getFullYear()} The Bucketeer Authors All Rights Reserved. <a href="https://github.com/bucketeer-io/bucketeer/blob/master/LICENSE" target="_blank">Privacy Policy</a>`,
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
