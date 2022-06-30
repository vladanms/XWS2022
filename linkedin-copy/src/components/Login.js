import styled from "styled-components";
import React, { useContext } from "react";
import { useNavigate } from "react-router-dom";
import Cookies from "js-cookie";
import AuthApi from "../AuthApi";
const Login = (props) => {
  const Auth = useContext(AuthApi);
  let navigate = useNavigate();
  const handleGuest = () => {
    Auth.setAuth(true);
    Cookies.set("guest", "guestUser");
    navigate("/home");
  }
  return (
    <Container>
      <Nav>
        <a href="/">
          <img src="/images/login-logo-dislinkt.png" alt="" />
        </a>
        <div>
          <Join href="/register">Join now</Join>
          <Signin href = "/signin">Sign in</Signin>
          <Guest onClick={handleGuest}>continue as guest</Guest>
        </div>
      </Nav>
      <Section>
        <Hero>
          <h1>Welcome to hell boi</h1>
          <img src="/images/login-hero-dislinkt.png" alt="" />
        </Hero>
      </Section>
    </Container>
  );
};

const Section = styled.section`
  align-content: start;
  display: flex;
  min-height: 700px;
  padding-bottom: 138px;
  padding-top: 40px;
  padding: 60px 0;
  position: relative;
  flex-wrap: wrap;
  width: 100%;
  max-width: 1128px;
  align-items: center;
  margin: auto;
`;
const Hero = styled.div`
  width: 100%;
  h1 {
    font-size: 56px;
    font-weight: 200;
    color: #2977c9;
    padding-bottom: 0;
    width: 55%;
    line-height: 70px;
  }
  img {
    // z-index: -1;
    width: 660px;
    height: 640px;
    position: absolute;
    bottom: 100px;
    right: -100px;
  }
`;

const Signin = styled.a`
  color: blue;
  font-size: 16px;
  font-weight: 600;
  padding-left: 25px;
  padding-right: 25px;
  padding-top: 12px;
  padding-bottom: 12px;
  border: 2px solid  #0a66c2;
  border-radius: 24px;
  transition-duration: 167ms;
  text-decoration: none;
  &:hover {
    background-color: rgba(10, 102, 194, 0.2);
  }
`;
const Join = styled.a`
  font-size: 16px;
  padding: 10px 12px;
  text-decoration: none;
  color: rgba(0, 0, 0, 0.6);
  border-radius: 4px;
  margin-right: 12px;
  transition-duration: 130ms;
  &:hover {
    background-color: rgba(0, 0, 0, 0.08);
    color: rgba(0, 0, 0, 0.9);
    text-decoration: none;
  }
`;
const Guest = styled.a`
  font-size: 16px;
  padding: 10px 12px;
  text-decoration: none;
  color: rgba(0, 0, 0, 0.6);
  border-radius: 4px;
  margin-left: 15px;
  transition-duration: 130ms;
  &:hover {
    background-color: rgba(0, 0, 0, 0.08);
    color: rgba(0, 0, 0, 0.9);
    cursor: pointer;
  }

`;

const Container = styled.div`
  padding: 0px;
  overflow-y: hidden;
  height: 100vh;
`;
const Nav = styled.nav`
  max-width: 1328px;
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
    @media (max-width: 768px) {
      padding: 0 5px;
    }
  }
`;

export default Login;
