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

    onChangeHandle(e) {
        let form = $(e.target).closest('.form');

        form.find('button').prop('disabled', !form[0].checkValidity());
    }

    addOrder(e) {
        let form = $(e.target).closest('.form');
        let order = serialize(form, {hash: true});

        order.Code = parseInt(order.Code);
        order.PhoneNumber = parseInt(order.PhoneNumber);

        console.log(order);

        this.props.createOrder(order);
    }

    render() {
        let formActive = this.state.active ? 'is-active' : '',
            options = genStatusSelect();

        return (
            <div className='content-new'>
                <button className='pure-button' type='button' onClick={::this.showForm}>+ Добавить заказ</button>
                <form className={'form form_edit pure-form pure-form-aligned ' + formActive}>
                    <fieldset>
                        <div className='pure-control-group'>
                            <label htmlFor='Code' title='Уникальный код от 1 до 99999999'>
                                Код заказа (?)
                            </label>
                            <input
                                onChange={::this.onChangeHandle}
                                name='Code'
                                type='number'
                                className='pure-input-1-2'
                                placeholder='Идентификатор заказа'
                                max='99999999'
                                min='1'
                                required
                            />
                        </div>

                        <div className='pure-control-group'>
                            <label htmlFor='SendAddress'>Адрес отправителя</label>
                            <input
                                onChange={::this.onChangeHandle}

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
                                onChange={::this.onChangeHandle}

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
                                onChange={::this.onChangeHandle}

                                name='PhoneNumber'
                                type='number'
                                className='pure-input-1-2'
                                placeholder='Телефон получателя'
                                min='70000000000'
                                max='89999999999'
                                required
                            />
                        </div>

                        <div className='pure-control-group'>
                            <label htmlFor='Status'>
                                Статус
                            </label>

                            <select
                                onChange={::this.onChangeHandle}
                                name='Status'
                                defaultValue='Готовится'
                                required
                            >
                                {options}
                            </select>
                        </div>

                        <div className='pure-controls'>
                            <button
                                onClick={::this.addOrder}
                                className='form_btn pure-button pure-button-primary is-active'
                                type='button'
                                disabled
                            >
                                Добавить
                            </button>
                        </div>
                    </fieldset>
                </form>
            </div>
        );
    }
}
