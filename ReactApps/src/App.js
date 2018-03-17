import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to Utakata React</h1>
        </header>
    <main role="main" className="container" >
      <div className="d-flex align-items-center p-3 my-3 text-white-50 bg-purple rounded box-shadow">
      </div>
      <div className="my-3 p-3 bg-white rounded box-shadow" id="TaskListView">
        <h6 className="border-bottom border-gray pb-2 mb-0">Your Task</h6>
      </div>
      <div className="my-3 p-3 bg-white rounded box-shadow">
        <div className="row media text-muted pt-3">
            <div className="text-center mb-2 col-7">
              <input id="newTask" className="form-control" placeholder="new Task" required=""></input>
            </div>
            <div className="text-center mb-2 col-4">
              <input id="targetDate" className="form-control" placeholder="Target date" required=""></input>
            </div>
            <div className="text-center mb-2 col-1">
              <button /*onClick="setTask();"*/>Add </button>
            </div>
        </div>
      </div>
    </main>
    <template id="mytemplate">
    <div className="media text-muted pt-3">
      <p className="media-body pb-3 mb-0 small lh-125 border-bottom border-gray">
        Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus.
      </p>
    </div>
    </template>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default App;
