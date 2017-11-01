// utils
import { ReduceStore } from 'flux/utils';
// consts
import * as actionTypes from '../actions/types';
import appDispatcher from '../utils/dispatcher';


class AppStore extends ReduceStore {

    getInitialState () {
        return {
            errors: null,
            weather: {},
        };
    }

    reduce (state, action) {
        switch (action.type) {
            case actionTypes.APP_INIT:
                return {
                    ...state,
                    weather: action.data
                };

            default:
                return state;
        };
    };
}

export default new AppStore(appDispatcher);
