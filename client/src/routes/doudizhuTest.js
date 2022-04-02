import '../App.css'
import React from 'react'
import post from '../common/api'
import { InputNumber, InputText, UserList } from '../common/input'

const PlayersPos = ({ playersPosChanged }) => {
    const [players, setPlayers] = React.useState(['unknown', 'unknown', 'unknown', 'unknown'])
    const updateUser = (data, id) => {
        const n = players
        n[id] = data
        setPlayers(n)
        playersPosChanged(players)
    }
    return (
        <div class='flex flex-wrap overflow-hidden'>
            <div class='w-1/3 overflow-hidden' />
            <div class='w-1/3 overflow-hidden'> <UserList userChanged={data => updateUser(data, 0)} />
            </div>
            <div class='w-1/3 overflow-hidden' />
            <div class='w-1/3 overflow-hidden'> <UserList userChanged={data => updateUser(data, 1)} />  </div>
            <div class='w-1/3 overflow-hidden' />
            <div class='w-1/3 overflow-hidden'> <UserList userChanged={data => updateUser(data, 2)} />  </div>
            <div class='w-1/3 overflow-hidden' />
            <div class='w-1/3 overflow-hidden'> <UserList userChanged={data => updateUser(data, 3)} />  </div>
            <div class='w-1/3 overflow-hidden' />
        </div>
    )
}

// const input = ({input})

export default class DouDiZhuTest extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            players: ['unknown', 'unknown', 'unknown', 'unknown'],
            winner: '',
            weight: {},
            lord: 0,
            points: 0,
            stake: 0,
            bonus_tiles: 0,
            final_points: {}
        }
    };

    handleSubmit(event) {
        console.log(this.state)
        // event.preventDefault()
        post('save', { ...this.state, timestamp: new Date().toISOString() })
    }

    render() {
        return (
            <>
                <PlayersPos playersPosChanged={data => this.setState({ players: data })} />
                <div class='ml-1 pb-4 max-w-sm'>
                    <p class='text-left pb-4'> Lord:  <UserList userChanged={user => this.setState({ Lord: user })} /> </p>
                    <InputNumber
                        initLabel='Stake'
                        initValue={this.state.stake}
                        inputChanged={data => this.setState({ stake: data })}
                    />
                    <InputNumber
                        initLabel='Lord Bonus Tiles'
                        initValue={this.state.bonus_tiles}
                        inputChanged={data => this.setState({ bonus_tiles: data })}
                    />
                    <p class='text-lg font-bold text-left'> weight: </p>
                    {this.state.players.map((player) => (
                        <InputNumber
                            initLabel={player}
                            initValue={this.state.weight[player]}
                            inputChanged={data => {
                                const newWeight = this.state.weight
                                newWeight[player] = data
                                this.setState({ weight: newWeight })
                            }}
                        />
                    ))}

                    <p class='text-left pb-4'> Winner:  <UserList userChanged={user => this.setState({ winner: user })} /> </p>
                    <InputNumber
                        initLabel='Points'
                        initValue={this.state.points}
                        inputChanged={data => this.setState({ points: data })}
                    />
                    <p class='text-lg font-bold text-left'> Final Points: </p>
                    {this.state.players.map((player) => (
                        <InputNumber
                            initLabel={player}
                            initValue={this.state.final_points[player]}
                            inputChanged={data => {
                                const newFP = this.state.final_points
                                newFP[player] = data
                                this.setState({ final_points: newFP })
                            }}
                        />
                    ))}
                    <button class='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded' onClick={e => this.handleSubmit(e)}>submit</button>
                </div>
            </>
        )
    }
}

// export default class DouDiZhuView extends React.Component {
//     constructor(props) {
//         super(props)
//         this.state = {
//             players: ['unknown', 'unknown', 'unknown', 'unknown'],
//             winner: '',
//             weight: {},
//             lord: 0,
//             points: 0,
//             stake: 0,
//             bonus_tiles: 0,
//             final_points: {}
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
//                 <PlayersPos playersPosChanged={data => this.setState({ players: data })} />
//                 <div class='ml-1 pb-4 max-w-sm'>
//                     <p class='text-left pb-4'> Lord:  <UserList userChanged={user => this.setState({ Lord: user })} /> </p>
//                     <InputNumber
//                         initLabel='Stake'
//                         initValue={this.state.stake}
//                         inputChanged={data => this.setState({ stake: data })}
//                     />
//                     <InputNumber
//                         initLabel='Lord Bonus Tiles'
//                         initValue={this.state.bonus_tiles}
//                         inputChanged={data => this.setState({ bonus_tiles: data })}
//                     />
//                     <p class='text-lg font-bold text-left'> weight: </p>
//                     {this.state.players.map((player) => (
//                         <InputNumber
//                             initLabel={player}
//                             initValue={this.state.weight[player]}
//                             inputChanged={data => {
//                                 const newWeight = this.state.weight
//                                 newWeight[player] = data
//                                 this.setState({ weight: newWeight })
//                             }}
//                         />
//                     ))}

//                     <p class='text-left pb-4'> Winner:  <UserList userChanged={user => this.setState({ winner: user })} /> </p>
//                     <InputNumber
//                         initLabel='Points'
//                         initValue={this.state.points}
//                         inputChanged={data => this.setState({ points: data })}
//                     />
//                     <button class='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded' onClick={e => this.handleSubmit(e)}>submit</button>
//                 </div>
//             </>
//         )
//     }
// }
