import React from 'react';
import style from './style.css'

export default function Highlight({ children }) {
  return (
    <div className='button-shelf'>
      { children }
    </div>
  );
}