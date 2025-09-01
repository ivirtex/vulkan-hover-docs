# VkRayTracingLssPrimitiveEndCapsModeNV(3) Manual Page

## Name

VkRayTracingLssPrimitiveEndCapsModeNV - LSS endcaps mode



## [](#_c_specification)C Specification

The default behavior with endcaps in a LSS chain is that both endcaps will be enabled for all beginning and end points. To change the LSS chain’s endcaps mode use [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html)::`endCapsMode`. The possible values for `endCapsMode` are:

```c++
// Provided by VK_NV_ray_tracing_linear_swept_spheres
typedef enum VkRayTracingLssPrimitiveEndCapsModeNV {
    VK_RAY_TRACING_LSS_PRIMITIVE_END_CAPS_MODE_NONE_NV = 0,
    VK_RAY_TRACING_LSS_PRIMITIVE_END_CAPS_MODE_CHAINED_NV = 1,
} VkRayTracingLssPrimitiveEndCapsModeNV;
```

## [](#_description)Description

- `VK_RAY_TRACING_LSS_PRIMITIVE_END_CAPS_MODE_NONE_NV` disables all endcaps and the chain boundaries have no influence.
- `VK_RAY_TRACING_LSS_PRIMITIVE_END_CAPS_MODE_CHAINED_NV` specifies that when `VK_RAY_TRACING_LSS_INDEXING_MODE_SUCCESSIVE_NV` is used as indexing mode for the LSS primitive, the first primitive in each chain will have both endcaps enabled, and every following primitive in the chain only has endcaps at the trailing position enabled.

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing\_linear\_swept\_spheres](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_linear_swept_spheres.html), [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRayTracingLssPrimitiveEndCapsModeNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0