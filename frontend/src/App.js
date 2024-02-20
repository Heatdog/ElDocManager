import React from "react";
import Main from "./pages/Main";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Login from "./pages/Login";
import NotFound from "./pages/NotFound";

class App extends React.Component{
  render() {
    return (
        <BrowserRouter>
            <Routes>
              <Route path="/" element={<Main />}/>
              <Route path="/login" element={<Login login="" password=""/>}/>
              <Route path="*" element={<NotFound />}/>
            </Routes>
        </BrowserRouter>
    );
  }
}

export default App;
