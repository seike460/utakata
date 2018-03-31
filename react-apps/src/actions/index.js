let nextTodoId = 0
export const addTodo = text => {
  return {
    type: 'ADD_TODO',
    id: nextTodoId++,
    text
  }
}

export const getTodo = () => {
  const json = (async () => {
    try {
      const url = 'https://thuv9pflq4.execute-api.ap-northeast-1.amazonaws.com/dev/utakata/getItem';
      const response = await fetch(url);
      const json = await response.json();
      console.log(json.origin);
      return json
    } catch (error) {
      console.log(error);
    }
  })();
  return {
    type: 'GET_TODO',
    payload: json
  }
}

export const setVisibilityFilter = filter => {
  return {
    type: 'SET_VISIBILITY_FILTER',
    filter
  }
}

export const toggleTodo = id => {
  return {
    type: 'TOGGLE_TODO',
    id
  }
}

export const VisibilityFilters = {
  SHOW_ALL: 'SHOW_ALL',
  SHOW_COMPLETED: 'SHOW_COMPLETED',
  SHOW_ACTIVE: 'SHOW_ACTIVE'
}
