import appDispatcher from '../utils/dispatcher';
import * as http from '../utils/http';
import * as actionTypes from "./types";
import * as consts from '../utils/consts';

// export const submit = (phone) => {
//     const url =  consts.getHostFn().replace('{API}', consts.postPhone);
//     appDispatcher.dispatch({
//         type: actionTypes.PHONE_SAVE_REQUEST,
//     });
//     http.post(url, phone)
//         .then(data => appDispatcher.dispatch({
//             type: actionTypes.PHONE_SAVE_SUCCESS,
//             phone
//         }))
//         .catch(error => appDispatcher.dispatch({
//             type: actionTypes.PHONE_SAVE_FAILED,
//             error
//         }))
// };

export const getData = (url) => {
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

export const deletePhone = (phone) => {
    const url =  consts.getHostFn().replace('{API}', consts.delPhone);
    appDispatcher.dispatch({
        type: actionTypes.PHONE_DELETE_REQUEST,
    });
    http.del(url, phone)
        .then(data => appDispatcher.dispatch({
            type: actionTypes.PHONE_DELETE_SUCCESS,
            phone
        }))
        .catch(error => appDispatcher.dispatch({
            type: actionTypes.PHONE_DELETE_FAILED,
            error
        }))
}
