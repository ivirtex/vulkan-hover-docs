# FrontFacing(3) Manual Page

## Name

FrontFacing - Front face determination of a fragment



## <a href="#_description" class="anchor"></a>Description

`FrontFacing`  
Decorating a variable with the `FrontFacing` built-in decoration will
make that variable contain whether the fragment is front or back facing.
This variable is non-zero if the current fragment is considered to be
part of a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-polygons-basic"
target="_blank" rel="noopener">front-facing</a> polygon primitive or of
a non-polygon primitive and is zero if the fragment is considered to be
part of a back-facing polygon primitive.

Valid Usage

- <a href="#VUID-FrontFacing-FrontFacing-04229"
  id="VUID-FrontFacing-FrontFacing-04229"></a>
  VUID-FrontFacing-FrontFacing-04229  
  The `FrontFacing` decoration **must** be used only within the
  `Fragment` `Execution` `Model`

- <a href="#VUID-FrontFacing-FrontFacing-04230"
  id="VUID-FrontFacing-FrontFacing-04230"></a>
  VUID-FrontFacing-FrontFacing-04230  
  The variable decorated with `FrontFacing` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-FrontFacing-FrontFacing-04231"
  id="VUID-FrontFacing-FrontFacing-04231"></a>
  VUID-FrontFacing-FrontFacing-04231  
  The variable decorated with `FrontFacing` **must** be declared as a
  boolean value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#FrontFacing"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
