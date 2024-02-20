import React, { Component } from 'react'

export class Item extends Component {
  render() {
    return (
      <div className='item'>
        <h2>{this.props.item.title}</h2>
      </div>
    )
  }
}

export default Item