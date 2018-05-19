// No need to wait for jQuery
Dropzone.options.uploadBios = {
  paramName: "file",
  init: function () {
    this.on("success", function (event) {
      $.getJson("/", function (data) {
        console.log(data);
      });
    });
  }
};
