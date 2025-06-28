# VkRayTracingInvocationReorderModeNV(3) Manual Page

## Name

VkRayTracingInvocationReorderModeNV - Enum providing a hint on how the application may: reorder



## [](#_c_specification)C Specification

Values which **may** be returned in the `rayTracingInvocationReorderReorderingHint` field of `VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV` are:

```c++
// Provided by VK_NV_ray_tracing_invocation_reorder
typedef enum VkRayTracingInvocationReorderModeNV {
    VK_RAY_TRACING_INVOCATION_REORDER_MODE_NONE_NV = 0,
    VK_RAY_TRACING_INVOCATION_REORDER_MODE_REORDER_NV = 1,
} VkRayTracingInvocationReorderModeNV;
```

## [](#_description)Description

- `VK_RAY_TRACING_INVOCATION_REORDER_MODE_NONE_NV` specifies that the implementation is likely to not reorder at reorder calls.
- `VK_RAY_TRACING_INVOCATION_REORDER_MODE_REORDER_NV` specifies that the implementation is likely to reorder at reorder calls.

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing\_invocation\_reorder](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_invocation_reorder.html), [VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRayTracingInvocationReorderModeNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0