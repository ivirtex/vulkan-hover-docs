# CullPrimitiveEXT(3) Manual Page

## Name

CullPrimitiveEXT - Application-specified culling state per primitive



## <a href="#_description" class="anchor"></a>Description

`CullPrimitiveEXT`  
Decorating a variable with the `CullPrimitiveEXT` built-in decoration
will make that variable contain the culling state of output primitives.
If the per-primitive boolean value is `true`, the primitive will be
culled, if it is `false` it will not be culled.

Valid Usage

- <a href="#VUID-CullPrimitiveEXT-CullPrimitiveEXT-07034"
  id="VUID-CullPrimitiveEXT-CullPrimitiveEXT-07034"></a>
  VUID-CullPrimitiveEXT-CullPrimitiveEXT-07034  
  The `CullPrimitiveEXT` decoration **must** be used only within the
  `MeshEXT` `Execution` `Model`

- <a href="#VUID-CullPrimitiveEXT-CullPrimitiveEXT-07035"
  id="VUID-CullPrimitiveEXT-CullPrimitiveEXT-07035"></a>
  VUID-CullPrimitiveEXT-CullPrimitiveEXT-07035  
  The variable decorated with `CullPrimitiveEXT` **must** be declared
  using the `Output` `Storage` `Class`

- <a href="#VUID-CullPrimitiveEXT-CullPrimitiveEXT-07036"
  id="VUID-CullPrimitiveEXT-CullPrimitiveEXT-07036"></a>
  VUID-CullPrimitiveEXT-CullPrimitiveEXT-07036  
  The variable decorated with `CullPrimitiveEXT` **must** be declared as
  an array of boolean values

- <a href="#VUID-CullPrimitiveEXT-CullPrimitiveEXT-07037"
  id="VUID-CullPrimitiveEXT-CullPrimitiveEXT-07037"></a>
  VUID-CullPrimitiveEXT-CullPrimitiveEXT-07037  
  The size of the array decorated with `CullPrimitiveEXT` **must** match
  the value specified by `OutputPrimitivesEXT`

- <a href="#VUID-CullPrimitiveEXT-CullPrimitiveEXT-07038"
  id="VUID-CullPrimitiveEXT-CullPrimitiveEXT-07038"></a>
  VUID-CullPrimitiveEXT-CullPrimitiveEXT-07038  
  The variable decorated with `CullPrimitiveEXT` within the `MeshEXT`
  `Execution` `Model` **must** also be decorated with the
  `PerPrimitiveEXT` decoration

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#CullPrimitiveEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
