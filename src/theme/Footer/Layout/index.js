import React from 'react';
import Layout from '@theme-original/Footer/Layout';
import Social from './../../../theme/sections/social/social';
import Footer from './../../../theme/sections/footer/footer';

// <Layout {...props} />

export default function LayoutWrapper(props) {
  return (
    <>
    	<Social />
		<Footer />
    </>
  );
}
