import styled from "styled-components";
import React from "react";
import { useState } from "react";
import "./addConnection.css";
import Cookies from "js-cookie";
import axios from "axios";
import { useEffect } from "react";

const UserInformation = (props) => {
  const [follow, setFollow] = useState("Follow");
  const [borderColor, setBorderColor] = useState("#0A66C2");
  const handleFollow = async (e) => {
    if (follow !== "Following") {
      if (follow === "Follow") {
        setFollow("Request Sent!");
        setBorderColor("grey");
        const res = await axios.post(
          "http://localhost:9090/follow/" + props.data.Username,
          {},
          { withCredentials: true }
        );
      } else {
        setFollow("Follow");
        setBorderColor("#0A66C2");
        await axios.delete(
          "http://localhost:9090/follow/request/" + props.data.Username,
          { withCredentials: true }
        );
      }
    }
  };
  useEffect(() => {
    const updateState = async () => {
      if (Cookies.get("username") !== undefined) {
        if (props.followingUser) {
          setFollow("Following");
          setBorderColor("grey");
        } else {
          const res = await axios.get(
            "http://localhost:9090/singleRequest/" + props.data.Username,
            { withCredentials: true }
          );
          console.log(res.status);
          if (res.status === 200) {
            setFollow("Request sent!");
            setBorderColor("grey");
          }
        }
      }
    };

    updateState();
  }, [props]);
  return (
    <Container>
      <ArtCard>
        <UserInfo>
          <CardBackground />
          <a>
            <Photo />
            <Link>{props.data.Username}</Link>
            <span>{props.data.Email}</span>
          </a>
        </UserInfo>
        {Cookies.get("username") !== undefined ? (
          <Widget>
            <a onClick={handleFollow}>
              <Follow
                style={{
                  borderColor: `${borderColor}`,
                  color: `${borderColor}`,
                }}
              >
                <h2>{follow}</h2>
              </Follow>
            </a>
          </Widget>
        ) : (
          <></>
        )}
        <Item>
          <span>
            <img src="/images/item-icon.svg" alt="" />
            My Items
          </span>
        </Item>
      </ArtCard>
    </Container>
  );
};

const Follow = styled.div`
  margin-left: 40px;
  margin-right: 40px;
  padding-top: 10px;
  padding-bottom: 10px;
  border: 2px solid #0a66c2;
  border-radius: 23px;
  color: #0a66c2;
  display: flex;
  flex-direction: row;
  align-items: center;

  &:hover {
    cursor: pointer;
    background-color: rgba(10, 102, 194, 0.2);
  }
`;
const Box = styled.div`
  z-index: 3;
  font-size: 16px;
  display: flex;
  padding: 20px 20px 20px 20px;
  position: absolute;
  top: 325px;
  left: 310px;
  border: 3px solid lightblue;
  border-left: 1px solid rgba(0, 0, 0, 0.15);
  border-radius: 0 5px 5px 0;
  background-color: white;
  width: 150px;
  height: 25px;
  justify-content: space-between;
  span {
    align-self: center;
    h3 {
      font-size: 12px;
      margin-bottom: 5px;
      padding: 2px 2px 2px 2px;
      border: 3px solid lightblue;
      border-radius: 5px 5px 5px 5px;
      background-color: aliceblue;
    }
    input {
      border: 2px solid black;
      border-radius: 2px;
      width: 100px;
      margin-bottom: 10px;
    }
  }
  a {
    border: 2px solid black;
    border-radius: 50%;
    padding-left: 3px;
    padding-right: 3px;
    padding-top: 3px;
    padding-bottom: 3px;
    transition-delay: 120ms;
    &:hover {
      background-color: lightblue;
      cursor: pointer;
    }
  }
`;

const Container = styled.div`
  grid-area: leftside;
`;

const ArtCard = styled.div`
  text-align: center;
  overflow: hidden;
  margin-bottom: 8px;
  background-color: #fff;
  border-radius: 5px;
  transition: box-shadow 83ms;
  position: relative;
  border: none;
  box-shadow: 0 0 0 1px rgb(0 0 0 / 15%), 0 0 0 rgb(0 0 0 / 20%);
`;

const UserInfo = styled.div`
  border-bottom: 1px solid rgba(0, 0, 0, 0.15);
  padding: 12px 12px 16px;
  word-wrap: break-word;
  word-break: break-word;
`;

const CardBackground = styled.div`
  background: url("/images/card-bg.svg");
  background-position: center;
  background-size: 462px;
  height: 54px;
  margin: -12px -12px 0;
`;

const Photo = styled.div`
  box-shadow: none;
  background-image: url("/images/photo.svg");
  width: 72px;
  height: 72px;
  box-sizing: border-box;
  background-clip: content-box;
  background-color: white;
  background-position: center;
  background-size: 60%;
  background-repeat: no-repeat;
  border: 2px solid white;
  margin: -38px auto 12px;
  border-radius: 50%;
`;

const Link = styled.div`
  font-size: 25px;
  line-height: 1.5;
  color: rgba(0, 0, 0, 0.9);
  font-weight: 600;
  margin: 5px;
`;

const Widget = styled.div`
  border-bottom: 1px solid rgba(0, 0, 0, 0.15);
  padding-top: 12px;
  padding-bottom: 12px;
  & > a {
    text-decoration: none;
    padding: 4px 12px;
    div {
      display: flex;
      flex-direction: column;
      text-align: left;
      span {
        font-size: 12px;
        line-height: 1.333;
        &:first-child {
          color: rgba(0, 0, 0, 0.6);
        }
        &:nth-child(2) {
          color: rgba(0, 0, 0, 1);
        }
      }
    }
  }
  svg {
    color: rgba(0, 0, 0, 1);
  }
`;

const Item = styled.a`
  border-color: rgba(0, 0, 0, 0.8);
  text-align: left;
  padding: 12px;
  font-size: 12px;
  display: block;
  span {
    display: flex;
    align-items: center;
    color: rgba(0, 0, 0, 1);
    svg {
      color: rgba(0, 0, 0, 0.6);
    }
  }
  &:hover {
    background-color: rgba(0, 0, 0, 0.08);
  }
`;

const CommunityCard = styled(ArtCard)`
  padding: 8px 0 0;
  text-align: left;
  display: flex;
  flex-direction: column;
  a {
    color: black;
    padding: 4px 12px 4px 12px;
    font-size: 12px;
    &:hover {
      color: #0a66c2;
    }
    span {
      display: flex;
      align-items: center;
      justify-content: space-between;
    }
    &:last-child {
      color: rgba(0, 0, 0, 0.6);
      text-decoration: none;
      border-top: 1px solid #d6cec2;
      padding: 12px;
      &:hover {
        background-color: rgba(0, 0, 0, 0.08);
      }
    }
  }
`;

export default UserInformation;
