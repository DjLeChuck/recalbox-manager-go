$(function () {
  $("#confirmModal").on("show.bs.modal", function (event) {
    var $link = $(event.relatedTarget);
    var $modal = $(this);
    var content = $link.data("message");

    $modal.find("[data-href]").attr("href", $link.attr("href"));
    $modal.find(".modal-body").html(
      content.replace("%s", "<strong>" + $link.data("message-arg") + "</strong>")
    );
  });
});
