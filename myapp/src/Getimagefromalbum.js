import React, { Component } from 'react';
import './App.css';

class Getimagefromalbum extends Component {
    constructor() {
        super();
        this.test=this.test.bind(this);
        this.state = {
           albumName:'',
           imageName:'',
        };
        this.handleInputChange = this.handleInputChange.bind(this);
        this.handlefilenamechange = this.handlefilenamechange.bind(this);
    }
    handleInputChange(event){
        this.setState({albumName:event.target.value});
    }
    handlefilenamechange(event){
        this.setState({imageName:event.target.value});
    }

    test() {
        /*global fetch */
        fetch("http://localhost:3333/api/v1/albums/"+this.state.albumName.trim()+"/images/"+this.state.imageName.trim(),{
            method:'GET'
        })
            .then((response)=>{return response.blob();})
           .then((blob)=>{
            const href = window.URL.createObjectURL(blob);
            const link = document.createElement('a');
            link.href = href;
            link.setAttribute('download', this.state.imageName); //or any other extension
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
           })
           .catch((err)=>{
            return Promise.reject({ Error: 'Something Went Wrong', err });
           })
    }

    render() {
        return (
            <div>
            <input type="text" value={this.state.albumName} placeholder="albumname" onChange={this.handleInputChange}/>
            <input type="text"  value={this.state.imageName} placeholder="imagename" onChange={this.handlefilenamechange} />
            <button onClick={this.test} variant="Primary">GetImage</button>
          </div>
        );
    }
}
export default Getimagefromalbum;