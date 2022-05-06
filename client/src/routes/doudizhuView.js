import React, { useEffect } from 'react'
import { ButtonToolbar, ButtonGroup, Message, toaster } from 'rsuite'
import {
    MdAddCircleOutline,
    MdRemoveCircleOutline,
    MdModeEdit,
    MdEdit,
    MdEditOff,
    MdVisibility,
    MdVisibilityOff,
    MdRefresh
} from 'react-icons/md'
import { Edit } from '@rsuite/icons'
import '../App.css'
import Header from '../navbar/navbar'
import data from './data.json'
import { get, post } from '../common/api'
import { useNavigate } from 'react-router-dom'

function enable(timestamp) {
    get('doudizhu/enable', { timestamp: timestamp })
}

function disable(timestamp) {
    get('doudizhu/disable', { timestamp: timestamp })
}

const DisabledItem = ({ deltas, readOnly, timestamp, hide }) => {
    const navigate = useNavigate()
    const routeChange = (path) => {
        navigate(path)
    }

    return (
        hide
            ? null
            : <tr class='bg-red-100 border-b dark:bg-gray-800 dark:border-gray-700'>

                <td key='td-disabled' scope='row' class='px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap'>
                    {readOnly
                        ? <i>disabled</i>
                        : <ButtonGroup>
                            <button class='pr-2' onClick={_ => enable(timestamp)}>
                                <MdAddCircleOutline />
                            </button>
                            <button onClick={_ => navigate('doudizhu-edit/' + encodeURIComponent(timestamp))}>
                                <MdModeEdit />
                            </button>
                        </ButtonGroup>}
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
    const navigate = useNavigate()
    const routeChange = (path) => {
        navigate(path)
    }

    return (
        <tr class='bg-white border-b dark:bg-gray-800 dark:border-gray-700'>
            <td key='td-round' scope='row' class='px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap'>
                {readOnly
                    ? <p> Round {rnd}</p>
                    : <ButtonGroup>
                        <button class='pr-2' onClick={_ => disable(timestamp)}>
                            <MdRemoveCircleOutline />
                        </button>
                        <button onClick={_ => navigate('/doudizhu-edit/' + encodeURIComponent(timestamp))}>
                            <MdModeEdit />
                        </button>
                    </ButtonGroup>}
            </td>
            {deltas.map((delta, tdIdx) => (
                <td key={'td-' + tdIdx} scope='row' class='px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap'>
                    {delta}
                </td>
            ))}
        </tr>
    )
}

async function updateView() {
    await get('doudizhu/view/update').catch(e => console.log('can not update view'))
}

const DouDiZhuView = () => {
    const [data, setData] = React.useState(null)
    const [hideDisabled, setHide] = React.useState(true)
    const [readOnly, setReadOnly] = React.useState(true)

    const message = (
        <Message key='noDataToday' showIcon type='info' header='Informational'>
            No Data Today
        </Message>
    )

    const isEmpty = v => {
        return v == null || !Array.isArray(v.deltaPoints) || !(v.deltaPoints.length > 0)
    }

    async function getData() {
        const resp = await get('doudizhu/view/now')
        if (!isEmpty(resp.data)) {
            setData(resp.data)
        }
    }

    useEffect(() => {
        updateView()
        getData()
        const interval = setInterval(() => {
            getData()
        }, 500)
        return () => clearInterval(interval)
    }, [])

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

    const refresh = async () => {
        await updateView()
        getData()
        if (isEmpty(data)) {
            toaster.push(message, 'bottomCenter')
        }
    }

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
            {!isEmpty(data) && <div>
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
                                {data.deltaPoints.map((o, idx) => {
                                    if (o.enabled) {
                                        return <EnabledItem key={'td' + idx.toString()} deltas={o.deltas} readOnly={readOnly} timestamp={o.timestamp} rnd={o.round} />
                                    } else {
                                        return <DisabledItem key={'td' + idx.toString()} deltas={o.deltas} readOnly={readOnly} timestamp={o.timestamp} hide={hideDisabled} />
                                    }
                                }
                                )}
                                <tr class='bg-white border-b dark:bg-gray-800 dark:border-gray-700'>
                                    <td scope='row' class='px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap'>Points</td>
                                    {data.finalPoints.map((score) => <td scope='row' class='px-6 py-4 font-bold text-gray-900 dark:text-white whitespace-nowrap'>{score}</td>)}
                                </tr>
                            </tbody>
                        </table>

                    </div>
                </div>
            </div>}
        </>
    )
}

export default DouDiZhuView
