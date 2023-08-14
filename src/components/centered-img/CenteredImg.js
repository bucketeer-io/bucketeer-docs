import React from 'react';
import useBaseUrl from '@docusaurus/useBaseUrl';

export default function Highlight({ imgURL, wSize, alt }) {
  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
      <img
        src={useBaseUrl(imgURL)}
        alt={alt}
        style={{
          width: wSize,
          // border: `${borderWidth} solid #555`,
          margin: '0 0 var(--ifm-paragraph-margin-bottom)',
          borderRadius: '14px', 
          boxShadow: '0px 0px 4px rgba(0, 0, 0, 0.15)', 
        }}
      />
    </div>
  );
}



