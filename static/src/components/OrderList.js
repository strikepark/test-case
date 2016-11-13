import React, { PropTypes, Component } from 'react'

import OrderItem from './OrderItem'

export default class OrderList extends Component {
    render() {
        let { orders, updateOrder } = this.props;

        // Sort by id
        orders.sort((a, b) => {
            if (a.id < b.id) {
                return 1;
            } else if (a.id > b.id) {
                return -1;
            } else {
                return 0;
            }
        });

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
