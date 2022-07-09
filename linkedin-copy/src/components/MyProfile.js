import styled from "styled-components";
import React from "react";
import { useState } from "react";
import axios from "axios";
import { useEffect } from "react";
import Cookies from "js-cookie";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import "./addConnection.css";
import PhoneInput from "react-phone-number-input";
import "react-phone-number-input/style.css";

const MyProfile = (props) => {
  const [disabledFN, setDisabledFN] = useState(true);
  const [disabledLN, setDisabledLN] = useState(true);
  const [disabledE, setDisabledE] = useState(true);
  const [disabledG, setDisabledG] = useState(true);
  const [disabledV, setDisabledV] = useState(true);
  const [disabledXP, setDisabledXP] = useState(true);
  const [disabledED, setDisabledED] = useState(true);
  const [disabledBD, setDisabledBD] = useState(true);
  const [disabledPN, setDisabledPN] = useState(true);
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [gender, setGender] = useState(0);
  const [experience, setExperience] = useState(0);
  const [education, setEducation] = useState(0);
  const [visibility, setVisible] = useState(0);
  const [birthday, setBirthday] = useState(null);
  const [phone, setPhone] = useState("");

  const handleEdit = async (origin, event) => {
    switch (origin) {
      case "FN":
        if (disabledFN === false) {
          const body = JSON.stringify({
            op: "replace",
            path: "/Name",
            value: firstName,
          });
          const res = await axios.patch(
            "http://localhost:9090/update/user",
            body,
            {
              withCredentials: true,
            }
          );
          console.log(res.status);
        }
        setDisabledFN(!disabledFN);
        break;
      case "LN":
        if (disabledLN === false) {
          const body = JSON.stringify({
            op: "replace",
            path: "/Foo",
            value: lastName,
          });
          /* const res = await axios.patch(
            "http://localhost:9090/update/user",
            body,
            {
              withCredentials: true,
            }
          );*/
          //console.log(res.status);
        }
        setDisabledLN(!disabledLN);
        break;
      case "E":
        if (disabledE === false) {
          const body = JSON.stringify({
            op: "replace",
            path: "/Email",
            value: email,
          });
          const res = await axios.patch(
            "http://localhost:9090/update/user",
            body,
            {
              withCredentials: true,
            }
          );
          console.log(res.status);
        }
        setDisabledE(!disabledE);
        break;
      case "G":
        if (disabledG === false) {
          console.log(gender);
          const body = JSON.stringify({
            op: "replace",
            path: "/Gender",
            value: Number(gender),
          });
          const res = await axios.patch(
            "http://localhost:9090/update/user",
            body,
            {
              withCredentials: true,
            }
          );
          console.log(res.status);
        }
        setDisabledG(!disabledG);
        break;
      case "V":
        if (disabledV === false) {
          console.log(Boolean(visibility));
          const body = JSON.stringify({
            op: "replace",
            path: "/Public",
            value: Boolean(visibility),
          });
          const res = await axios.patch(
            "http://localhost:9090/update/user",
            body,
            {
              withCredentials: true,
            }
          );
          console.log(res.status);
        }
        setDisabledV(!disabledV);
        break;
      case "XP":
        if (disabledXP === false) {
          console.log(experience);
          const body = JSON.stringify({
            op: "replace",
            path: "/Experience",
            value: Number(experience),
          });
          const res = await axios.patch(
            "http://localhost:9090/update/user",
            body,
            {
              withCredentials: true,
            }
          );
          console.log(res.status);
        }
        setDisabledXP(!disabledXP);
        break;
      case "ED":
        if (disabledED === false) {
          const body = JSON.stringify({
            op: "replace",
            path: "/Education",
            value: Number(education),
          });
          const res = await axios.patch(
            "http://localhost:9090/update/user",
            body,
            {
              withCredentials: true,
            }
          );
          console.log(res.status);
        }
        setDisabledED(!disabledED);
        break;
      case "BD":
        if (disabledBD === false) {
          const body = JSON.stringify({
            op: "replace",
            path: "/DateOfBirth",
            value: birthday,
          });
          const res = await axios.patch(
            "http://localhost:9090/update/user",
            body,
            {
              withCredentials: true,
            }
          );
          console.log(res.status);
        }
        setDisabledBD(!disabledBD);
        break;
      case "PN":
        if(disabledPN === false){
          console.log(phone);
          const body = JSON.stringify({
            op: "replace",
            path: "/PhoneNumber",
            value: phone,
          });
          const res = await axios.patch(
            "http://localhost:9090/update/user",
            body,
            {
              withCredentials: true,
            }
          );
          console.log(res.status);
        }
        setDisabledPN(!disabledPN);
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
      setFirstName(userRes.data.Name);
      setLastName("Stojanovic");
      setEmail(userRes.data.Email);
      setGender(userRes.data.Gender);
      if (userRes.data.Gender === 1) {
        const a = document.getElementById("female");
        a.checked = true;
      } else {
        const b = document.getElementById("male");
        b.checked = true;
      }
      setExperience(userRes.data.Experience);
      setEducation(userRes.data.Education);
      setVisible(Number(userRes.data.Public));
      if (userRes.data.Public === true) {
        const a = document.getElementById("public");
        a.checked = true;
      } else {
        const b = document.getElementById("private");
        b.checked = true;
      }
      let d = new Date(String(userRes.data.DateOfBirth));
      setBirthday(d);
      setPhone(userRes.data.PhoneNumber);
    };
    userInfo();
  }, []);
  return (
    <Container>
      <Profile>
        <h1>My Profile</h1>
        <form id="edit">
          <Box>
            <GeneralLabel>First Name: </GeneralLabel>
            <GeneralInput
              id="username1"
              type="text"
              name="username1"
              placeholder="First Name"
              value={firstName}
              disabled={disabledFN}
              onChange={(e) => setFirstName(e.target.value)}
            />
            <EditLabel
              htmlFor="username"
              onClick={(event) => handleEdit("FN", event)}
            >
              Edit
            </EditLabel>
          </Box>
          <Box>
            <GeneralLabel>Last Name: </GeneralLabel>
            <GeneralInput
              id="lastName"
              type="text"
              name="lastName"
              disabled={disabledLN}
              value={lastName}
              placeholder = "Last Name"
              onChange={(e) => setLastName(e.target.value)}
            />
            <EditLabel
              htmlFor="lastName"
              onClick={(event) => handleEdit("LN", event)}
            >
              Edit
            </EditLabel>
          </Box>
          <Box>
            <EmailLabel>Email: </EmailLabel>
            <GeneralInput
              id="email"
              type="email"
              name="email"
              disabled={disabledE}
              value={email}
              placeholder = "Email"
              onChange={(e) => setEmail(e.target.value)}
            />
            <EditLabel
              htmlFor="email"
              onClick={(event) => handleEdit("E", event)}
            >
              Edit
            </EditLabel>
          </Box>
          <Box>
            <GenderLabel>Gender: </GenderLabel>
            <GeneralLabel for="male">Male</GeneralLabel>
            <Radio
              type="radio"
              id="male"
              name="gender"
              value={0}
              disabled={disabledG}
              onChange={(e) => setGender(e.target.value)}
            ></Radio>
            <GeneralLabel for="female">Female</GeneralLabel>
            <Radio
              type="radio"
              id="female"
              name="gender"
              value={1}
              disabled={disabledG}
              onChange={(e) => setGender(e.target.value)}
            ></Radio>
            <EditLabel
              htmlFor="gender"
              onClick={(event) => handleEdit("G", event)}
            >
              Edit
            </EditLabel>
          </Box>
          <Box>
            <VisibilityLabel>Visibility: </VisibilityLabel>
            <GeneralLabel for="public">Public</GeneralLabel>
            <Radio
              type="radio"
              id="public"
              name="visibility"
              value={1}
              disabled={disabledV}
              onChange={(e) => setVisible(e.target.value)}
            ></Radio>
            <GeneralLabel for="private">Private</GeneralLabel>
            <Radio
              type="radio"
              id="private"
              name="visibility"
              value={0}
              disabled={disabledV}
              onChange={(e) => setVisible(e.target.value)}
            ></Radio>
            <EditLabel
              htmlFor="visibility"
              onClick={(event) => handleEdit("V", event)}
            >
              Edit
            </EditLabel>
          </Box>
          <Box>
            <GeneralLabel>Experience: </GeneralLabel>
            <ExperienceInput
              id="xp"
              type="number"
              name="experience"
              disabled={disabledXP}
              min={0}
              max={70}
              value={experience}
              onChange={(e) => setExperience(e.target.value)}
            />
            <EditLabel
              htmlFor="experience"
              onClick={(event) => handleEdit("XP", event)}
            >
              Edit
            </EditLabel>
          </Box>
          <Box>
            <EducationLabel>Education level: </EducationLabel>
            <EducationSelector
              name="education"
              value={education}
              onChange={(e) => setEducation(e.target.value)}
              disabled={disabledED}
            >
              <option value={0}>Lower secondary</option>
              <option value={1}>Upper secondary</option>
              <option value={2}>Bachelors</option>
              <option value={3}>Masters</option>
              <option value={4}>Doctoral</option>
            </EducationSelector>
            <EditLabel
              htmlFor="education"
              onClick={(event) => handleEdit("ED", event)}
            >
              Edit
            </EditLabel>
          </Box>
          <Box>
            <BirthdayLabel htmlFor="birthday">Birthday: </BirthdayLabel>
            <DatePicker
              name="birthday"
              id="birthday"
              selected={birthday}
              onChange={(date) => setBirthday(date)}
              dateFormat="dd/MM/yyyy"
              maxDate={new Date()}
              showYearDropdown
              scrollableMonthYearDropdown
              fixedHeight
              disabled={disabledBD}
              className="prelepi"
            />
            <EditLabel
              htmlFor="birthday"
              onClick={(event) => handleEdit("BD", event)}
            >
              Edit
            </EditLabel>
          </Box>
          <Box>
            <PhoneLabel htmlFor="phonenumber">Phone number: </PhoneLabel>
            <PhoneInput
              className="phone"
              placeholder="Phone number"
              value={phone}
              onChange={setPhone}
              disabled={disabledPN}
            ></PhoneInput>
            <EditLabel
              htmlFor="phonenumber"
              onClick={(event) => handleEdit("PN", event)}
            >
              Edit
            </EditLabel>
          </Box>
        </form>
      </Profile>
    </Container>
  );
};
const PhoneLabel = styled.label`
  margin-left: 10px;
`;
const BirthdayLabel = styled.label`
  margin-left: 15px;
  margin-right: 100px;
`;
const ExperienceInput = styled.input`
  width: 50px;
  height: 25px;
  margin-left: 60px;
  margin-right: 90px;
  font-size: 16px;
`;
const EducationLabel = styled.label`
  margin-left: 10px;
  margin-right: 15px;
`;
const EducationSelector = styled.select`
  height: 30px;
  width: 150px;
  font-size: 15px;
  margin-left: 23px;
  option {
    font-size: 15px;
  }
`;
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
const Container = styled.div`
  padding: 0;
  height: 100vh;
  display: flex;
  justify-content: center;
`;
const Profile = styled.div`
  margin-top: 150px;
  width: 600px;
  height: 550px;
  background-color: white;
  border: none;
  border-radius: 10px 10px 10px 10px;
  box-shadow: 0 0 0 1px rgb(0 0 0 / 15%), 0 0 0 rgb(0 0 0 / 20%);
  display: flex;
  flex-direction: column;

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
