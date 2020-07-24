import {applyMiddleware, createStore} from 'redux'
import createSagaMiddleware from 'redux-saga'
import {createWrapper} from 'next-redux-wrapper'

import rootReducer from './reducers'
import sagas from "./sagas";

const bindMiddleware = (middleware) => {
    if (process.env.NODE_ENV !== 'production') {
        const {composeWithDevTools} = require('redux-devtools-extension')
        return composeWithDevTools(applyMiddleware(...middleware))
    }
    return applyMiddleware(...middleware)
};


export const makeStore = (_, initialState = {}) => {
    const sagaMiddleware = createSagaMiddleware();

    const store = createStore(
        rootReducer,
        initialState,
        bindMiddleware([sagaMiddleware])
    );

    // @ts-ignore
    store.runSagaTask = () => {
        // @ts-ignore
        store.sagaTask = sagaMiddleware.run(sagas);
    };

    // @ts-ignore
    store.runSagaTask();
    return store
};


export const wrapper = createWrapper(makeStore, {debug: false});
