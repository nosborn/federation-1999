BEGIN {
  printf "const (\n";

  msg_number = 1;
  process = "";
}

/^#message/ {
  if (process == "L") {
    printf "\t%s_Brief\n", $2; msg_number++;
    printf "\t%s_Full\n", $2; msg_number++;
  } else if (process == "O") {
    printf "\t%s_Desc\n", $2; msg_number++;
    printf "\t%s_Scan\n", $2; msg_number++;
  } else {
    if (msg_number == 1) {
      printf "\t%s MsgNum = %d + iota\n", $2, msg_number++;
    } else {
      printf "\t%s\n", $2; msg_number++;
    }
  }
}

/^#process/ {
   process = $2;
}

END {
  printf ")\n";
  printf "\n";
  printf "var maxMsgNum MsgNum = %d\n", msg_number - 1;
}
