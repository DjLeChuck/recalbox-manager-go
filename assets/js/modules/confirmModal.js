$(function () {
  $("#confirmModal").on("show.bs.modal", function (event) {
    var $button = $(event.relatedTarget);
    var $modal = $(this);
    var content = $button.data("message");

    $modal.find("[data-href]").attr("href", $button.data('href'));
    $modal.find(".modal-body").html(
      content.replace("%s", "<strong>" + $button.data("message-arg") + "</strong>")
    );
  });
});
