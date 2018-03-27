import React, { Component } from 'react';
import './App.css';
import aws_exports from './aws-exports';
import Amplify from 'aws-amplify';
import { withAuthenticator } from 'aws-amplify-react';
import AddTodo from './containers/AddTodo'
import VisibleTodoList from './containers/VisibleTodoList'
import AppBarComposition from './components/AppBarComposition'
import Footer from './components/Footer'
import {Card, CardTitle, CardText} from 'material-ui/Card';
Amplify.configure(aws_exports); 

class App extends Component {
  render() {
    return (
        <div>
		  <AppBarComposition />
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
