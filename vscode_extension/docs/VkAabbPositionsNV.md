# VkAabbPositionsKHR(3) Manual Page

## Name

VkAabbPositionsKHR - Structure specifying two opposing corners of an axis-aligned bounding box



## [](#_c_specification)C Specification

The `VkAabbPositionsKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkAabbPositionsKHR {
    float    minX;
    float    minY;
    float    minZ;
    float    maxX;
    float    maxY;
    float    maxZ;
} VkAabbPositionsKHR;
```

or the equivalent

```c++
// Provided by VK_NV_ray_tracing
typedef VkAabbPositionsKHR VkAabbPositionsNV;
```

## [](#_members)Members

- `minX` is the x position of one opposing corner of a bounding box.
- `minY` is the y position of one opposing corner of a bounding box.
- `minZ` is the z position of one opposing corner of a bounding box.
- `maxX` is the x position of the other opposing corner of a bounding box.
- `maxY` is the y position of the other opposing corner of a bounding box.
- `maxZ` is the z position of the other opposing corner of a bounding box.

## [](#_description)Description

Valid Usage

- [](#VUID-VkAabbPositionsKHR-minX-03546)VUID-VkAabbPositionsKHR-minX-03546  
  `minX` **must** be less than or equal to `maxX`
- [](#VUID-VkAabbPositionsKHR-minY-03547)VUID-VkAabbPositionsKHR-minY-03547  
  `minY` **must** be less than or equal to `maxY`
- [](#VUID-VkAabbPositionsKHR-minZ-03548)VUID-VkAabbPositionsKHR-minZ-03548  
  `minZ` **must** be less than or equal to `maxZ`

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAabbPositionsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0