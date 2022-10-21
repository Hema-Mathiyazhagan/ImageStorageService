import "./App.css";
import React from "react";
import { useState } from "react";

function DeleteImage() {
  const [albumName, setalbumname] = useState("");
  const [imageName, setImagename] = useState("");
  // eslint-disable-next-line
  const [message, setMessage] = useState("");

  let handleSubmit = async (e) => {
    e.preventDefault();
    fetch("http://localhost:3333/api/v1/albums/"+ albumName.trim()+"/images/"+imageName.trim(), {
        method: "DELETE",
      })
      .then((response)=> response.json())
      .then((data)=>{
        if(data.message){
           alert(data.message);
        }
        else{
            alert(data.error);
        }}).catch((err)=>{console.log(err)});
  };

  return (
    <div className="App">
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={albumName}
          placeholder="albumname"
          onChange={(e) => setalbumname(e.target.value)}
        />
        <input
          type="text"
          value={imageName}
          placeholder="image name"
          onChange={(e) => setImagename(e.target.value)}
        />
        <button type="submit">Submit</button>

        <div className="message">{message ? <p>{message}</p> : null}</div>
      </form>
    </div>
  );
}

export default DeleteImage;
