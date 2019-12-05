/**
 * Autogenerated by Avro
 *
 * DO NOT EDIT DIRECTLY
 */
package com.pharbers.kafka.schema;

import org.apache.avro.specific.SpecificData;

@SuppressWarnings("all")
@org.apache.avro.specific.AvroGenerated
public class ConvertExcel extends org.apache.avro.specific.SpecificRecordBase implements org.apache.avro.specific.SpecificRecord {
  private static final long serialVersionUID = -6010935530319630724L;
  public static final org.apache.avro.Schema SCHEMA$ = new org.apache.avro.Schema.Parser().parse("{\"type\":\"record\",\"name\":\"ConvertExcel\",\"namespace\":\"com.pharbers.kafka.schema\",\"fields\":[{\"name\":\"traceId\",\"type\":\"string\"},{\"name\":\"assetId\",\"type\":\"string\"},{\"name\":\"type\",\"type\":\"string\"}]}");
  public static org.apache.avro.Schema getClassSchema() { return SCHEMA$; }
  @Deprecated public java.lang.CharSequence traceId;
  @Deprecated public java.lang.CharSequence assetId;
  @Deprecated public java.lang.CharSequence type;

  /**
   * Default constructor.  Note that this does not initialize fields
   * to their default values from the schema.  If that is desired then
   * one should use <code>newBuilder()</code>.
   */
  public ConvertExcel() {}

  /**
   * All-args constructor.
   * @param traceId The new value for traceId
   * @param assetId The new value for assetId
   * @param type The new value for type
   */
  public ConvertExcel(java.lang.CharSequence traceId, java.lang.CharSequence assetId, java.lang.CharSequence type) {
    this.traceId = traceId;
    this.assetId = assetId;
    this.type = type;
  }

  public org.apache.avro.Schema getSchema() { return SCHEMA$; }
  // Used by DatumWriter.  Applications should not call.
  public java.lang.Object get(int field$) {
    switch (field$) {
    case 0: return traceId;
    case 1: return assetId;
    case 2: return type;
    default: throw new org.apache.avro.AvroRuntimeException("Bad index");
    }
  }

  // Used by DatumReader.  Applications should not call.
  @SuppressWarnings(value="unchecked")
  public void put(int field$, java.lang.Object value$) {
    switch (field$) {
    case 0: traceId = (java.lang.CharSequence)value$; break;
    case 1: assetId = (java.lang.CharSequence)value$; break;
    case 2: type = (java.lang.CharSequence)value$; break;
    default: throw new org.apache.avro.AvroRuntimeException("Bad index");
    }
  }

  /**
   * Gets the value of the 'traceId' field.
   * @return The value of the 'traceId' field.
   */
  public java.lang.CharSequence getTraceId() {
    return traceId;
  }

  /**
   * Sets the value of the 'traceId' field.
   * @param value the value to set.
   */
  public void setTraceId(java.lang.CharSequence value) {
    this.traceId = value;
  }

  /**
   * Gets the value of the 'assetId' field.
   * @return The value of the 'assetId' field.
   */
  public java.lang.CharSequence getAssetId() {
    return assetId;
  }

  /**
   * Sets the value of the 'assetId' field.
   * @param value the value to set.
   */
  public void setAssetId(java.lang.CharSequence value) {
    this.assetId = value;
  }

  /**
   * Gets the value of the 'type' field.
   * @return The value of the 'type' field.
   */
  public java.lang.CharSequence getType() {
    return type;
  }

  /**
   * Sets the value of the 'type' field.
   * @param value the value to set.
   */
  public void setType(java.lang.CharSequence value) {
    this.type = value;
  }

  /**
   * Creates a new ConvertExcel RecordBuilder.
   * @return A new ConvertExcel RecordBuilder
   */
  public static com.pharbers.kafka.schema.ConvertExcel.Builder newBuilder() {
    return new com.pharbers.kafka.schema.ConvertExcel.Builder();
  }

  /**
   * Creates a new ConvertExcel RecordBuilder by copying an existing Builder.
   * @param other The existing builder to copy.
   * @return A new ConvertExcel RecordBuilder
   */
  public static com.pharbers.kafka.schema.ConvertExcel.Builder newBuilder(com.pharbers.kafka.schema.ConvertExcel.Builder other) {
    return new com.pharbers.kafka.schema.ConvertExcel.Builder(other);
  }

  /**
   * Creates a new ConvertExcel RecordBuilder by copying an existing ConvertExcel instance.
   * @param other The existing instance to copy.
   * @return A new ConvertExcel RecordBuilder
   */
  public static com.pharbers.kafka.schema.ConvertExcel.Builder newBuilder(com.pharbers.kafka.schema.ConvertExcel other) {
    return new com.pharbers.kafka.schema.ConvertExcel.Builder(other);
  }

  /**
   * RecordBuilder for ConvertExcel instances.
   */
  public static class Builder extends org.apache.avro.specific.SpecificRecordBuilderBase<ConvertExcel>
    implements org.apache.avro.data.RecordBuilder<ConvertExcel> {

    private java.lang.CharSequence traceId;
    private java.lang.CharSequence assetId;
    private java.lang.CharSequence type;

    /** Creates a new Builder */
    private Builder() {
      super(SCHEMA$);
    }

    /**
     * Creates a Builder by copying an existing Builder.
     * @param other The existing Builder to copy.
     */
    private Builder(com.pharbers.kafka.schema.ConvertExcel.Builder other) {
      super(other);
      if (isValidValue(fields()[0], other.traceId)) {
        this.traceId = data().deepCopy(fields()[0].schema(), other.traceId);
        fieldSetFlags()[0] = true;
      }
      if (isValidValue(fields()[1], other.assetId)) {
        this.assetId = data().deepCopy(fields()[1].schema(), other.assetId);
        fieldSetFlags()[1] = true;
      }
      if (isValidValue(fields()[2], other.type)) {
        this.type = data().deepCopy(fields()[2].schema(), other.type);
        fieldSetFlags()[2] = true;
      }
    }

    /**
     * Creates a Builder by copying an existing ConvertExcel instance
     * @param other The existing instance to copy.
     */
    private Builder(com.pharbers.kafka.schema.ConvertExcel other) {
            super(SCHEMA$);
      if (isValidValue(fields()[0], other.traceId)) {
        this.traceId = data().deepCopy(fields()[0].schema(), other.traceId);
        fieldSetFlags()[0] = true;
      }
      if (isValidValue(fields()[1], other.assetId)) {
        this.assetId = data().deepCopy(fields()[1].schema(), other.assetId);
        fieldSetFlags()[1] = true;
      }
      if (isValidValue(fields()[2], other.type)) {
        this.type = data().deepCopy(fields()[2].schema(), other.type);
        fieldSetFlags()[2] = true;
      }
    }

    /**
      * Gets the value of the 'traceId' field.
      * @return The value.
      */
    public java.lang.CharSequence getTraceId() {
      return traceId;
    }

    /**
      * Sets the value of the 'traceId' field.
      * @param value The value of 'traceId'.
      * @return This builder.
      */
    public com.pharbers.kafka.schema.ConvertExcel.Builder setTraceId(java.lang.CharSequence value) {
      validate(fields()[0], value);
      this.traceId = value;
      fieldSetFlags()[0] = true;
      return this;
    }

    /**
      * Checks whether the 'traceId' field has been set.
      * @return True if the 'traceId' field has been set, false otherwise.
      */
    public boolean hasTraceId() {
      return fieldSetFlags()[0];
    }


    /**
      * Clears the value of the 'traceId' field.
      * @return This builder.
      */
    public com.pharbers.kafka.schema.ConvertExcel.Builder clearTraceId() {
      traceId = null;
      fieldSetFlags()[0] = false;
      return this;
    }

    /**
      * Gets the value of the 'assetId' field.
      * @return The value.
      */
    public java.lang.CharSequence getAssetId() {
      return assetId;
    }

    /**
      * Sets the value of the 'assetId' field.
      * @param value The value of 'assetId'.
      * @return This builder.
      */
    public com.pharbers.kafka.schema.ConvertExcel.Builder setAssetId(java.lang.CharSequence value) {
      validate(fields()[1], value);
      this.assetId = value;
      fieldSetFlags()[1] = true;
      return this;
    }

    /**
      * Checks whether the 'assetId' field has been set.
      * @return True if the 'assetId' field has been set, false otherwise.
      */
    public boolean hasAssetId() {
      return fieldSetFlags()[1];
    }


    /**
      * Clears the value of the 'assetId' field.
      * @return This builder.
      */
    public com.pharbers.kafka.schema.ConvertExcel.Builder clearAssetId() {
      assetId = null;
      fieldSetFlags()[1] = false;
      return this;
    }

    /**
      * Gets the value of the 'type' field.
      * @return The value.
      */
    public java.lang.CharSequence getType() {
      return type;
    }

    /**
      * Sets the value of the 'type' field.
      * @param value The value of 'type'.
      * @return This builder.
      */
    public com.pharbers.kafka.schema.ConvertExcel.Builder setType(java.lang.CharSequence value) {
      validate(fields()[2], value);
      this.type = value;
      fieldSetFlags()[2] = true;
      return this;
    }

    /**
      * Checks whether the 'type' field has been set.
      * @return True if the 'type' field has been set, false otherwise.
      */
    public boolean hasType() {
      return fieldSetFlags()[2];
    }


    /**
      * Clears the value of the 'type' field.
      * @return This builder.
      */
    public com.pharbers.kafka.schema.ConvertExcel.Builder clearType() {
      type = null;
      fieldSetFlags()[2] = false;
      return this;
    }

    @Override
    public ConvertExcel build() {
      try {
        ConvertExcel record = new ConvertExcel();
        record.traceId = fieldSetFlags()[0] ? this.traceId : (java.lang.CharSequence) defaultValue(fields()[0]);
        record.assetId = fieldSetFlags()[1] ? this.assetId : (java.lang.CharSequence) defaultValue(fields()[1]);
        record.type = fieldSetFlags()[2] ? this.type : (java.lang.CharSequence) defaultValue(fields()[2]);
        return record;
      } catch (Exception e) {
        throw new org.apache.avro.AvroRuntimeException(e);
      }
    }
  }

  private static final org.apache.avro.io.DatumWriter
    WRITER$ = new org.apache.avro.specific.SpecificDatumWriter(SCHEMA$);

  @Override public void writeExternal(java.io.ObjectOutput out)
    throws java.io.IOException {
    WRITER$.write(this, SpecificData.getEncoder(out));
  }

  private static final org.apache.avro.io.DatumReader
    READER$ = new org.apache.avro.specific.SpecificDatumReader(SCHEMA$);

  @Override public void readExternal(java.io.ObjectInput in)
    throws java.io.IOException {
    READER$.read(this, SpecificData.getDecoder(in));
  }

}
