import styled from "styled-components";
import React from "react";
import "./registration.css";
import { useState } from "react";
import axios from "axios";
import { Navigate, useNavigate } from "react-router-dom";

const Registration = (props) => {
  const [Username, setUsername] = useState("");
  const [Email, setEmail] = useState("");
  const [Password, setPassword] = useState("");
  const handleUsername = (e) => {
    setUsername(e.target.value);
  };

  // Handling the email change
  const handleEmail = (e) => {
    setEmail(e.target.value);
  };

  // Handling the password change
  const handlePassword = (e) => {
    setPassword(e.target.value);
  };
  let navigate = useNavigate();
  const handleSubmit = async(e) => {
    e.preventDefault();
    const res = await axios.post(
      "http://localhost:9090/register",
      JSON.stringify({
        Username,
        Email,
        Password,
      })
    );
    if(res.status == 200){
    navigate("/home");
    }
  };

  return (
    <Container>
      <Nav>
        <a href="/">
          <img src="/images/login-logo-dislinkt.png" alt="" />
        </a>
      </Nav>
      <BodyWrapper>
      <Body>
        <div className="logo">
          <img className="img" src="/images/login-logo-dislinkt.png" alt="" />
        </div>
        <div className="form">
          <div className="form-body">
            <div className="username">
              <input
                className="form__input"
                type="text"
                id="firstName"
                placeholder="Username"
                onChange={handleUsername}
              />
            </div>
            <div className="email">
              <input
                type="email"
                id="email"
                className="form__input"
                placeholder="Email"
                onChange={handleEmail}
              />
            </div>
            <div className="password">
              <input
                className="form__input"
                type="password"
                id="password"
                placeholder="Password"
                onChange={handlePassword}
              />
            </div>
          </div>
          <div class="footer">
            <Submit type="submit" onClick={handleSubmit}>
              Register
            </Submit>
          </div>
        </div>
      </Body>
      </BodyWrapper>
    </Container>
  );
};
const BodyWrapper = styled.div`
  padding: 0;
  height: 100vh;
  display: flex;
  justify-content: center;
`;
const Submit = styled.button`
  box-shadow: 0 0 0 1px rgb(0 0 0 / 15%), 0 0 0 rgb(0 0 0 / 20%);
  border-radius: 24px;
  width: 200px;
  height: 50px;
  background-color: #0a66c2;
  font-size: 20px;
  color: white;
  font-weight: 600;
  transition-duration: 130ms;
  &:hover {
    cursor: pointer;
    background-color: rgb(10, 102, 160);
  }
`;
const Container = styled.div`
  padding: 0;
  height: 100vh;
`;

const Body = styled.div`
  overflow: hidden;
  margin-top: 50px;
  z-index: 1000;
  width: 350px;
  height: 450px;
  border-radius: 15px 15px 15px 15px;
  background-color: white;
  box-shadow: 0 0 0 1px rgb(0 0 0 / 15%), 0 0 0 rgb(0 0 0 / 20%);
  input {
      width: 200px;
      height: 40px;
      font-size: 16px;
      border: 2px solid grey;
      border-radius: 5px 5px 5px 5px;
      transition-duration: 130ms;
      &:hover {
          border-color: rgba(105, 105, 105, 0.3);
      }
  }
`;

const Nav = styled.nav`
  max-width: 1128px;
  margin: auto;
  padding: 12px 0 16px;
  display: flex;
  align-items: center;
  position: relative;
  justify-content: space-between;
  flex-wrap: nowrap;

  & > a {
    width: 135px;
    height: 90px;
    img {
      width: 200px;
      height: 100px;
    }
    @media (max-width: 768px) {
      padding: 0 5px;
    }
  }
`;

export default Registration;
