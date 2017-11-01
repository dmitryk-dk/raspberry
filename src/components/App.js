import React, { Component } from 'react';
import { Container } from 'flux/utils';
import Weather from './Weather';


// store
import appStore from '../stores/store';
// actions
import * as actions from '../actions/actions'

class App extends Component {

    static getStores() {
        return [ appStore ];
    }

    static calculateState(prevState, props) {
        return {
            ...appStore.getState()
        }
    }

    render() {
        return (
            <Weather
                {...this.state.weather}
                {...actions}
            />
        );
    }
}

export default new Container.create(App);
