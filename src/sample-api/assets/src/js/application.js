window.sleep = (delay, callback) => setTimeout(callback, delay);


$.fn.exists = function() {
  return Boolean(this.length > 0);
};
