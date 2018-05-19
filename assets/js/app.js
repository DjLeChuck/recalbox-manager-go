// No need to wait for jQuery
var successClasses = "fa-check text-success";
var errorClasses = "fa-times-circle text-danger";

Dropzone.options.uploadBios = {
  paramName: "file",
  init: function () {
    this.on("success", function (file) {
      $.get({
        url: $("[data-url-check]").data("url-check"),
        dataType: "json",
        data: {
          file: file.name
        }
      }).done(function (data) {
        var $cell = $("[data-name='" + data.data.Name + "']");

        if (!$cell.length) {
          return;
        }

        var $icon = $cell.find("[data-fa-i2svg]");

        if (!$icon.length) {
          return;
        }

        if (data.data.IsValid) {
          $icon.removeClass(errorClasses).addClass(successClasses);
        } else {
          $icon.removeClass(successClasses).addClass(errorClasses);
        }
      });
    });
  }
};
