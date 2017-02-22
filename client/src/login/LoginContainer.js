import { connect } from 'react-redux';
import { LoginForm } from './LoginForm';
import { authenticateUser } from './LoginActions';

function mapDispatchToProps(dispatch) {
  return ({
    onSubmit: (values) => {
      console.log('RESUBMIT');
      dispatch(authenticateUser(values.email, values.password));
    },
  });
}

function mapStateToProps() {
  return ({});
}

export const LoginContainer = connect(mapStateToProps, mapDispatchToProps)(LoginForm);

export default LoginContainer;
