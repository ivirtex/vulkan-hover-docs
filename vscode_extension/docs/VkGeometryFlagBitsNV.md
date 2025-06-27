# VkGeometryFlagBitsKHR(3) Manual Page

## Name

VkGeometryFlagBitsKHR - Bitmask specifying additional parameters for a
geometry



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits specifying additional parameters for geometries in acceleration
structure builds, are:

``` c
// Provided by VK_KHR_acceleration_structure
typedef enum VkGeometryFlagBitsKHR {
    VK_GEOMETRY_OPAQUE_BIT_KHR = 0x00000001,
    VK_GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_KHR = 0x00000002,
  // Provided by VK_NV_ray_tracing
    VK_GEOMETRY_OPAQUE_BIT_NV = VK_GEOMETRY_OPAQUE_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_NV = VK_GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_KHR,
} VkGeometryFlagBitsKHR;
```

or the equivalent

``` c
// Provided by VK_NV_ray_tracing
typedef VkGeometryFlagBitsKHR VkGeometryFlagBitsNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_GEOMETRY_OPAQUE_BIT_KHR` indicates that this geometry does not
  invoke the any-hit shaders even if present in a hit group.

- `VK_GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_KHR` indicates that
  the implementation **must** only call the any-hit shader a single time
  for each primitive in this geometry. If this bit is absent an
  implementation **may** invoke the any-hit shader more than once for
  this geometry.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkGeometryFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGeometryFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
