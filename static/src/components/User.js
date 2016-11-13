import React, { Component } from 'react'

import { UserOrderList } from './UserOrderList'

export default class User extends Component {
    render() {
        const { active, orders } = this.props
        const isActive = active ? '' : 'hidden'

        return (
            <div className={'content ' + isActive}>
                <UserOrderList orders={orders} />
            </div>
        );
    }
}
