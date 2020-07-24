// @ts-ignore
import {takeLatest, put} from "redux-saga/effects";

// @ts-ignore
import {translationType} from "../actionCreators";
import {translationFailure, translationSuccess} from "../actions/translate";

// @ts-ignore
export function* changeLanguage(action) {
    const {language} = action;
    if (language) {
        try {
            const translate = require(`../../../config/translates/${language}.json`);
            console.log(`=== Translate lang from folder: `, translate.language);
            yield put(translationSuccess(translate));
        } catch (err) {
            yield put(
                translationFailure("Error to change language: ", err)
            );
        }
    } else {
        yield put(
            translationFailure("There's no language selected.")
        );
    }
}

// @ts-ignore
export default function* translateWatcher() {
    yield takeLatest(translationType.TRANSLATION_REQUEST, changeLanguage);
}
