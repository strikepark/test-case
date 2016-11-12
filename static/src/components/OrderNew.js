import { findDOMNode } from 'react-dom'
import React, { PropTypes, Component } from 'react'
import { genStatusSelect } from '../helpers'

import serialize from 'form-serialize'
import $ from 'jquery'

export default class OrderNew extends Component {
    constructor(props) {
        super(props);

        this.state = {
            active: false
        };
    }

    showForm() {
        this.setState({
            active: !this.state.active
        });
    }

    onchange(e) {
        console.log(e);
    }

    addOrder(e) {
        let form = e.target.parentNode;
        let order = serialize(form, {hash: true});

        order.Code = parseInt(order.Code);
        order.PhoneNumber = parseInt(order.PhoneNumber);

        this.props.createOrder(order);
    }

    render() {
        let formActive = this.state.active ? 'is-active' : '',
            options = genStatusSelect(true);

        return (
            <form className={'form form_edit pure-form ' + formActive}>
                <button className='pure-button' type='button' onClick={this.showForm}>+ Добавить заказ</button>

                <fieldset className='pure-group'>
                    <input
                        onChange={this.onChange.bind(this)}
                        name='Code'
                        type='number'
                        className='pure-input-1-2'
                        placeholder='Идентификатор заказа'
                        maxlength='8'
                        minlength='1'
                        required
                    />

                    <input
                        onChange={this.onChange.bind(this)}

                        name='SendAddress'
                        type='text'
                        className='pure-input-1-2'
                        placeholder='Адрес отправителя'
                        required
                    />

                    <input
                        onChange={this.onChange.bind(this)}

                        name='RecipientAddress'
                        type='text'
                        className='pure-input-1-2'
                        placeholder='Адрес получателя'
                        required
                    />

                    <input
                        onChange={this.onChange.bind(this)}

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
                    name='Status'
                >
                    {options}
                </select>

                <button
                    onClick={this.addOrder.bind(this)}
                    className='form_btn pure-button pure-button-primary'
                    type='button'
                >
                    Добавить
                </button>
            </form>
        );
    }
}
