# VK\_SHADER\_UNUSED\_KHR(3) Manual Page

## Name

VK\_SHADER\_UNUSED\_KHR - Sentinel for an unused shader index



## [](#_c_specification)C Specification

`VK_SHADER_UNUSED_KHR` is a special shader index used to indicate that a ray generation, miss, or callable shader member is not used.

```c++
#define VK_SHADER_UNUSED_KHR              (~0U)
```

or the equivalent

```c++
#define VK_SHADER_UNUSED_NV               VK_SHADER_UNUSED_KHR
```

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_SHADER_UNUSED_KHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0