import React, { Component } from 'react'

export default class User extends Component {
    render() {
        const { active } = this.props
        const isActive = active ? '' : 'hidden'

        return (
            <div className={'content ' + isActive}>ghgfhgfhd</div>
        );
    }
}
