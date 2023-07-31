import React, { useState } from 'react';
import style from './style.css'

export default function FilterComponent({ dataList }) {


  const [ filterValue, setFilterValue ] = useState('');
  const [ selectValue, setSelectValue ] = useState('');
  const data = dataList;
  const availableLetters = [ ...new Set(dataList.map(item => item.name.charAt(0).toUpperCase())) ];

  const handleFilterChange = (event) => {
    setFilterValue(event.target.value);
  };

  const handleOptionChange = (event) => {
    const optionValue = event.target.value;
    setSelectValue(optionValue);
  };

  return (
    <div className='filterComponent'>
      <div className='inputs'>
        <input
          type="text"
          value={ filterValue }
          onChange={ handleFilterChange }
          placeholder='Search vocabulary'
        />
        <div>
          <label htmlFor="selectLetter">Letter selector:</label>
          <select value={ selectValue } onChange={ handleOptionChange }>
            <option value="">All</option>
            { availableLetters.map((option) => (
              <option key={ option } value={ option }>
                { option }
              </option>
            )) }
          </select>
        </div>
      </div>
      { data && data
        .filter((item) => item.name.toLowerCase().includes(filterValue.toLowerCase()))
        .filter((item) => item.name.charAt(0).toUpperCase().includes(selectValue))
        .map((item) => (
          <div className='vocabulary-card' key={ item.name }>
            <h4>{ item.name }</h4>
            <p>{ item.description }</p>
          </div>
        )) }
    </div>
  )
}