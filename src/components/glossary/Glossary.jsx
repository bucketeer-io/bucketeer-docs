import React, { useState, useRef, useEffect } from 'react';
import style from './style.css'

export default function Glossary({word, content}) {
  const [isHovered, setIsHovered] = useState(false);
  const [hoveredWord, setHoveredWord] = useState('');
  const [boxPosition, setBoxPosition] = useState({ top: 0, left: 0 });
  const wordRef = useRef(null);

  const handleMouseEnter = (word, event) => {
    const wordRect = wordRef.current.getBoundingClientRect();
    setHoveredWord(word);
    /*setBoxPosition({
      top: event.clientY, // Adjust the values as needed
      left: event.clientX, // Adjust the values as needed
    });*/
    setBoxPosition({
      top: wordRect.top - 45, // Adjust the values as needed
      left: wordRect.left - 240, // Adjust the values as needed
    });
    setIsHovered(true);
  };

  const handleMouseLeave = () => {
    setIsHovered(false);
  };

  return (
    <span>
        <span
          onMouseEnter={(e) => handleMouseEnter('hoveredWord', e)}
          onMouseLeave={handleMouseLeave}
          style={{ textDecoration: 'underline', cursor: 'pointer' }}
          ref={wordRef}
        >
          Word
        </span>{' '}


      {isHovered && (
        <div
          style={{
            position: 'absolute',
            background: 'white',
            border: '1px solid #ccc',
            width: '200px', // Adjust the width as needed
            padding: '10px',
            top: boxPosition.top,
            left: boxPosition.left,
            zIndex: 999,
          }}
        >
          <p>{content}</p>
        </div>
      )}
    </span>
  );
};
