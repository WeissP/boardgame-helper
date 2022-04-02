import React, { Component } from 'react'

import './App.css'

import 'rsuite/dist/rsuite.min.css'
import { Button } from 'rsuite'
import Header from "./navbar/navbar"

class App extends Component {
    render() {
        return (
            <div className='App'>
                <Header />
                <Button appearance='primary'> Hello world </Button>
            </div>
        )
    }
}

export default App
