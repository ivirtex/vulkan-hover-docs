# VkAccelerationStructureTypeKHR(3) Manual Page

## Name

VkAccelerationStructureTypeKHR - Type of acceleration structure



## [](#_c_specification)C Specification

Values which **can** be set in [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type` or [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html)::`type` specifying the type of acceleration structure, are:

```c++
// Provided by VK_KHR_acceleration_structure
typedef enum VkAccelerationStructureTypeKHR {
    VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR = 0,
    VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR = 1,
    VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR = 2,
  // Provided by VK_NV_ray_tracing
    VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_NV = VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR,
  // Provided by VK_NV_ray_tracing
    VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_NV = VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR,
} VkAccelerationStructureTypeKHR;
```

or the equivalent

```c++
// Provided by VK_NV_ray_tracing
typedef VkAccelerationStructureTypeKHR VkAccelerationStructureTypeNV;
```

## [](#_description)Description

- `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR` is a top-level acceleration structure containing instance data referring to bottom-level acceleration structures.
- `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR` is a bottom-level acceleration structure containing the AABBs or geometry to be intersected.
- `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR` is an acceleration structure whose type is determined at build time used for special circumstances. In these cases, the acceleration structure type is not known at creation time, but **must** be specified at build time as either top or bottom.

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html), [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html), [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureTypeKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0