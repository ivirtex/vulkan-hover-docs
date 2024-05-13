# InstanceIndex(3) Manual Page

## Name

InstanceIndex - Index of an instance



## <a href="#_description" class="anchor"></a>Description

`InstanceIndex`  
Decorating a variable in a vertex shader with the `InstanceIndex`
built-in decoration will make that variable contain the index of the
instance that is being processed by the current vertex shader
invocation. `InstanceIndex` begins at the `firstInstance` parameter to
[vkCmdDraw](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDraw.html) or [vkCmdDrawIndexed](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexed.html)
or at the `firstInstance` member of a structure consumed by
[vkCmdDrawIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirect.html) or
[vkCmdDrawIndexedIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirect.html).

Valid Usage

- <a href="#VUID-InstanceIndex-InstanceIndex-04263"
  id="VUID-InstanceIndex-InstanceIndex-04263"></a>
  VUID-InstanceIndex-InstanceIndex-04263  
  The `InstanceIndex` decoration **must** be used only within the
  `Vertex` `Execution` `Model`

- <a href="#VUID-InstanceIndex-InstanceIndex-04264"
  id="VUID-InstanceIndex-InstanceIndex-04264"></a>
  VUID-InstanceIndex-InstanceIndex-04264  
  The variable decorated with `InstanceIndex` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-InstanceIndex-InstanceIndex-04265"
  id="VUID-InstanceIndex-InstanceIndex-04265"></a>
  VUID-InstanceIndex-InstanceIndex-04265  
  The variable decorated with `InstanceIndex` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#InstanceIndex"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
