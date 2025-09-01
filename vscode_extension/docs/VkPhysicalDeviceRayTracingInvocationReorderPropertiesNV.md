# VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV - Structure describing shader module identifier properties of an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_ray_tracing_invocation_reorder
typedef struct VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV {
    VkStructureType                        sType;
    void*                                  pNext;
    VkRayTracingInvocationReorderModeNV    rayTracingInvocationReorderReorderingHint;
} VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `rayTracingInvocationReorderReorderingHint` is a hint indicating if the implementation will actually reorder at the reorder calls.

## [](#_description)Description

Note

Because the extension changes how hits are managed there is a compatibility reason to expose the extension even when an implementation does not have sorting active.

If the `VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV-sType-sType)VUID-VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_INVOCATION_REORDER_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing\_invocation\_reorder](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_invocation_reorder.html), [VkRayTracingInvocationReorderModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingInvocationReorderModeNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0