BEGIN {
   printf "#ifndef MESSAGES_HH\n";
   printf "#define MESSAGES_HH\n\n";
   printf "/****** GENERATED FILE - DO NOT EDIT! ******/\n\n";
   printf "#include <typedefs.hh>\n\n";
   printf "#define mnNULL 0\n";

   msg_number = 1;
   process = "";
}

/^\#message/ {
   if (process == "L") {
      printf "#define mn%s_Brief %d\n", $2, msg_number++;
      printf "#define mn%s_Full %d\n", $2, msg_number++;
   }
   else if (process == "O") {
      printf "#define mn%s_Desc %d\n", $2, msg_number++;
      printf "#define mn%s_Scan %d\n", $2, msg_number++;
   }
   else {
      printf "#define mn%s %d\n", $2, msg_number++;
   }
}

/^\#process/ {
   process = $2;
}

END {
   printf "\n";
   printf "extern const message_t messages[%d];\n", msg_number;
   printf "extern const char* message( msg_id_t, ... );\n\n";
   printf "#endif /* MESSAGES_HH */\n";
}
