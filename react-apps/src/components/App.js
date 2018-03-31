import React, { Component } from 'react';
import '../css/App.css';
import AddTodo from '../containers/AddTodo'
import VisibleTodoList from '../containers/VisibleTodoList'
import AppBar from './AppBar'
import Footer from './Footer'
import {Card, CardTitle, CardText} from 'material-ui/Card';
import Amplify from 'aws-amplify';
import aws_exports from '../config/aws-exports';
import { withAuthenticator } from 'aws-amplify-react';
Amplify.configure(aws_exports); 

class App extends Component {
  render() {
    return (
        <div>
		  <AppBar />
          <Card>
            <CardTitle title="タスク登録"/>
            <CardText>
              <AddTodo />
            </CardText>
            <CardTitle title="タスクリスト"/>
            <CardText>
              <VisibleTodoList />
            </CardText>
          </Card>
          <Footer />
        </div>
    );
  }
}
export default withAuthenticator(App);
