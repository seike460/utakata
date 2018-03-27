import React, {Component} from 'react';
import {BottomNavigation, BottomNavigationItem} from 'material-ui/BottomNavigation';
import Paper from 'material-ui/Paper';
import Sort from 'material-ui/svg-icons/content/sort';
import Add from 'material-ui/svg-icons/content/add';

const sortIcon = <Sort />;
const addIcon = <Add />;

class Footer extends Component {
  state = {
    selectedIndex: 0,
  };

  select = (index) => this.setState({selectedIndex: index});

  render() {
    return (
      <Paper zDepth={1}>
        <BottomNavigation selectedIndex={this.state.selectedIndex}>
          <BottomNavigationItem
            label="list"
            icon={sortIcon}
            onClick={() => this.select(0)}
          />
          <BottomNavigationItem
            label="add"
            icon={addIcon}
            onClick={() => this.select(1)}
          />
        </BottomNavigation>
      </Paper>
    );
  }
}
export default Footer
