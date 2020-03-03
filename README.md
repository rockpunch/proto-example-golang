# proto-example-golang

### Updating Protocol Rules
1. Don't change the numeric tags for any existing fields.
2. You can add new fields, and old code will just ignore them.
3. Likewise, if the old/new code reads unknown data, the default will take place.
4. Fields can be removed, as long as the tag number is not used again 
    in your updated message type. You may want to rename the field instead,
    perhaps adding the prefix "OBSOLETE_", or make the tag RESERVED, so that
    future users of your .proto can't accidentally reuse the number.
5. For data type changes (int32 to int64 for example, please refer to the documentation)  