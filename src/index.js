import React from 'react';
import {render} from 'react-dom';
import App from './components/App';
import css from './styles/style.scss';
import * as actions from './actions/actions'
import * as consts from './utils/consts';

function init() {
    const url = consts.getHostFn().replace('{API}', consts.getTemperature);
    actions.getData(url);
    render(<App />,document.getElementById('app'));
}

init();

