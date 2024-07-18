# VkGeometryTypeKHR(3) Manual Page

## Name

VkGeometryTypeKHR - Enum specifying which type of geometry is provided



## <a href="#_c_specification" class="anchor"></a>C Specification

Geometry types are specified by
[VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryTypeKHR.html), which takes values:

``` c
// Provided by VK_KHR_acceleration_structure
typedef enum VkGeometryTypeKHR {
    VK_GEOMETRY_TYPE_TRIANGLES_KHR = 0,
    VK_GEOMETRY_TYPE_AABBS_KHR = 1,
    VK_GEOMETRY_TYPE_INSTANCES_KHR = 2,
  // Provided by VK_NV_ray_tracing
    VK_GEOMETRY_TYPE_TRIANGLES_NV = VK_GEOMETRY_TYPE_TRIANGLES_KHR,
  // Provided by VK_NV_ray_tracing
    VK_GEOMETRY_TYPE_AABBS_NV = VK_GEOMETRY_TYPE_AABBS_KHR,
} VkGeometryTypeKHR;
```

or the equivalent

``` c
// Provided by VK_NV_ray_tracing
typedef VkGeometryTypeKHR VkGeometryTypeNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_GEOMETRY_TYPE_TRIANGLES_KHR` specifies a geometry type consisting
  of triangles.

- `VK_GEOMETRY_TYPE_AABBS_KHR` specifies a geometry type consisting of
  axis-aligned bounding boxes.

- `VK_GEOMETRY_TYPE_INSTANCES_KHR` specifies a geometry type consisting
  of acceleration structure instances.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureGeometryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryKHR.html),
[VkGeometryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGeometryTypeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
