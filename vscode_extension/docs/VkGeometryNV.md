# VkGeometryNV(3) Manual Page

## Name

VkGeometryNV - Structure specifying a geometry in a bottom-level acceleration structure



## [](#_c_specification)C Specification

The `VkGeometryNV` structure describes geometry in a bottom-level acceleration structure and is defined as:

```c++
// Provided by VK_NV_ray_tracing
typedef struct VkGeometryNV {
    VkStructureType       sType;
    const void*           pNext;
    VkGeometryTypeKHR     geometryType;
    VkGeometryDataNV      geometry;
    VkGeometryFlagsKHR    flags;
} VkGeometryNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `geometryType` specifies the [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTypeKHR.html) which this geometry refers to.
- `geometry` contains the geometry data as described in [VkGeometryDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryDataNV.html).
- `flags` has [VkGeometryFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryFlagBitsKHR.html) describing options for this geometry.

## [](#_description)Description

Valid Usage

- [](#VUID-VkGeometryNV-geometryType-03503)VUID-VkGeometryNV-geometryType-03503  
  `geometryType` **must** be `VK_GEOMETRY_TYPE_TRIANGLES_NV` or `VK_GEOMETRY_TYPE_AABBS_NV`

Valid Usage (Implicit)

- [](#VUID-VkGeometryNV-sType-sType)VUID-VkGeometryNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GEOMETRY_NV`
- [](#VUID-VkGeometryNV-pNext-pNext)VUID-VkGeometryNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkGeometryNV-geometryType-parameter)VUID-VkGeometryNV-geometryType-parameter  
  `geometryType` **must** be a valid [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTypeKHR.html) value
- [](#VUID-VkGeometryNV-geometry-parameter)VUID-VkGeometryNV-geometry-parameter  
  `geometry` **must** be a valid [VkGeometryDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryDataNV.html) structure
- [](#VUID-VkGeometryNV-flags-parameter)VUID-VkGeometryNV-flags-parameter  
  `flags` **must** be a valid combination of [VkGeometryFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryFlagBitsKHR.html) values

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html), [VkGeometryDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryDataNV.html), [VkGeometryFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryFlagsKHR.html), [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTypeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGeometryNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0