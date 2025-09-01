# PrimitiveCountNV(3) Manual Page

## Name

PrimitiveCountNV - Number of primitives output by a mesh shader



## [](#_description)Description

`PrimitiveCountNV`

Decorating a variable with the `PrimitiveCountNV` decoration will make that variable contain the primitive count. The primitive count specifies the number of primitives in the output mesh produced by the mesh shader that will be processed by subsequent pipeline stages.

Valid Usage

- [](#VUID-PrimitiveCountNV-PrimitiveCountNV-04327)VUID-PrimitiveCountNV-PrimitiveCountNV-04327  
  The `PrimitiveCountNV` decoration **must** be used only within the `MeshNV` `Execution` `Model`
- [](#VUID-PrimitiveCountNV-PrimitiveCountNV-04328)VUID-PrimitiveCountNV-PrimitiveCountNV-04328  
  The variable decorated with `PrimitiveCountNV` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-PrimitiveCountNV-PrimitiveCountNV-04329)VUID-PrimitiveCountNV-PrimitiveCountNV-04329  
  The variable decorated with `PrimitiveCountNV` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PrimitiveCountNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0