import React, { PropTypes, Component } from 'react'

import OrderItem from './OrderItem'
import OrderNew from './OrderNew'

export default class OrderList extends Component {
    render() {
        const updateOrder= this.props.updateOrder;
        let orders = this.props.orders;

        // Sort by id
        orders.sort((a, b) => {
            if (a.value > b.value) {
                return 1;
            } else if (a.value < b.value) {
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
