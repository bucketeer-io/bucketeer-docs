import React from 'react';
import useBaseUrl from '@docusaurus/useBaseUrl';

export default function Highlight({imgURL,wSize, alt, borderWidth}) {
  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
      <img src={useBaseUrl(imgURL)} alt={alt} style={{ width: wSize, border: `${borderWidth} solid #555`, margin: '0 0 var(--ifm-paragraph-margin-bottom)' }}/>
    </div>
  );
}

