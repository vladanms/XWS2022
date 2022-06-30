import React, { useContext } from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
} from "react-router-dom";
import "./App.css";
import Login from "./components/Login";
import Header from "./components/Header";
import Home from "./components/Home";
import Registration from "./components/Registration";
import Signin from "./components/Signin";
import { useState } from "react";
import AuthApi from "./AuthApi";
import ProtectedRoutes from "./ProtectedRoutes";
import AuthedRoute from "./AuthedRoute";
import MyProfile from "./components/MyProfile"
import UserProfile from "./components/UserProfile";

function App() {
  const [auth, setAuth] = useState(false);
  return (
    <div className="App">
      <AuthApi.Provider value={{ auth, setAuth }}>
        <Router>
          <AllRoutes />
        </Router>
      </AuthApi.Provider>
    </div>
  );
}

function AllRoutes() {
  return (
    <Routes>
      <Route element={<ProtectedRoutes />}>
        <Route exact path="/home" element={<><Header /><Home /></> }/>
        <Route exact path = "/myProfile" element ={<><Header/><MyProfile/></>}/>
      </Route>
      <Route element={<AuthedRoute/>}>
        <Route exact path="/" element={<Login />} />
        <Route exact path="/register" element={<Registration />} />
        <Route exact path="/signin" element={<Signin />} />
      </Route>
      <Route exact path="/user/profile" element={<><Header /><UserProfile/></>}/>
    </Routes>
  );
}

export default App;
