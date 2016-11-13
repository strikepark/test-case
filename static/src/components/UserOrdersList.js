import React, { PropTypes, Component } from 'react'

// import OrderItem from './OrderItem'

export default class UserOrderList extends Component {
    render() {
        let { orders } = this.props

        console.log(orders)

        const listItems = orders.map((order) =>
            // <OrderItem order={order} key={order.id} />
        );

        return (
            <div className='list'>
                fdsffdsf
                // {listItems}
            </div>
        )
    }
}

OrderList.propTypes = {
    orders: PropTypes.array.isRequired
}
