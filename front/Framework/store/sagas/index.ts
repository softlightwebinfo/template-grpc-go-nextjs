// @ts-ignore
import {all} from "redux-saga/effects";

// @ts-ignore
import translate from "./translate";

// @ts-ignore
export default function* rootSaga() {
    yield all([translate()]);
}
