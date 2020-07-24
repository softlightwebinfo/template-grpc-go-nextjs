import React, {Component, Fragment} from "react";
import {connect} from "react-redux";
import Head from "next/head";
import {translationRequest} from "../Framework/store/actions/translate";
//import {END} from "redux-saga";

class Index extends Component {
    static async getInitialProps(ctx) {
        const {isServer, store} = ctx;
        store.dispatch(translationRequest("es"));
        // if (ctx.req) {
        //     ctx.store.dispatch(END);
        //     await ctx.store.sagaTask.toPromise();
        // }
        return {isServer};
    }

    componentDidMount() {
        // @ts-ignore
        const {dispatch, isServer, translate} = this.props;

        if (isServer && !translate) {
            //dispatch(ActionCreator.translationRequest("es"));
        }
    }

    render() {
        return (
            <Fragment>
                <Head>
                    <title>My website</title>
                    <meta name="title" content="my website"/>
                </Head>

            </Fragment>
        );
    }
}

export default connect()(Index);
