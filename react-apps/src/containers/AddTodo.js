import React from 'react'
import { connect } from 'react-redux'
import { addTodo } from '../actions'
import TextField from 'material-ui/TextField';
import DatePicker from 'material-ui/DatePicker';
import TimePicker from 'material-ui/TimePicker';
import RaisedButton from 'material-ui/RaisedButton';

let AddTodo = ({ dispatch }) => {
  return (
    <div>
      <form
        onSubmit={e => {
          e.preventDefault()
          if (!document.getElementById("taskVal").value.trim()) {
            return
          }
          const obj = {body: document.getElementById('taskVal').value};
          const method = "POST";
          const body = JSON.stringify(obj);
          const headers = {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          };
          fetch("https://thuv9pflq4.execute-api.ap-northeast-1.amazonaws.com/dev/utakata/setItem", {method, headers, body})
          .then(response => {
            dispatch(addTodo(document.getElementById("taskVal").value.trim()))
            document.getElementById("taskVal").value = 
            document.getElementById("taskDate").value = ''
            document.getElementById("taskTime").value = ''
            return response.json();
          })
          .then(json => {
            console.log(json);
          });
        }}
      >
      <TextField  id="taskVal" floatingLabelText="タスク" fullWidth={true} />
      <DatePicker id="taskDate" hintText="期限日" floatingLabelText="期限日" fullWidth={true}/>
      <TimePicker id="taskTime" format="24hr" hintText="時刻" floatingLabelText="時刻" minutesStep={5} fullWidth={true}/>
      <br/>
      <RaisedButton type="submit" label="タスク登録" primary={true} fullWidth={true}/>
      </form>
    </div>
  )
}
AddTodo = connect()(AddTodo)

export default AddTodo
