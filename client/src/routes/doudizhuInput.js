import '../App.css'
import React from 'react'
import post from '../common/api'
import { PlayerList, INum, UserList } from '../common/input'
import { InputNumber, Form } from 'rsuite'
import Header from '../navbar/navbar'
import { playerValue, playersArray } from '../common/players'

function useForceUpdate() {
    const [value, setValue] = React.useState(0) // integer state
    return () => setValue(value => value + 1) // update the state to force render
}

// const PlayersPos = ({ playersPosChanged }) => {
//     // const [playersPos, setPlayersPos] = React.useState(['unknown', 'unknown', 'unknown', 'unknown'])
//     // const updateUser = (key, id) => {
//     //     const n = playersPos
//     //     n[id] = key
//     //     setPlayersPos(n)
//     //     // console.log(players)
//     //     playersPosChanged(playersPos)
//     // }
//     return (
//         <div class='grid grid-cols-3 gap-0'>
//             <div />
//             <div> <UserList key='user0' userChanged={key => updateUser(key, 0)} /> </div>
//             <div />
//             <div> <UserList key='user1' userChanged={key => updateUser(key, 1)} />  </div>
//             <div />
//             <div> <UserList key='user2' userChanged={key => updateUser(key, 2)} />  </div>
//             <div />
//             <div> <UserList key='user3' userChanged={key => updateUser(key, 3)} />  </div>
//             <div />
//         </div>
//     )
// }

const DouDiZhuInput = () => {
    const [players, setPlayers] = React.useState(['unknown', 'unknown', 'unknown', 'unknown'])
    // const [player0, setPlayer0] = React.useState('unknown')
    // const [player1, setPlayer0] = React.useState('unknown')
    // const [player1, setPlayer0] = React.useState('unknown')
    const [points, setPoints] = React.useState(8)
    const [winner, setWinner] = React.useState('')
    const [weight, setWeight] = React.useState({})
    const [lord, setLord] = React.useState('')
    const [stake, setStake] = React.useState(0)
    const [BonusTiles, setBonusTiles] = React.useState(3)

    function handleSubmit(event) {
        console.log(this.state)
        // event.preventDefault()
        post('save', { ...this.state, timestamp: new Date().toISOString() })
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
                    <div> <PlayerList key='user1' player={players[1]} playerOnChange={key => updatePlayer(key, 1)} /> </div>
                    <div />
                    <div> <PlayerList key='user2' player={players[2]} playerOnChange={key => updatePlayer(key, 2)} /> </div>
                    <div />
                    <div> <PlayerList key='user3' player={players[3]} playerOnChange={key => updatePlayer(key, 3)} /> </div>
                    <div />
                </div>
                <div class='mt-2 pb-4 max-w-sm'>
                    <InputNumber prefix='Stake' value={stake} onChange={(n, _) => setStake(n)} min={0} />
                    <br />
                    <InputNumber prefix='Tiles' value={BonusTiles} onChange={(n, _) => setBonusTiles(n)} min={3} />
                    <br />
                    <div class='grid grid-cols-6 items-center gap-4'>
                        <div> <p class=' text-right'>Lord: </p>  </div>
                        <div class='col-span-2'> <UserList key='lord-select' userChanged={setLord} />  </div>
                        <div> <label>Winner: </label>  </div>
                        <div class='col-span-2'> <UserList key='winner-select' userChanged={setWinner} />  </div>
                    </div>
                    <br/>
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
                    <InputNumber prefix='Points' value={points} onChange={(n, _) => setPoints(n)} min={8} />
                </div>
            </div>
        </>
    )
}

// export default class DouDiZhuInput extends React.Component {
//     constructor(props) {
//         super(props)
//         this.state = {
//             players: ['unknown', 'unknown', 'unknown', 'unknown'],
//             winner: '',
//             weight: {},
//             lord: '',
//             points: 0,
//             stake: 0,
//             bonusTiles: 0
//         }
//     };

//     handleSubmit(event) {
//         console.log(this.state)
//         // event.preventDefault()
//         post('save', { ...this.state, timestamp: new Date().toISOString() })
//     }

//     render() {
//         return (
//             <>
//                 <Header />
//                 <div className='flex flex-col items-center justify-center min-h-screen py-2'>
//                     <PlayersPos playersPosChanged={data => this.setState({ players: data })} />
//                     <div class='ml-1 pb-4 max-w-sm'>
//                         <p class='text-left pb-4'> Lord:  <UserList userChanged={user => this.setState({ lord: user })} /> </p>
//                         <span>Stake</span>
//                         <InputNumber
//                             min={0}
//                             defaultValue={this.state.stake}
//                             onChange={data => this.setState({ stake: data })}
//                         />
//                         <InputNumber
//                             initLabel='Lord Bonus Tiles'
//                             initValue={this.state.bonusTiles}
//                             inputChanged={data => this.setState({ bonus_tiles: data })}
//                         />
//                         <p class='text-lg font-bold text-left'> weight: </p>
//                         {this.state.players.map((player) => (
//                             <InputNumber
//                                 initLabel={playerValue(player)}
//                                 initValue={this.state.weight[player]}
//                                 inputChanged={data => {
//                                     const newWeight = this.state.weight
//                                     newWeight[player] = data
//                                     this.setState({ weight: newWeight })
//                                 }}
//                             />
//                         ))}

//                         <p class='text-left pb-4'> Winner:  <UserList userChanged={user => this.setState({ winner: user })} /> </p>
//                         <InputNumber
//                             initLabel='Points'
//                             initValue={this.state.points}
//                             inputChanged={data => this.setState({ points: data })}
//                         />
//                         <button class='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded' onClick={e => this.handleSubmit(e)}>submit</button>
//                     </div>
//                 </div>

//                 <Form layout='inline'>
//                     <Form.Group controlId='stake'>
//                         <Form.ControlLabel>Stake</Form.ControlLabel>
//                         <Form.Control name='stake' style={{ width: 160 }} />
//                     </Form.Group>
//                 </Form>
//             </>
//         )
//     }
// }

export default DouDiZhuInput
