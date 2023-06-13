import React from 'react';
import style from './style.css'
import { useHistory } from '@docusaurus/router';

export default function Highlight({redirect,title, info}) {
  const history = useHistory();

  const handleClick = () => {
    history.push(redirect)
  }

  return (
    <div className='button' onClick={handleClick}>
      <h3>{title}</h3>
      <p>{info}</p>
    </div>
  );
}