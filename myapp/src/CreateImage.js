import "./App.css";
import React from "react";
import "bootstrap/dist/css/bootstrap.min.css";
export default class ImageUpload extends React.Component{
    constructor(props){
        super(props);
        this.test=this.test.bind(this);
        this.state={
            fileUploadOngoing:false,
            albumName:''
        };
        this.handleInputChange = this.handleInputChange.bind(this)
    }
    handleInputChange(event){
        this.setState({albumName:event.target.value});
    }
    test(){
        this.setState = ({fileUploadOngoing:true});
        const fileInput = document.querySelector('#fileInput');
        const formData = new FormData();
        formData.append("file", fileInput.files[0]);
        const options = {
            method:"POST",
            body:formData
        };
        fetch("http://localhost:3333/api/v1/albums/"+this.state.albumName.trim()+"/images",options)
        .then((response)=> response.json())
        .then((data)=>{
          if(data.message){
          alert(data.message);
          }
          else{
              alert(data.error);
          }})
        .catch((err)=>{console.log(err.values)});
    }
    render() {
        console.log("this.state.fileUploadOngoing=" + this.state.fileUploadOngoing);
        return (
          <div>
          <input type="text" value={this.state.albumName} placeholder="albumname" onChange={this.handleInputChange}/>
            <input id="fileInput" type="file" name="file" />
            <button onClick={this.test} variant="Primary">Upload</button>   
          </div>
        );
      }
    
}