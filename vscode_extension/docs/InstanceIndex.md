# InstanceIndex(3) Manual Page

## Name

InstanceIndex - Index of an instance



## [](#_description)Description

`InstanceIndex`

Decorating a variable in a vertex shader with the `InstanceIndex` built-in decoration will make that variable contain the index of the instance that is being processed by the current vertex shader invocation. `InstanceIndex` begins at the `firstInstance` parameter to [vkCmdDraw](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDraw.html) or [vkCmdDrawIndexed](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexed.html) or at the `firstInstance` member of a structure consumed by [vkCmdDrawIndirect](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndirect.html) or [vkCmdDrawIndexedIndirect](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexedIndirect.html).

Valid Usage

- [](#VUID-InstanceIndex-InstanceIndex-04263)VUID-InstanceIndex-InstanceIndex-04263  
  The `InstanceIndex` decoration **must** be used only within the `Vertex` `Execution` `Model`
- [](#VUID-InstanceIndex-InstanceIndex-04264)VUID-InstanceIndex-InstanceIndex-04264  
  The variable decorated with `InstanceIndex` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-InstanceIndex-InstanceIndex-04265)VUID-InstanceIndex-InstanceIndex-04265  
  The variable decorated with `InstanceIndex` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#InstanceIndex)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0