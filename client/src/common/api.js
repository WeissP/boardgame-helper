import data from '../config.json'
import axios from 'axios'

export function post(endpoint, data) {
    return axios.post('/api/' + endpoint, data)
}

export function get(endpoint, data) {
    return axios.get('/api/' + endpoint)
}

// class TestV extends React.Component {
//     constructor(props) {
//         super(props)
//         this.state = {
//             url: endpoint + 'test/123f',
//             ttt: 'qwr',
//             xxx: 'a'
//         }
//     };

//     componentDidMount() {
//         // this.setState({ ttt: 'asdff' })
//         axios.get('http://localhost:8888/api/test/123f').then(response => {
//             // this.setState({ xxx: 'success' })
//             const result = response.data

//             // console.log(result)
//             // this.setState({ xxx: response.data.point })
//             this.setState({ xxx: result.point })
//         }
//         ).catch(err => {
//             console.log(err)
//             this.setState({ xxx: 'err' })
//         }
//         )
//     }

//     render() {
//         return (
//             <div> <p> TestV: {this.state.xxx}  </p> </div>
//         )
//     }
// }



