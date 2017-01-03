
/*
 * Import CSS
 **/
require('../css/application.scss');

window.sleep = (delay, callback) => setTimeout(callback, delay);


$.fn.exists = function() {
  return Boolean(this.length > 0);
};
