// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const {themes} = require('prism-react-renderer');
const lightTheme = themes.github;
const darkTheme = themes.dracula;

/** @type {import('@docusaurus/types').Config} */
const config = {
  themes: [
    [
      require.resolve("@easyops-cn/docusaurus-search-local"),
      ({
        // ... Your options.
        // `hashed` is recommended as long-term-cache of index file is possible.
        hashed: true,
        searchResultContextMaxLength: 100,
        indexBlog: false,
        // For Docs using Chinese, The `language` is recommended to set to:
        // ```
        language: ["en"],
        // ```
      }),
    ],
  ],
  plugins: [
    require.resolve('docusaurus-plugin-image-zoom'),
    [
      '@scalar/docusaurus',
      {
        label: 'API Reference',
        route: '/api',
        configuration: {
          spec: {
            // Put the URL to your OpenAPI document here:
            url: 'https://raw.githubusercontent.com/bucketeer-io/bucketeer/main/api-description/openapi.yaml'
          },
        },
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
    locales: [ 'en' ],
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
      zoom: {
        selector: '.markdown :not(em) > img',
        background: {
          light: 'rgb(255, 255, 255)',
          dark: 'rgb(50, 50, 50)'
        }
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
            to: 'changelog',
            label: 'Changelog',
            position: 'left',
            activeBasePath: "/changelog",
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
        theme: lightTheme,
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
