import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import * as orderActions from '../actions/OrderActions'
import OrderList from '../components/OrderList'
import OrderNew from '../components/OrderNew'

import $ from 'jquery';

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

    if (error !== '') {
        alert(error);
    }

    return (
        <div className='content'>
            <OrderNew createOrder={createOrder} />
            <OrderList orders={orders} updateOrder={updateOrder} />
        </div>
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
