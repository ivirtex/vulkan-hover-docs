# DrawIndex(3) Manual Page

## Name

DrawIndex - Index of the current draw



## <a href="#_description" class="anchor"></a>Description

`DrawIndex`  
Decorating a variable with the `DrawIndex` built-in will make that
variable contain the integer value corresponding to the zero-based index
of the draw that invoked the current task, mesh, or vertex shader
invocation. For *indirect drawing commands*, `DrawIndex` begins at zero
and increments by one for each draw executed. The number of draws is
given by the `drawCount` parameter. For *direct drawing commands*, if
[vkCmdDrawMultiEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMultiEXT.html) or
[vkCmdDrawMultiIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMultiIndexedEXT.html) is used, this
variable contains the integer value corresponding to the zero-based
index of the draw. Otherwise `DrawIndex` is always zero. `DrawIndex` is
dynamically uniform.

When task or mesh shaders are used, only the first active stage will
have proper access to the variable. The value read by other stages is
undefined.

Valid Usage

- <a href="#VUID-DrawIndex-DrawIndex-04207"
  id="VUID-DrawIndex-DrawIndex-04207"></a>
  VUID-DrawIndex-DrawIndex-04207  
  The `DrawIndex` decoration **must** be used only within the `Vertex`,
  `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution` `Model`

- <a href="#VUID-DrawIndex-DrawIndex-04208"
  id="VUID-DrawIndex-DrawIndex-04208"></a>
  VUID-DrawIndex-DrawIndex-04208  
  The variable decorated with `DrawIndex` **must** be declared using the
  `Input` `Storage` `Class`

- <a href="#VUID-DrawIndex-DrawIndex-04209"
  id="VUID-DrawIndex-DrawIndex-04209"></a>
  VUID-DrawIndex-DrawIndex-04209  
  The variable decorated with `DrawIndex` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#DrawIndex"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
