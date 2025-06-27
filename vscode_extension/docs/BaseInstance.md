# BaseInstance(3) Manual Page

## Name

BaseInstance - First instance being rendered



## <a href="#_description" class="anchor"></a>Description

`BaseInstance`  
Decorating a variable with the `BaseInstance` built-in will make that
variable contain the integer value corresponding to the first instance
that was passed to the command that invoked the current vertex shader
invocation. `BaseInstance` is the `firstInstance` parameter to a *direct
drawing command* or the `firstInstance` member of a structure consumed
by an *indirect drawing command*.

Valid Usage

- <a href="#VUID-BaseInstance-BaseInstance-04181"
  id="VUID-BaseInstance-BaseInstance-04181"></a>
  VUID-BaseInstance-BaseInstance-04181  
  The `BaseInstance` decoration **must** be used only within the
  `Vertex` `Execution` `Model`

- <a href="#VUID-BaseInstance-BaseInstance-04182"
  id="VUID-BaseInstance-BaseInstance-04182"></a>
  VUID-BaseInstance-BaseInstance-04182  
  The variable decorated with `BaseInstance` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-BaseInstance-BaseInstance-04183"
  id="VUID-BaseInstance-BaseInstance-04183"></a>
  VUID-BaseInstance-BaseInstance-04183  
  The variable decorated with `BaseInstance` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#BaseInstance"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
