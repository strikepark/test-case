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
            readOnly = this.state.disabled,
            editable = order.status === 'Доставлен' ? 'hidden' : '';

        return (
            <form className={'form form_edit pure-form pure-form-aligned ' + formActive}>
                <fieldset>
                    <div className='pure-controls'>
                        <button
                            className={'button-xsmall button-secondary pure-button ' + editable}
                            type='button'
                            onClick={::this.onEditClick}>
                            Редкатировать
                        </button>
                    </div>

                    <div className='pure-control-group'>
                        <label htmlFor='Code' title='Уникальный код от 1 до 99999999'>
                            Код заказа (?)
                        </label>

                        <input
                            value={code}

                            type='number'
                            className='pure-input-1-2'
                            placeholder='Идентификатор заказа'
                            required
                            readOnly
                        />
                    </div>

                    <div className='pure-control-group'>
                        <label htmlFor='SendAddress'>Адрес отправителя</label>
                        <input
                            onChange={this.onChange.bind(this)}
                            value={order.sendAddress}
                            readOnly={readOnly}

                            name='SendAddress'
                            type='text'
                            className='pure-input-1-2'
                            placeholder='Адрес отправителя'
                            required
                        />
                    </div>

                    <div className='pure-control-group'>
                        <label htmlFor='RecipientAddress'>Адрес получателя</label>
                        <input
                            onChange={this.onChange.bind(this)}
                            value={this.state.order.recipientAddress}
                            readOnly={readOnly}

                            name='RecipientAddress'
                            type='text'
                            className='pure-input-1-2'
                            placeholder='Адрес получателя'
                            required
                        />
                    </div>

                    <div className='pure-control-group'>
                        <label htmlFor='PhoneNumber' title='Номер состоит из 11 цифр, например 89203002023'>
                            Телефон (?)
                        </label>
                        <input
                            onChange={this.onChange.bind(this)}
                            value={this.state.order.phoneNumber}
                            readOnly={readOnly}

                            name='PhoneNumber'
                            type='number'
                            className='pure-input-1-2'
                            placeholder='Телефон получателя'
                            max='9'
                            min='9'
                            required
                        />
                    </div>

                    <div className='pure-control-group'>
                        <label htmlFor='Status'>
                            Статус
                        </label>

                        <select
                            onChange={this.onChange.bind(this)}
                            value={this.state.order.status}
                            readOnly={readOnly}
                            disabled={readOnly}
                            name='Status'
                        >
                            {options}
                        </select>
                    </div>

                    <div className='pure-controls'>
                        <button
                            onClick={this.onSaveClick.bind(this)}
                            className={'form_btn pure-button pure-button-primary ' + buttonActive}
                            type='button'
                        >
                            Сохранить
                        </button>
                    </div>
                </fieldset>
            </form>
        );
    }
}

OrderEdit.propTypes = {
    order: PropTypes.object.isRequired
}
