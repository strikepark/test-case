import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import * as orderActions from '../actions/OrderActions'
import Manage from '../components/Manage'
import Login from '../components/Login'
import NotFound from '../components/NotFound'
import $ from 'jquery';

import { Router, Route, browserHistory } from 'react-router'


class App extends Component {
  render() {
    const orders = this.props.order.orderList;
    const fetching = this.props.order.fetching;
    const error = this.props.order.error;
    const { updateOrder, createOrder } = this.props.orderActions

    if (fetching) {
        $('body').addClass('fetching');
    } else {
        $('body').removeClass('fetching');
    }

    return (
        <Router history={browserHistory}>
            <Route path='/' component={Login}>
                <Route path='manage' component={Manage} createOrder={createOrder} orders={orders} updateOrder={updateOrder} />
            </Route>

            <Route path='*' component={NotFound} />
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
