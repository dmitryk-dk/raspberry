// utils
import { ReduceStore } from 'flux/utils';
// consts
import * as actionTypes from '../actions/types';
import appDispatcher from '../utils/dispatcher';


class AppStore extends ReduceStore {

    getInitialState () {
        return {
            errors: null,
            wheather: {},
        };
    }

    reduce (state, action) {
        switch (action.type) {
            case actionTypes.APP_INIT:
                return {
                    ...state,
                    wheather: action.data
                };

            // case actionTypes.PHONE_SAVE_SUCCESS:
            //     return {
            //         ...state,
            //         phones: [
            //             ...state.phones,
            //             action.phone
            //         ]
            //     };
            //
            // case actionTypes.PHONE_DELETE_SUCCESS:
            //     return {
            //         ...state,
            //         phones: [
            //             ...state.phones.filter(phone => phone.number !== action.phone.number)
            //         ]
            //     };

            default:
                return state;
        };
    };
}

export default new AppStore(appDispatcher);
