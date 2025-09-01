# InstanceCustomIndexKHR(3) Manual Page

## Name

InstanceCustomIndexKHR - Custom index associated with an intersected instance



## [](#_description)Description

`InstanceCustomIndexKHR`

A variable decorated with the `InstanceCustomIndexKHR` decoration will contain the application-defined value of the instance that intersects the current ray. This variable contains the value that was specified in [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html)::`instanceCustomIndex` for the current acceleration structure instance in the lower 24 bits and the upper 8 bits will be zero.

Valid Usage

- [](#VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04251)VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04251  
  The `InstanceCustomIndexKHR` decoration **must** be used only within the `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04252)VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04252  
  The variable decorated with `InstanceCustomIndexKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04253)VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04253  
  The variable decorated with `InstanceCustomIndexKHR` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#InstanceCustomIndexKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0