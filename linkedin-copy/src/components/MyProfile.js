import styled from "styled-components";
import React from "react";
import { useState } from "react";
import axios from "axios";
import { useEffect } from "react";
import Cookies from "js-cookie";

const MyProfile = (props) => {
  const [disabledFN, setDisabledFN] = useState(true);
  const [disabledLN, setDisabledLN] = useState(true);
  const [disabledE, setDisabledE] = useState(true);
  const [disabledG, setDisabledG] = useState(true);
  const [disabledV, setDisabledV] = useState(true);
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [gender, setGender] = useState(0);

  const handleEdit = (origin, event) => {
    switch (origin) {
      case "FN":
        setDisabledFN(!disabledFN);
        break;
      case "LN":
        setDisabledLN(!disabledLN);
        break;
      case "E":
        setDisabledE(!disabledE);
        break;
      case "G":
        setDisabledG(!disabledG);
        break;
      case "V":
        setDisabledV(!disabledV);
        break;
      default:
        break;
    }
  };
  useEffect(() => {
    const userInfo = async () => {
      const userRes = await axios.get(
        "http://localhost:9090/users/" + Cookies.get("username")
      );
      console.log(userRes.data);
      setFirstName(userRes.data.username);
      setLastName("Stojanovic");
      setEmail(userRes.data.email);
      setGender(userRes.data.gender);
    };
    userInfo();
  }, []);
  return (
    <Container>
      <Profile>
        <h1>My Profile</h1>
        <form id="edit">
          <Box>
            <h2>
              <GeneralLabel>First Name: </GeneralLabel>
              <GeneralInput
                id="username1"
                type="text"
                name="username1"
                value={firstName}
                disabled={disabledFN}
              />
              <EditLabel
                htmlFor="username"
                onClick={(event) => handleEdit("FN", event)}
              >
                Edit
              </EditLabel>
            </h2>
          </Box>
          <Box>
            <h2>
              <GeneralLabel>Last Name: </GeneralLabel>
              <GeneralInput
                id="lastName"
                type="text"
                name="lastName"
                disabled={disabledLN}
                value={lastName}
              />
              <EditLabel
                htmlFor="lastName"
                onClick={(event) => handleEdit("LN", event)}
              >
                Edit
              </EditLabel>
            </h2>
          </Box>
          <Box>
            <h2>
              <EmailLabel>Email: </EmailLabel>
              <GeneralInput
                id="email"
                type="email"
                name="email"
                disabled={disabledE}
                value={email}
              />
              <EditLabel
                htmlFor="email"
                onClick={(event) => handleEdit("E", event)}
              >
                Edit
              </EditLabel>
            </h2>
          </Box>
          <Box>
            <Gender>
              <GenderLabel>Gender: </GenderLabel>
              <GeneralLabel for="male">Male</GeneralLabel>
              <Radio
                type="radio"
                id="male"
                name="gender"
                value={gender}
                disabled={disabledG}
              ></Radio>
              <GeneralLabel for="female">Female</GeneralLabel>
              <Radio
                type="radio"
                id="female"
                name="gender"
                value={gender}
                disabled={disabledG}
              ></Radio>
              <EditLabel
                htmlFor="gender"
                onClick={(event) => handleEdit("G", event)}
              >
                Edit
              </EditLabel>
            </Gender>
          </Box>
          <Box>
            <Gender>
              <VisibilityLabel>Visibility: </VisibilityLabel>
              <GeneralLabel for="public">Public</GeneralLabel>
              <Radio
                type="radio"
                id="public"
                name="visibility"
                value={gender}
                disabled={disabledV}
              ></Radio>
              <GeneralLabel for="private">Private</GeneralLabel>
              <Radio
                type="radio"
                id="private"
                name="visibility"
                value={gender}
                disabled={disabledV}
              ></Radio>
              <EditLabel
                htmlFor="visibility"
                onClick={(event) => handleEdit("V", event)}
              >
                Edit
              </EditLabel>
            </Gender>
          </Box>
        </form>
      </Profile>
    </Container>
  );
};
const VisibilityLabel = styled.label`
  margin-left: 10px;
  margin-right: 65px;
`;
const Gender = styled.h2``;
const GeneralInput = styled.input`
  margin-left: 50px;
  width: 150px;
  height: 25px;
  font-size: 15px;
`;
const GeneralLabel = styled.label`
  margin-right: 10px;
  margin-left: 10px;
`;
const GenderLabel = styled.label`
  margin-right: 85px;
  margin-left: 10px;
`;

const EmailLabel = styled.label`
  margin-right: 50px;
  margin-left: 10px;
`;
const Box = styled.div`
  display: flex;
  justify-content: space-around;
  margin-left: 100px;
  flex-direction: row;
  padding-top: 10px;
  padding-bottom: 10px;
  border-top: 2px solid rgba(0, 0, 0, 0.1);
  margin-right: 100px;
  align-items: baseline;
`;
const EditLabel = styled.label`
  padding: 5px;
  margin-left: 20px;
  font-weight: 500;
  font-size: 14px;
  border: 2px solid rgba(0, 0, 0, 0.3);
  border-radius: 5px 5px 5px 5px;
  &:hover {
    cursor: pointer;
    color: rgba(0, 0, 0, 0.3);
  }
`;
const Container = styled.div``;
const Profile = styled.div`
  position: absolute;
  top: 100px;
  left: 450px;
  align-self: center;
  width: 600px;
  height: 500px;
  background-color: white;
  border: none;
  border-radius: 5px 5px 5px 5px;
  box-shadow: 0 0 0 1px rgb(0 0 0 / 15%), 0 0 0 rgb(0 0 0 / 20%);

  h1 {
    margin-left: 100px;
    font-size: 25px;
    padding-top: 20px;
    padding-left: 20px;
    padding-right: 20px;
    margin-right: 100px;
    font-family: sans-serif;
    padding-bottom: 10px;
  }
`;

const Radio = styled.input`
  width: 15px;
  height: 15px;
  margin-left: 0px;
  margin-right: 0px;
`;

export default MyProfile;
