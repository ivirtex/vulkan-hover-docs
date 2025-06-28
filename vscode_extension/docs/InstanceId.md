# InstanceId(3) Manual Page

## Name

InstanceId - Id associated with an intersected instance



## [](#_description)Description

`InstanceId`

Decorating a variable in an intersection, any-hit, or closest hit shader with the `InstanceId` decoration will make that variable contain the index of the instance that intersects the current ray.

Valid Usage

- [](#VUID-InstanceId-InstanceId-04254)VUID-InstanceId-InstanceId-04254  
  The `InstanceId` decoration **must** be used only within the `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-InstanceId-InstanceId-04255)VUID-InstanceId-InstanceId-04255  
  The variable decorated with `InstanceId` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-InstanceId-InstanceId-04256)VUID-InstanceId-InstanceId-04256  
  The variable decorated with `InstanceId` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#InstanceId)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0