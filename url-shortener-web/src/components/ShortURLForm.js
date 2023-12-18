import React, {Component} from 'react';
import configData from "./../config.json";

class ShortURLForm extends Component {
    constructor(props) {
        super(props);
        this.state = {
            urlValue: '',
            apiResponse: {data: []},
            apiBaseUrl: configData.URL_SHORTENER_API_HOST + ':' + configData.URL_SHORTENER_API_PORT + '/',
            apiError: ''
        };
    }

    handleChange = (event) => {
        this.setState({
            urlValue: event.target.value,
            apiResponse: {data: []},
            apiError: ''
        });
    }

    handleSubmit = async (event) => {
        event.preventDefault();
        event.stopPropagation();

        fetch(this.state.apiBaseUrl + "short", {
            method: 'POST',
            body: JSON.stringify({url: this.state.urlValue})
        })
            .then(response => {
                if (response.ok) {
                    return response.json();
                } else {
                    throw new Error('Invalid URL!');
                }
            })
            .then(data => this.setState({apiResponse: data}))
            .catch(error => this.setState({apiError: error.message}));
    }

    render() {
        return (
            <form className="shortUrlForm" onSubmit={this.handleSubmit}>
                <div>{this.state.apiError !== "" && this.state.apiError}</div>
                <div><span>URL:</span>
                    <input type="text" value={this.state.urlValue} onChange={this.handleChange}/>
                    <input className="submit" type="submit" value="Generate"/>
                </div>
                <div><span>Short URL:</span>
                    {this.state.apiResponse == null || this.state.apiResponse.code === undefined
                        ? ""
                        : <a target="_blank" rel="noopener noreferrer"
                             href={this.state.apiBaseUrl + this.state.apiResponse.code}>
                            {this.state.apiBaseUrl + this.state.apiResponse.code}</a>}
                </div>
            </form>
        );
    }
}

export default ShortURLForm