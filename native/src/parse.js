'use strict';

const {
  createSourceFile,
  ScriptTarget,
  forEachChild
} = require('typescript');
const nodeKinds = require('./kinds');

/**
 * Parses the given TypeScript source code and returns a string with
 * the AST encoded as JSON.
 *
 * @function parse
 * @param {string} source - TypeScript source code to parse
 * @returns {Object} the representation of the parsed AST
 */
function parse(source) {
  let sourceFile = createSourceFile(
    'input.ts',
    source,
    ScriptTarget.ES6,
    false
  );

  const diagnostics = (sourceFile.parseDiagnostics || [])
    .filter(d => d.category === 1)
    .map(d => `at ${d.start}:${d.start+d.length}: ${d.messageText}`);

  // normalization of the AST and remove some properties that are not
  // useful in the output. Text we already had and fileName is not real
  // so there's no point in keeping it
  normalizeNode(sourceFile);
  delete sourceFile.text;
  delete sourceFile.fileName;
  delete sourceFile.parseDiagnostics;

  return {
    ast: sourceFile,
    diagnostics
  };
}

/**
 * Recursively normalizes the node and all its children. By normalization 
 * we mean replace the kind of the node by a string, instead of an integer
 * and also we remove the flags from the node.
 *
 * @function normalizeNode
 * @param {ts.Node} node - Node to normalize
 */
function normalizeNode(node) {
  forEachChild(node, normalizeNode);
  // node.kind substitution MUST be done after normalizing the children
  // otherwise forEachChild does not know how to access them
  node.kind = nodeKinds[node.kind];
  delete node.flags;
}

const OK = 'ok';
const FATAL = 'fatal';
const ERROR = 'error';

/**
 * Builds the result, either a successful or an errored one.
 *
 * @function printResult
 * @param {string} status - one of OK, FATAL or ERROR
 * @param {Object|null} ast - parsed AST, null if status is not ok
 * @param {Array<string>} diagnostics - List of diagnostic messages
 * @return {string} JSON encoded result
 */
function result(status, ast, ...diagnostics) {
  const result = {
    status,
    ast,
    errors: diagnostics.map(m => ({ message: m }))
  }

  try {
    return JSON.stringify(result);
  } catch (e) {
    return result(FATAL, null, `Unable to encode response to JSON: ${e.message}`);
  }
}

/**
 * Processes one single request to parse a TypeScript source file. Will print a
 * fatal or error result if the process fails, a valid one with the AST in any
 * other case.
 *
 * @function processRequest
 * @param {string} line - Line read from stdin containing the request
 * @param {string} the JSON encoded response output
 */
function processRequest(line) {
  let req;
  try {
    req = JSON.parse(line);
  } catch (e) {
    return result(FATAL, null, `Unable to decode request from JSON: ${e.message}`);
  }

  const { ast, diagnostics } = parse(req.content || '');
  if (diagnostics.length > 0) {
    return result(ERROR, null, ...diagnostics);
  }

  return result(OK, ast);
}

module.exports = {
  parse,
  processRequest
};
