# vkBindAccelerationStructureMemoryNV(3) Manual Page

## Name

vkBindAccelerationStructureMemoryNV - Bind acceleration structure memory



## [](#_c_specification)C Specification

To attach memory to one or more acceleration structures at a time, call:

```c++
// Provided by VK_NV_ray_tracing
VkResult vkBindAccelerationStructureMemoryNV(
    VkDevice                                    device,
    uint32_t                                    bindInfoCount,
    const VkBindAccelerationStructureMemoryInfoNV* pBindInfos);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the acceleration structures and memory.
- `bindInfoCount` is the number of elements in `pBindInfos`.
- `pBindInfos` is a pointer to an array of [VkBindAccelerationStructureMemoryInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindAccelerationStructureMemoryInfoNV.html) structures describing acceleration structures and memory to bind.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkBindAccelerationStructureMemoryNV-device-parameter)VUID-vkBindAccelerationStructureMemoryNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkBindAccelerationStructureMemoryNV-pBindInfos-parameter)VUID-vkBindAccelerationStructureMemoryNV-pBindInfos-parameter  
  `pBindInfos` **must** be a valid pointer to an array of `bindInfoCount` valid [VkBindAccelerationStructureMemoryInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindAccelerationStructureMemoryInfoNV.html) structures
- [](#VUID-vkBindAccelerationStructureMemoryNV-bindInfoCount-arraylength)VUID-vkBindAccelerationStructureMemoryNV-bindInfoCount-arraylength  
  `bindInfoCount` **must** be greater than `0`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkBindAccelerationStructureMemoryInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindAccelerationStructureMemoryInfoNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkBindAccelerationStructureMemoryNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0