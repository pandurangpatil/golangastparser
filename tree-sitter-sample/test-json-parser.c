#include <assert.h>
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <tree_sitter/api.h>
#include <tree_sitter/parser.h>

// Declare the `tree_sitter_json` function, which is
// implemented by the `tree-sitter-json` library.
TSLanguage *tree_sitter_go();

char* read_file_to_string(const char* filename) {
  // Open the file for reading
  FILE* file = fopen(filename, "r");
  if (!file) {
    printf("Failed to open file\n");
    return NULL;
  }

  // Determine the size of the file
  fseek(file, 0L, SEEK_END);
  long file_size = ftell(file);
  rewind(file);

  // Allocate a buffer to hold the file contents
  char* buffer = (char*) malloc(file_size + 1);
  if (!buffer) {
    printf("Failed to allocate memory\n");
    fclose(file);
    return NULL;
  }

  // Read the file contents into the buffer
  if (fread(buffer, sizeof(char), file_size, file) != file_size) {
    printf("Failed to read file\n");
    fclose(file);
    free(buffer);
    return NULL;
  }

  // Add a null terminator to the end of the buffer
  buffer[file_size] = '\0';

  // Close the file and return the buffer
  fclose(file);
  return buffer;
}

// Define a callback function to print each node's fields
void print_node_fields(const TSTree *tree, TSNode node, int depth, const char* source_code) {

    // if (ts_node_is_named(node)) {
        // Get the string content of the node
  const char *type = ts_node_type(node);
  char *node_string = ts_node_string(node);
    // Get the start and end positions of the node in the source code
  uint32_t start = ts_node_start_byte(node);
  uint32_t end = ts_node_end_byte(node);
  char token[end - start + 1];
  strncpy(token, &source_code[start], end - start);
  token[end - start] = '\0';
  TSPoint startPoint = ts_node_start_point(node);
  TSPoint endPoint = ts_node_end_point(node);

  // Print the node's type and text
  printf("----->  <%s> line no -> %d Token -> %s\n", type, startPoint.row , token);

  // Recursively print out each of the node's child nodes
  for (uint32_t i = 0, child_count = ts_node_child_count(node); i < child_count; i++) {
    TSNode child = ts_node_child(node, i);
    print_node_fields(tree, child, depth + 1, source_code);
  }

  // Free the string content after use
  free(node_string);
    // }
}

int main() {
  // Create a parser.
  TSParser *parser = ts_parser_new();

  // Set the parser's language (JSON in this case).
  ts_parser_set_language(parser, tree_sitter_go());

  // Build a syntax tree based on source code stored in a string.
  // const char *source_code = "package main\n\nfunc main() {\n    fmt.Println(\"Hello, world!\"\n}\n";

  char *source_code = read_file_to_string("/Users/pandurang/projects/golang/helloworld/hello.go");

  TSTree *tree = ts_parser_parse_string(
    parser,
    NULL,
    source_code,
    strlen(source_code)
  );

  // Get the root node of the syntax tree.
  TSNode root_node = ts_tree_root_node(tree);

  // Print the syntax tree as an S-expression.
  char *string = ts_node_string(root_node);
  printf("Syntax tree: %s\n", string);

  // Print out each field of the root node
  print_node_fields(tree, root_node, 0, source_code);

  // Free all of the heap-allocated memory.
  free(string);
  free(source_code);
  ts_tree_delete(tree);
  ts_parser_delete(parser);
  return 0;
}
