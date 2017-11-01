import appDispatcher from '../utils/dispatcher';
import * as http from '../utils/http';
import * as actionTypes from "./types";
import * as consts from '../utils/consts';

export const getData = () => {
    const url = consts.getHostFn().replace('{API}', consts.getTemperature);
    http.get(url)
        .then(data => appDispatcher.dispatch({
            type: actionTypes.APP_INIT,
            data
        }))
        .catch(error => appDispatcher.dispatch({
            type: actionTypes.APP_INIT_ERROR,
            error
        }))
};

