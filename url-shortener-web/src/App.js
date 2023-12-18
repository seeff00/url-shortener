// import logo from './logo.svg';
import './App.css';
import React, { Component }  from 'react';
import ShortURLForm from "./components/ShortURLForm";

class App extends Component {
    render () {
        return (
            <div className="App">
                <ShortURLForm />
            </div>
        );
    }
}

export default App;
