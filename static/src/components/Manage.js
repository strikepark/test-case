import React, { Component } from 'react'

import OrderList from './OrderList'
import OrderNew from './OrderNew'

export default class Manage extends Component {
    showLogin() {
        this.props.showLogin()
    }

    render() {
        const { active, orders, createOrder, updateOrder } = this.props;
        const isActive = active ? '' : 'hidden'

        return (
            <div className={'content ' + isActive}>
                <button onClick={::this.showLogin} type='button' className='button-xsmall pure-button content__back'>← Вернуться</button>

                <OrderNew createOrder={createOrder} />
                <OrderList orders={orders} updateOrder={updateOrder} />
            </div>
        );
    }
}
