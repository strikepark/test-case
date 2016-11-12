import { findDOMNode } from 'react-dom'
import React, { PropTypes, Component } from 'react'
import { fmtCode, genStatusSelect } from '../helpers'

import serialize from 'form-serialize'
import $ from 'jquery'

// import { connect } from 'react-redux';
// import { updateOrder } from '../actions/OrderActions'

export default class OrderEdit extends Component {
    constructor(props) {
        super(props);

        this.state = {
            order: this.props.order,
            active: this.props.active,
            disabled: true
        };
    }

    componentWillReceiveProps(nextProps) {
        let nextPropsStr = JSON.stringify(nextProps.orders) + JSON.stringify(nextProps.active);
        let curPropsStr = JSON.stringify(this.state.orders) + JSON.stringify(this.state.active);

        if (nextPropsStr !== curPropsStr) {
            this.setState({
                order: nextProps.order,
                active: nextProps.active
            });
        }
    }

    onEditClick() {
        this.setState({
            disabled: !this.state.disabled
        });
    }

    onChange(e) {
        let fieldName = e.target.name;
        fieldName = fieldName.charAt(0).toLowerCase() + fieldName.slice(1);

        let order = Object.assign({}, this.state.order);

        order[fieldName] = e.target.value;

        this.setState({
            order: order
        });
    }

    onSaveClick(e) {
        let form = e.target.parentNode;
        let order = serialize(form, {hash: true});

        order.Code = this.state.order.code;
        order.PhoneNumber = parseInt(order.PhoneNumber);

        this.props.updateOrder(order, this.state.order.id);
    }

    render() {
        let order = this.state.order,
            code = fmtCode(order.code),
            formActive = this.state.active ? 'is-active' : '',
            buttonActive = this.state.disabled ? '' : 'is-active',
            options = genStatusSelect(),
            disabled = this.state.disabled && order.status !== 'Доставлен';

        return (
            <form className={'form form_edit pure-form ' + formActive}>
                <button className='pure-button' type='button' onClick={::this.onEditClick}>Редкатировать</button>

                <fieldset className='pure-group'>
                    <input
                        value={code}

                        type='number'
                        className='pure-input-1-2'
                        placeholder='Идентификатор заказа'
                        required
                        disabled
                    />

                    <input
                        onChange={this.onChange.bind(this)}
                        value={order.sendAddress}
                        disabled={disabled}

                        name='SendAddress'
                        type='text'
                        className='pure-input-1-2'
                        placeholder='Адрес отправителя'
                        maxlength='8'
                        minlength='1'
                        required
                    />

                    <input
                        onChange={this.onChange.bind(this)}
                        value={this.state.order.recipientAddress}
                        disabled={disabled}

                        name='RecipientAddress'
                        type='text'
                        className='pure-input-1-2'
                        placeholder='Адрес получателя'
                        required
                    />

                    <input
                        onChange={this.onChange.bind(this)}
                        value={this.state.order.phoneNumber}
                        disabled={disabled}

                        name='PhoneNumber'
                        type='number'
                        className='pure-input-1-2'
                        placeholder='Телефон получателя'
                        maxlength='9'
                        minlength='9'
                        required
                    />
                </fieldset>

                <select
                    onChange={this.onChange.bind(this)}
                    value={this.state.order.status}
                    disabled={this.state.disabled}
                    name='Status'
                >
                    {options}
                </select>

                <button
                    onClick={this.onSaveClick.bind(this)}
                    className={'form_btn pure-button pure-button-primary ' + buttonActive}
                    type='button'
                >
                    Сохранить
                </button>
            </form>
        );
    }
}

OrderEdit.propTypes = {
    order: PropTypes.object.isRequired
}
