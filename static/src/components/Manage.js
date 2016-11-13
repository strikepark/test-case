import React, { Component } from 'react'

import OrderList from './OrderList'
import OrderNew from './OrderNew'

export default class Manage extends Component {
    render() {
        const { orders, createOrder, updateOrder } = this.props.route;

        return (
            <div className='content'>
                <OrderNew createOrder={createOrder} />
                <OrderList orders={orders} updateOrder={updateOrder} />
            </div>
        );
    }
}
