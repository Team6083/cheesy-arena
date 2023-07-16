// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Client-side logic for the scoring interface.

var websocket;
let alliance;

// Handles a websocket message to update the teams for the current match.
const handleMatchLoad = function(data) {
  $("#matchName").text(data.Match.LongName);
  if (alliance === "red") {
    $("#team1").text(data.Match.Red1);
    $("#team2").text(data.Match.Red2);
    $("#team3").text(data.Match.Red3);
  } else {
    $("#team1").text(data.Match.Blue1);
    $("#team2").text(data.Match.Blue2);
    $("#team3").text(data.Match.Blue3);
  }
};

// Handles a websocket message to update the match status.
const handleMatchTime = function(data) {
  switch (matchStates[data.MatchState]) {
    case "PRE_MATCH":
      // Pre-match message state is set in handleRealtimeScore().
      $("#postMatchMessage").hide();
      $("#commitMatchScore").hide();
      break;
    case "POST_MATCH":
      $("#postMatchMessage").hide();
      $("#commitMatchScore").css("display", "flex");
      break;
    default:
      $("#postMatchMessage").hide();
      $("#commitMatchScore").hide();
  }
};

// Handles a websocket message to update the realtime scoring fields.
var handleRealtimeScore = function(data) {
  var realtimeScore;
  if (alliance === "red") {
    realtimeScore = data.Red;
  } else {
    realtimeScore = data.Blue;
  }
  var score = realtimeScore.Score;

  console.log(score);

  $("#cube>.value").text(score.CubeStatus ? "Yes" : "No");
  $("#cube").attr("data-value", score.CubeStatus);

  for (let i = 0; i < 3; i++) {
    const i1 = i + 1;
    $(`#parkStatus${i1}>.value`).text(score.ParkStatuses[i] ? "Yes" : "No");
    $("#parkStatus" + i1).attr("data-value", score.ParkStatuses[i]);
  }

  $("#teleopCells").text(score.PowerCell);
  $("#teleopBlocks").text(score.Block);
};

// Handles a keyboard event and sends the appropriate websocket message.
var handleKeyPress = function(event) {
  websocket.send(String.fromCharCode(event.keyCode));
};

// Handles an element click and sends the appropriate websocket message.
const handleClick = function(command, teamPosition = 0, gridRow = 0, gridNode = 0, nodeState = 0) {
  console.log(`click_${command}`);
  websocket.send(command, {TeamPosition: teamPosition, GridRow: gridRow, GridNode: gridNode, NodeState: nodeState});
};

// Sends a websocket message to indicate that the score for this alliance is ready.
const commitMatchScore = function() {
  websocket.send("commitMatch");
  $("#postMatchMessage").css("display", "flex");
  $("#commitMatchScore").hide();
};

$(function() {
  alliance = window.location.href.split("/").slice(-1)[0];
  $("#alliance").attr("data-alliance", alliance);

  // Set up the websocket back to the server.
  websocket = new CheesyWebsocket("/panels/scoring/" + alliance + "/websocket", {
    matchLoad: function(event) { handleMatchLoad(event.data); },
    matchTime: function(event) { handleMatchTime(event.data); },
    realtimeScore: function(event) { handleRealtimeScore(event.data); },
  });

  $(document).keypress(handleKeyPress);
});
