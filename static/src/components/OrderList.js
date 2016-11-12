import React, { PropTypes, Component } from 'react'

import OrderItem from './OrderItem'

export default class OrderList extends Component {
    render() {
        const orders = this.props.orders;
        const updateOrder= this.props.updateOrder;
        const listItems = orders.map((order) =>
            <OrderItem order={order} key={order.id} updateOrder={updateOrder} />
        );

        return (
            <div className='list'>
                {listItems}
            </div>
        )
    }
}

OrderList.propTypes = {
    orders: PropTypes.array.isRequired
}
