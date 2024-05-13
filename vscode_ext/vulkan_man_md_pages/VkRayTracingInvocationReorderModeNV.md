# VkRayTracingInvocationReorderModeNV(3) Manual Page

## Name

VkRayTracingInvocationReorderModeNV - Enum providing a hint on how the
application may: reorder



## <a href="#_c_specification" class="anchor"></a>C Specification

Values which **may** be returned in the
`rayTracingInvocationReorderReorderingHint` field of
`VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV` are:

``` c
// Provided by VK_NV_ray_tracing_invocation_reorder
typedef enum VkRayTracingInvocationReorderModeNV {
    VK_RAY_TRACING_INVOCATION_REORDER_MODE_NONE_NV = 0,
    VK_RAY_TRACING_INVOCATION_REORDER_MODE_REORDER_NV = 1,
} VkRayTracingInvocationReorderModeNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_RAY_TRACING_INVOCATION_REORDER_MODE_NONE_NV` indicates that the
  implementation is likely to not reorder at reorder calls.

- `VK_RAY_TRACING_INVOCATION_REORDER_MODE_REORDER_NV` indicates that the
  implementation is likely to reorder at reorder calls.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing_invocation_reorder](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing_invocation_reorder.html),
[VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRayTracingInvocationReorderModeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
