-- Author: Cael Shoop, cshoop2018@my.fit.edu
-- Course: CSE 4250, Fall 2020
-- Project: Proj2, Simply Sort 
-- Language implementation: GNAT Community 2020 (20200429-93)
-- GNATMAKE 10.2.0

with Ada.Text_IO; use Ada.Text_IO;
with Ada.Integer_Text_IO;
with Ada.Containers.Generic_Array_Sort;

procedure main is
    type Int_Array is array (Integer range <>) of Integer;
    size: Integer := 0;
    procedure sort_arr is new Ada.Containers.Generic_Array_Sort(
        Element_Type => Integer,
        Index_Type => Integer,
        Array_Type => Int_Array
    );
begin
    Ada.Integer_Text_IO.Get(size);
    declare
        arr : Int_Array(1..size);
    begin
        for i in 1..size loop
            Ada.Integer_Text_IO.Get(arr(i));
        end loop;
        sort_arr(arr);
        for i in 1..size loop
            Ada.Integer_Text_IO.Put(arr(i));
        end loop;
    end;
end main;