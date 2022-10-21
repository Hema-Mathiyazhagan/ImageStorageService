import React, { Component } from 'react';
import './App.css';

class GetAllImagesfromAlbum extends Component {
    constructor() {
        super();
        this.test=this.test.bind(this);
        this.state = {
           albumName:''
        };
        this.handleInputChange = this.handleInputChange.bind(this);
    }
    handleInputChange(event){
        this.setState({albumName:event.target.value});
    }

    test() {
        /*global fetch */
        fetch("http://localhost:3333/api/v1/albums/"+this.state.albumName.trim()+"/images",{
            method:'GET'
        })
        .then((response)=>response.json())
           .then((data)=>{
            if(Object.values(data)[0].length>0){
               alert(data.imageList);
            }else{
                alert(data.error);
           }
           })
           .catch((err)=>{console.log(err)})
    }

    render() {
        return (
            <div>
            <input type="text" value={this.state.albumName} placeholder="albumname" onChange={this.handleInputChange}/>
            <button onClick={this.test} variant="Primary">GetImage</button>
          </div>
        );
    }
}
export default GetAllImagesfromAlbum;