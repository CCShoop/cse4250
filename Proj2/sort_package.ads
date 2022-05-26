generic
   Size: Natural;
package Sort_Package with SPARK_Mode is
   subtype Array_Index is Integer range 0 .. Size;
   type Data_Array is array (Array_Index) of Integer
      with Predicate => Data_Array'First = 0;
   procedure Insertion_Sort (Data : in out Data_Array);
end Sort_Package