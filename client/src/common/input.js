import '../App.css'
import React from 'react'
import { playerValue, playersArray } from './players'
import { InputGroup, InputNumber } from 'rsuite'

const PlayerList = ({ player, playerOnChange }) => {
    return (
        <select
            class='form-select block w-fit px-1 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none'
            value={player}
            onChange={e => playerOnChange(e.target.value)}
        >
            <option value='unknown'>name</option>
            {playersArray().map(([key, val]) => <option value={key}>{val}</option>)}
        </select>
    )
}

const UserList = ({ userChanged }) => {
    return (
        <select
            class='form-select block w-fit px-1 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none'
            onChange={e => userChanged(e.target.value)}
        >
            <option value='unknown'>name</option>
            {playersArray().map(([key, val]) => <option value={key}>{val}</option>)}
        </select>
    )
}

const InputText = ({ initLabel, initValue, inputChanged }) => {
    return (
        <div class='md:flex md:items-center mb-6'>
            <div class='md:w-1/3'>
                <label class='block text-gray-700 text-sm font-bold mb-2 text-left'>
                    {initLabel}:
                </label>
            </div>
            <div class='md:w-2/3'>
                <input class='bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500' type='text' value={initValue} onChange={e => inputChanged(e.target.value)} />
            </div>
        </div>
    )
}

const INum = ({ dftValue, label }) => {
    const [value, setValue] = React.useState(dftValue)
    const handleMinus = () => {
        setValue(parseInt(value, 10) - 1)
    }
    const handlePlus = () => {
        setValue(parseInt(value, 10) + 1)
    }
    return (
        <div style={{ width: 160 }}>
            <span>stake: </span>
            <InputGroup>
                <InputGroup.Button onClick={handleMinus}>-</InputGroup.Button>
                <InputNumber className='custom-input-number' value={value} onChange={setValue} />
                <InputGroup.Button onClick={handlePlus}>+</InputGroup.Button>
            </InputGroup>
        </div>
    )
}

export { INum, InputText, UserList ,PlayerList}
