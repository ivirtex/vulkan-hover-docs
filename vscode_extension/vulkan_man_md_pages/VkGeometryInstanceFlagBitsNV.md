# VkGeometryInstanceFlagBitsKHR(3) Manual Page

## Name

VkGeometryInstanceFlagBitsKHR - Instance flag bits



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of `flags` in the instance modifying the behavior of
that instance are:

``` c
// Provided by VK_KHR_acceleration_structure
typedef enum VkGeometryInstanceFlagBitsKHR {
    VK_GEOMETRY_INSTANCE_TRIANGLE_FACING_CULL_DISABLE_BIT_KHR = 0x00000001,
    VK_GEOMETRY_INSTANCE_TRIANGLE_FLIP_FACING_BIT_KHR = 0x00000002,
    VK_GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_KHR = 0x00000004,
    VK_GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_KHR = 0x00000008,
  // Provided by VK_EXT_opacity_micromap
    VK_GEOMETRY_INSTANCE_FORCE_OPACITY_MICROMAP_2_STATE_EXT = 0x00000010,
  // Provided by VK_EXT_opacity_micromap
    VK_GEOMETRY_INSTANCE_DISABLE_OPACITY_MICROMAPS_EXT = 0x00000020,
    VK_GEOMETRY_INSTANCE_TRIANGLE_FRONT_COUNTERCLOCKWISE_BIT_KHR = VK_GEOMETRY_INSTANCE_TRIANGLE_FLIP_FACING_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_GEOMETRY_INSTANCE_TRIANGLE_CULL_DISABLE_BIT_NV = VK_GEOMETRY_INSTANCE_TRIANGLE_FACING_CULL_DISABLE_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_GEOMETRY_INSTANCE_TRIANGLE_FRONT_COUNTERCLOCKWISE_BIT_NV = VK_GEOMETRY_INSTANCE_TRIANGLE_FRONT_COUNTERCLOCKWISE_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_NV = VK_GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_NV = VK_GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_KHR,
} VkGeometryInstanceFlagBitsKHR;
```

or the equivalent

``` c
// Provided by VK_NV_ray_tracing
typedef VkGeometryInstanceFlagBitsKHR VkGeometryInstanceFlagBitsNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_GEOMETRY_INSTANCE_TRIANGLE_FACING_CULL_DISABLE_BIT_KHR` disables
  face culling for this instance.

- `VK_GEOMETRY_INSTANCE_TRIANGLE_FLIP_FACING_BIT_KHR` indicates that the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-traversal-culling-face"
  target="_blank" rel="noopener">facing determination</a> for geometry
  in this instance is inverted. Because the facing is determined in
  object space, an instance transform does not change the winding, but a
  geometry transform does.

- `VK_GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_KHR` causes this instance to
  act as though `VK_GEOMETRY_OPAQUE_BIT_KHR` were specified on all
  geometries referenced by this instance. This behavior **can** be
  overridden by the SPIR-V `NoOpaqueKHR` ray flag.

- `VK_GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_KHR` causes this instance to
  act as though `VK_GEOMETRY_OPAQUE_BIT_KHR` were not specified on all
  geometries referenced by this instance. This behavior **can** be
  overridden by the SPIR-V `OpaqueKHR` ray flag.

`VK_GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_KHR` and
`VK_GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_KHR` **must** not be used in the
same flag.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkGeometryInstanceFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryInstanceFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGeometryInstanceFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
