import React from 'react'
import {
    MdAddCircleOutline,
    MdRemoveCircleOutline,
    MdEdit,
    MdEditOff,
    MdVisibility,
    MdVisibilityOff,
    MdRefresh
} from 'react-icons/md'
import '../App.css'
import Header from '../navbar/navbar'
import data from './data.json'
import { playerIDs, playerNames, playersArray } from '../common/players'

function enable(timestamp) {
    console.log('enable:' + timestamp)
}

function disable(timestamp) {
    console.log('disable:' + timestamp)
}

const DisabledItem = ({ deltas, readOnly, timestamp, hide }) => {
    return (
        hide
            ? null
            : <tr class='bg-red-100 border-b dark:bg-gray-800 dark:border-gray-700'>

                <td key='td-disabled' scope='row' class='px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap'>
                    {readOnly
                        ? <i>disabled</i>
                        : <button onClick={_ => enable(timestamp)}>
                            <MdAddCircleOutline />
                        </button>}
                </td>
                {deltas.map((delta, tdIdx) => (
                    <td key={'td-' + tdIdx} scope='row' class='px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap'>
                        {delta}
                    </td>
                ))}
            </tr>
    )
}

const EnabledItem = ({ deltas, readOnly, timestamp, rnd }) => {
    return (
        <tr class='bg-white border-b dark:bg-gray-800 dark:border-gray-700'>
            <td key='td-round' scope='row' class='px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap'>
                {readOnly
                    ? <p> Round {rnd}</p>
                    : <button onClick={_ => disable(timestamp)}>
                        <MdRemoveCircleOutline />
                    </button>}
            </td>
            {deltas.map((delta, tdIdx) => (
                <td key={'td-' + tdIdx} scope='row' class='px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap'>
                    {delta}
                </td>
            ))}
        </tr>
    )
}

function DouDiZhuView() {
    const [hideDisabled, setHide] = React.useState(true)
    const [readOnly, setReadOnly] = React.useState(true)

    const toggleVisible = () => {
        if (hideDisabled) {
            setHide(false)
        } else {
            setReadOnly(true)
            setHide(true)
        }
    }
    const toggleReadOnly = () => {
        if (readOnly) {
            setHide(false)
            setReadOnly(false)
        } else {
            setReadOnly(true)
        }
    }

    const refresh = () => console.log('refresh')

    return (
        <>
            <Header />
            <button class='top-0 left-0 ml-3 mt-3' onClick={refresh}>
                <MdRefresh />
            </button>
            <button class='top-0 left-0 ml-3 mt-3' onClick={toggleVisible}>
                {hideDisabled ? <MdVisibility /> : <MdVisibilityOff />}
            </button>
            <button class='top-0 left-0 ml-3 mt-3' onClick={toggleReadOnly}>
                {readOnly ? <MdEdit /> : <MdEditOff />}
            </button>
            <div className='flex flex-col items-center min-h-screen py-2 z-0'>
                <div class='overflow-x-auto shadow-md sm:rounded-lg'>
                    <table class='w-full text-sm text-left text-gray-500 dark:text-gray-400'>
                        <thead class='text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400'>
                            <tr>
                                <th />
                                {data.playerNames.map((player) => <th scope='col' class='px-6 py-3'> {player} </th>)}
                            </tr>
                        </thead>
                        <tbody>
                            {data.deltaScores.map((o, idx) => {
                                if (o.enabled) {
                                    return <EnabledItem key={'td' + idx.toString()} deltas={o.deltas} readOnly={readOnly} timestamp={o.timestamp} rnd={o.round} />
                                } else {
                                    return <DisabledItem key={'td' + idx.toString()} deltas={o.deltas} readOnly={readOnly} timestamp={o.timestamp} hide={hideDisabled} />
                                }
                            }
                            )}
                            <tr class='bg-white border-b dark:bg-gray-800 dark:border-gray-700'>
                                <td scope='row' class='px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap'>Scores</td>
                                {data.finalScores.map((score) => <td scope='row' class='px-6 py-4 font-bold text-gray-900 dark:text-white whitespace-nowrap'>{score}</td>)}
                            </tr>
                        </tbody>
                    </table>

                </div>
            </div>
        </>
    )
}

export default DouDiZhuView
