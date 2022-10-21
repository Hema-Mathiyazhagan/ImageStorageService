import React from "react";
import "./App.css";
import CreateDelete  from "./CreateDelete";
import DeleteAlbum  from "./DeleteAlbum";
import CreateImage from "./CreateImage";
import DeleteImage from "./DeleteImage";
import Getimagefromalbum from "./Getimagefromalbum";
import GetAllImagesfromAlbum from "./GetAllmagesfromAlbum";

class App extends React.Component {
  state = {
    showCreateDelete: false,
    createAlbum:false,
    createImage:false,
    deleteImage:false,
    getSingleImage:false,
    getAllImages:false,
  }
  redirection = () => {
    this.setState({showCreateDelete: true})
    this.renderRedirect();
  }
  getSingleImageredirect = () => {
    this.setState({getSingleImage: true})
    this.renderSingleImage();
  }
  getAllImageredirect = () => {
    this.setState({getAllImages: true})
    this.renderAllImages();
  }
  createImageredirect = () => {
    this.setState({createImage: true})
    this.renderCreateImageredirect();
  }
  deleteImageredirect = () => {
    this.setState({deleteImage: true})
    this.renderDeleteImageredirect();
  }
  createredirection = () => {
    this.setState({createAlbum:true})
    this.renderCreateRedirect();
  }
  renderRedirect = () => {
    if(this.state.showCreateDelete){
      return <DeleteAlbum/>
    }
  }
  renderAllImages = () => {
    if(this.state.getAllImages){
      return <GetAllImagesfromAlbum/>
    }
  }
  renderSingleImage = () => {
    if(this.state.getSingleImage){
      return <Getimagefromalbum/>
    }
  }
  renderCreateImageredirect = () => {
    if(this.state.createImage){
      return <CreateImage/>
    }
  }
  renderDeleteImageredirect = () => {
    if(this.state.deleteImage){
      return <DeleteImage/>
    }
  }
  renderCreateRedirect = () => {
    if(this.state.createAlbum){
      return <CreateDelete/>
    }
  }
  renderCreateRedirect = () => {
    if(this.state.createAlbum){
      return <CreateDelete/>
    }
  }

  async getAlbumList() {
    let albumList = []
    albumList = await fetch(`http://localhost:3333/api/v1/albums`).then(res=>res.json());
    let a=Object.values(albumList);
    for(var i=0;i<a[0].length;i++){
       console.log((a[0]));
    }
    return a[0];
  }


  render() {
    // eslint-disable-next-line

    return (
      <div className="App">
        <h1>Image Store using Golang and React</h1>
        <div className="Homepage">
          <div className="album">
          <h2>Create and Delete Album</h2>
          <div className="albumbuttons">
        <button onClick={this.createredirection} variant="contained">create album</button>
        {this.renderCreateRedirect()}
        <button onClick={this.redirection} variant="contained">delete album</button>
        {this.renderRedirect()}
        </div>
        </div>
        <h2>Create Delete and retrieve images</h2>
        <div className="images">
        <button onClick={this.createImageredirect} variant="contained">create Image</button>
        {this.renderCreateImageredirect()}
        <button onClick={this.deleteImageredirect} variant="contained">Delete Image</button>
        {this.renderDeleteImageredirect()}
        <button onClick={this.getAllImageredirect} variant="contained">get all images in an album</button>
        {this.renderAllImages()}
          <button onClick={this.getSingleImageredirect} variant="contained">get single image in an album</button>
          {this.renderSingleImage()}
          </div>
        </div>
      </div>
    );
  }
}

export default App;