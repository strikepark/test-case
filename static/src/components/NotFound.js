import React, { Component } from 'react'

import { Link } from 'react-router';

export default class NotFound extends Component {
    render() {
        return (
            <div className='login'>
                <h1>Страница не найдена</h1>

                <Link to='/'>На главную →</Link>
            </div>
        );
    }
}
