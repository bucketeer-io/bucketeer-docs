import React from 'react';
import Layout from '@theme-original/Footer/Layout';
import Social from './../../../theme/sections/Social/social';
import Footer from './../../../theme/sections/Footer/footer';

// <Layout {...props} />

export default function LayoutWrapper(props) {
  return (
    <>
    	<Social />
		<Footer />
    </>
  );
}
