$(function () {
  var $volume = $("[data-volume]");

  $("#audio-volume").on("input change", function () {
    $volume.text($(this).val());
  });
});
