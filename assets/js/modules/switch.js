$(function () {
  $("[data-switch]").on("click", function () {
    $("#" + $(this).data("switch")).trigger("click");
  });
});
