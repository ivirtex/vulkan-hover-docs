# VkGeometryTypeKHR(3) Manual Page

## Name

VkGeometryTypeKHR - Enum specifying which type of geometry is provided



## [](#_c_specification)C Specification

Geometry types are specified by [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTypeKHR.html), which takes values:

```c++
// Provided by VK_KHR_acceleration_structure
typedef enum VkGeometryTypeKHR {
    VK_GEOMETRY_TYPE_TRIANGLES_KHR = 0,
    VK_GEOMETRY_TYPE_AABBS_KHR = 1,
    VK_GEOMETRY_TYPE_INSTANCES_KHR = 2,
  // Provided by VK_NV_ray_tracing_linear_swept_spheres
    VK_GEOMETRY_TYPE_SPHERES_NV = 1000429004,
  // Provided by VK_NV_ray_tracing_linear_swept_spheres
    VK_GEOMETRY_TYPE_LINEAR_SWEPT_SPHERES_NV = 1000429005,
#ifdef VK_ENABLE_BETA_EXTENSIONS
  // Provided by VK_AMDX_dense_geometry_format
    VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX = 1000478000,
#endif
  // Provided by VK_NV_ray_tracing
    VK_GEOMETRY_TYPE_TRIANGLES_NV = VK_GEOMETRY_TYPE_TRIANGLES_KHR,
  // Provided by VK_NV_ray_tracing
    VK_GEOMETRY_TYPE_AABBS_NV = VK_GEOMETRY_TYPE_AABBS_KHR,
} VkGeometryTypeKHR;
```

or the equivalent

```c++
// Provided by VK_NV_ray_tracing
typedef VkGeometryTypeKHR VkGeometryTypeNV;
```

## [](#_description)Description

- `VK_GEOMETRY_TYPE_TRIANGLES_KHR` specifies a geometry type consisting of [triangles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-tracing-triangle-primitive).
- `VK_GEOMETRY_TYPE_AABBS_KHR` specifies a geometry type consisting of [axis-aligned bounding boxes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aabb-primitive).
- `VK_GEOMETRY_TYPE_INSTANCES_KHR` specifies a geometry type consisting of acceleration structure instances.
- `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX` specifies a geometry type consisting of triangles from compressed data.
- `VK_GEOMETRY_TYPE_SPHERES_NV` specifies a geometry type consisting of [spheres](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#sphere-primitive).
- `VK_GEOMETRY_TYPE_LINEAR_SWEPT_SPHERES_NV` specifies a geometry type consisting of [linear swept spheres](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#linear-swept-sphere-primitive).

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureGeometryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryKHR.html), [VkGeometryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGeometryTypeKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0