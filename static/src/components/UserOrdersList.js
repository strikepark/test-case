import React, { PropTypes, Component } from 'react'

import UserOrdersItem from './UserOrdersItem'

export default class UserOrdersList extends Component {
    render() {
        let { orders } = this.props

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
            <UserOrdersItem order={order} key={order.id} />
        )

        return (
            <div className='list list_user'>
                {listItems}
            </div>
        )
    }
}

UserOrdersList.propTypes = {
    orders: PropTypes.array.isRequired
}
