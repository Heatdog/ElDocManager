import React from "react";
import Header from "../components/Header";
import Footer from "../components/Footer";
import Items from "../components/Items";

class Main extends React.Component{
  constructor(props){
    super(props)
    this.state = {
      items: [
        {
          id: 1,
          title: 'Первый документ'
        },
      ]
    }
  }
  render() {
    return (
        <div className="wrapper">
          <Header />
          <Items items={this.state.items}/>
          <Footer />
        </div>
    );
  }
}

export default Main;