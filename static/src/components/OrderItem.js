import React, { PropTypes, Component } from 'react'
import { fmtCode } from '../helpers'

import OrderEdit from './OrderEdit'

export default class OrderItem extends Component {
    constructor(props) {
        super(props);

        this.state = {
            active: false
        };
    }

    onOrderClick() {
        this.setState({
            active: !this.state.active
        });
    }

    render() {
        const order = this.props.order;
        const updateOrder= this.props.updateOrder;
        let code = fmtCode(order.code);

        return (
            <div className='list__item'>
                <div className='list__name' onClick={::this.onOrderClick}>
                    Заказ № {code}
                </div>

                <OrderEdit order={order} active={this.state.active} updateOrder={updateOrder} />
            </div>
        )
    }
}

OrderItem.propTypes = {
    order: PropTypes.object.isRequired
}
