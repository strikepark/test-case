import React, { Component } from 'react'

import UserOrdersItem from './UserOrdersList'

export default class User extends Component {
    render() {
        const { active, orders } = this.props
        const isActive = active ? '' : 'hidden'

        console.log(orders)

        return (
            <div className={'content ' + isActive}>
                <UserOrdersItem orders={orders} />
            </div>
        )
    }
}
