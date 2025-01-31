// source: pkg/model/deployment_source.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

goog.exportSymbol('proto.model.DeploymentSource', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.DeploymentSource = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.DeploymentSource, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.DeploymentSource.displayName = 'proto.model.DeploymentSource';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.DeploymentSource.prototype.toObject = function(opt_includeInstance) {
  return proto.model.DeploymentSource.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.DeploymentSource} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.DeploymentSource.toObject = function(includeInstance, msg) {
  var f, obj = {
    applicationDirectory: jspb.Message.getFieldWithDefault(msg, 1, ""),
    revision: jspb.Message.getFieldWithDefault(msg, 2, ""),
    applicationConfig: msg.getApplicationConfig_asB64(),
    applicationConfigFilename: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.DeploymentSource}
 */
proto.model.DeploymentSource.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.DeploymentSource;
  return proto.model.DeploymentSource.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.DeploymentSource} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.DeploymentSource}
 */
proto.model.DeploymentSource.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setApplicationDirectory(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setRevision(value);
      break;
    case 3:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setApplicationConfig(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setApplicationConfigFilename(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.DeploymentSource.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.DeploymentSource.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.DeploymentSource} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.DeploymentSource.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getApplicationDirectory();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRevision();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getApplicationConfig_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      3,
      f
    );
  }
  f = message.getApplicationConfigFilename();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * optional string application_directory = 1;
 * @return {string}
 */
proto.model.DeploymentSource.prototype.getApplicationDirectory = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.DeploymentSource} returns this
 */
proto.model.DeploymentSource.prototype.setApplicationDirectory = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string revision = 2;
 * @return {string}
 */
proto.model.DeploymentSource.prototype.getRevision = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.DeploymentSource} returns this
 */
proto.model.DeploymentSource.prototype.setRevision = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional bytes application_config = 3;
 * @return {string}
 */
proto.model.DeploymentSource.prototype.getApplicationConfig = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * optional bytes application_config = 3;
 * This is a type-conversion wrapper around `getApplicationConfig()`
 * @return {string}
 */
proto.model.DeploymentSource.prototype.getApplicationConfig_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getApplicationConfig()));
};


/**
 * optional bytes application_config = 3;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getApplicationConfig()`
 * @return {!Uint8Array}
 */
proto.model.DeploymentSource.prototype.getApplicationConfig_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getApplicationConfig()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.model.DeploymentSource} returns this
 */
proto.model.DeploymentSource.prototype.setApplicationConfig = function(value) {
  return jspb.Message.setProto3BytesField(this, 3, value);
};


/**
 * optional string application_config_filename = 4;
 * @return {string}
 */
proto.model.DeploymentSource.prototype.getApplicationConfigFilename = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.DeploymentSource} returns this
 */
proto.model.DeploymentSource.prototype.setApplicationConfigFilename = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


goog.object.extend(exports, proto.model);
