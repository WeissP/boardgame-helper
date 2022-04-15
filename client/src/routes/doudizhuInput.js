import '../App.css'
import React from 'react'
import { post } from '../common/api'
import { PlayerList, INum, UserList } from '../common/input'
import { InputNumber, Button } from 'rsuite'
import Header from '../navbar/navbar'
import { playerValue, playersArray } from '../common/players'

const defaultOption = {
    players: ['unknown', 'unknown', 'unknown', 'unknown'],
    rawPoints: 8,
    winner: '',
    weight: {},
    lord: '',
    stake: 1,
    bonusTiles: 3,
    timestamp: ''
}

const DouDiZhuInput = ({ option }) => {
    const [players, setPlayers] = React.useState(option.players)
    const [rawPoints, setRawPoints] = React.useState(option.rawPoints)
    const [winner, setWinner] = React.useState(option.winner)
    const [weight, setWeight] = React.useState(option.weight)
    const [lord, setLord] = React.useState(option.lord)
    const [stake, setStake] = React.useState(option.stake)
    const [bonusTiles, setBonusTiles] = React.useState(option.bonusTiles)

    function handleSubmit(event) {
        const res = {
            players: players,
            rawPoints: rawPoints,
            winner: winner,
            weight: weight,
            lord: lord,
            stake: stake,
            bonusTiles: bonusTiles,
            timestamp: option.timestamp == '' ? new Date().toISOString() : option.timestamp
        }
        console.log(res)
        post('save', res)
    }

    const updatePlayer = (player, index) => {
        setPlayers((prev) => prev.map((el, i) => (i !== index ? el : player)))
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
                    <InputNumber prefix='Stake' value={stake} onChange={(n, _) => setStake(n)} min={1} />
                    <br />
                    <InputNumber prefix='Tiles' value={bonusTiles} onChange={(n, _) => setBonusTiles(n)} min={3} />
                    <br />
                    <div class='grid grid-cols-6 items-center gap-4'>
                        <div> <p class=' text-right'>Lord: </p>  </div>
                        <div class='col-span-2'> <UserList key='lord-select' userChanged={setLord} />  </div>
                        <div> <label>Winner: </label>  </div>
                        <div class='col-span-2'> <UserList key='winner-select' userChanged={setWinner} />  </div>
                    </div>
                    <br />
                    <p> Weight: </p>
                    {players.map((player, idx) => (
                        <InputNumber
                            key={'player' + idx.toString()}
                            prefix={playerValue(player)}
                            value={weight[player]}
                            onChange={(n, _) => {
                                const newWeight = weight
                                newWeight[player] = n
                                setWeight(newWeight)
                            }}
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

const DouDiZhuInputDft = () => <DouDiZhuInput option={defaultOption} />

export { DouDiZhuInput, DouDiZhuInputDft, defaultOption }
