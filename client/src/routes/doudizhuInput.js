import '../App.css'
import React, { useEffect } from 'react'
import { get, post } from '../common/api'
import { PlayerList, INum, UserList } from '../common/input'
import { InputNumber, Button, Tooltip, Whisper, PreventOverflowContainer, Message, toaster } from 'rsuite'
import Header from '../navbar/navbar'
import { playerValue, playersArray } from '../common/players'
import { useParams, useNavigate } from 'react-router-dom'
import { objectMap } from '../common/misc'

const defaultOption = {
    players: Array(4).fill('unknown'),
    rawPoints: 8,
    winner: '',
    weight: {},
    lord: '',
    stake: 1,
    bonusTiles: 3
}

// function getOption(timestamp) {
//     if (timestamp == '') {
//         return defaultOption
//     } else {

//     }
// }

async function curPlayers() {
    get('/doudizhu/curPlayers').then((response) => {
        return response.data.players
    }).catch(_ => { })
};

const DouDiZhuInput = ({ timestamp }) => {
    const navigate = useNavigate()
    const message = (
        <Message showIcon type='success'>
            success!
        </Message>
    )

    const [players, setPlayers] = React.useState(defaultOption.players)
    const [rawPoints, setRawPoints] = React.useState(defaultOption.rawPoints)
    const [winner, setWinner] = React.useState(defaultOption.winner)
    const [weight, setWeight] = React.useState(Array(4).fill(0))
    const [lord, setLord] = React.useState(defaultOption.lord)
    const [stake, setStake] = React.useState(defaultOption.stake)
    const [bonusTiles, setBonusTiles] = React.useState(defaultOption.bonusTiles)

    const updatePlayer = (player, index) => {
        setPlayers((prev) => prev.map((el, i) => (i !== index ? el : player)))
    }

    const updateWeight = (weight, index) => {
        setWeight((prev) => prev.map((el, i) => (i !== index ? el : weight)))
    }

    function setOptionWithoutPlayers(opt) {
        setRawPoints(opt.rawPoints)
        setWinner(opt.winner)
        setLord(opt.lord)
        setStake(opt.stake)
        setBonusTiles(opt.bonusTiles)

        opt.players.forEach((x, idx) => {
            const w = opt.weight[x]
            if (typeof w === 'number') {
                updateWeight(w, idx)
            }
        })
    }

    useEffect(() => {
        async function getLastPlayers() {
            const resp = await get('/doudizhu/curPlayers')
            const p = resp.data.players
            if (Array.isArray(p)) {
                setPlayers(p)
            }
        }

        async function getOption() {
            const resp = await get('/doudizhu/edit', { timestamp: timestamp })
            const opt = resp.data
            if (opt != null) {
                setPlayers(opt.players)
                setOptionWithoutPlayers(opt)
            }
        }

        if (timestamp != '') {
            getOption()
        } else {
            getLastPlayers()
        }
    }, []
    )

    function handleSubmit(event) {
        const res = {
            players: players,
            rawPoints: parseInt(rawPoints, 10),
            winner: winner,
            weight: Object.fromEntries(weight.map((v, idx) => [players[idx], parseInt(v, 10)])),
            lord: lord,
            stake: parseInt(stake, 10),
            bonusTiles: parseInt(bonusTiles, 10),
            timestamp: timestamp == '' ? new Date().toISOString() : timestamp
        }
        // console.log(res)
        post('doudizhu/new', JSON.stringify(res)).then((response) => {
            toaster.push(message, 'topCenter')
            if (timestamp == '') {
                setOptionWithoutPlayers(defaultOption)
                setWeight(Array(4).fill(0))
            } else {
                navigate('/doudizhu-view')
            }
        }).catch(function(error) {
            console.log(error)
            if (error.response) {
                window.alert(error.response.data)
            }
        })
    }

    return (
        <>
            <Header />
            <div className='mt-5 flex flex-col items-center min-h-screen py-2'>
                <div class='grid grid-cols-3 gap-0'>
                    <div />
                    <div> <PlayerList key='user0' player={players[0]} playerOnChange={key => updatePlayer(key, 0)} /> </div>
                    <div />
                    <div> <PlayerList key='user3' player={players[3]} playerOnChange={key => updatePlayer(key, 3)} /> </div>
                    <div />
                    <div> <PlayerList key='user1' player={players[1]} playerOnChange={key => updatePlayer(key, 1)} /> </div>
                    <div />
                    <div> <PlayerList key='user2' player={players[2]} playerOnChange={key => updatePlayer(key, 2)} /> </div>
                    <div />
                </div>
                <div class='mt-2 pb-4 max-w-sm'>
                    <InputNumber prefix='Stake' type='number' value={stake} onChange={(n, _) => setStake(n)} min={1} />
                    <br />
                    <InputNumber prefix='Tiles' value={bonusTiles} onChange={(n, _) => setBonusTiles(n)} min={3} />
                    <br />
                    <div class='grid grid-cols-6 items-center gap-4'>
                        <div> <p class=' text-right'>Lord: </p>  </div>
                        <div class='col-span-2'> <UserList key='lord-select' value={lord} setValue={setLord} />  </div>
                        <div> <label>Winner: </label>  </div>
                        <div class='col-span-2'> <UserList key='winner-select' value={winner} setValue={setWinner} />  </div>
                    </div>
                    <br />
                    <p> Weight: </p>
                    {players.map((player, idx) => (
                        <InputNumber
                            key={'playerWeight' + idx.toString()}
                            prefix={playerValue(player)}
                            value={weight[idx]}
                            onChange={(n, _) => updateWeight(n, idx)}
                        />
                    ))}

                    <br />
                    <InputNumber prefix='RawPoints' value={rawPoints} onChange={(n, _) => setRawPoints(n)} min={8} />
                    <br />
                </div>
                <Button appearance='primary' onClick={handleSubmit}>Submit</Button>
            </div>
        </>
    )
}

const DouDiZhuInputDft = () => <DouDiZhuInput timestamp='' />

const DouDiZhuInputEdit = () => {
    const { timestamp } = useParams()
    return <DouDiZhuInput timestamp={timestamp} />
}

export { DouDiZhuInputEdit, DouDiZhuInputDft }
