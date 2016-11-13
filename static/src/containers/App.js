import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import * as orderActions from '../actions/OrderActions'
import Manage from '../components/Manage'
import Login from '../components/Login'
import $ from 'jquery';

import { Router, Route, browserHistory } from 'react-router'


class App extends Component {
    constructor(props) {
        super(props)

        const orders = this.props.order.orderList;
        const { updateOrder, createOrder } = this.props.orderActions;

        this.state = {
            routes: (
                <Router>
                    <Route path='/' component={Login} />
                    <Route path='/manage' component={Manage} createOrder={createOrder} orders={orders} updateOrder={updateOrder} />
                </Router>
            )
        };
    }
    render() {
        const fetching = this.props.order.fetching;

        if (fetching) {
            $('body').addClass('fetching');
        } else {
            $('body').removeClass('fetching');
        }

        return (
            <Router history={browserHistory}>
                {this.state.routes}
            </Router>
        );
    }
}

function mapStateToProps(state) {
    return {
        order: state.order
    }
}

function mapDispatchToProps(dispatch) {
    return {
        orderActions: bindActionCreators(orderActions, dispatch)
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(App)
