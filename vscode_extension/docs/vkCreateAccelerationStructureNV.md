# vkCreateAccelerationStructureNV(3) Manual Page

## Name

vkCreateAccelerationStructureNV - Create a new acceleration structure object



## [](#_c_specification)C Specification

To create acceleration structures, call:

```c++
// Provided by VK_NV_ray_tracing
VkResult vkCreateAccelerationStructureNV(
    VkDevice                                    device,
    const VkAccelerationStructureCreateInfoNV*  pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkAccelerationStructureNV*                  pAccelerationStructure);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the buffer object.
- `pCreateInfo` is a pointer to a [VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoNV.html) structure containing parameters affecting creation of the acceleration structure.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pAccelerationStructure` is a pointer to a [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handle in which the resulting acceleration structure object is returned.

## [](#_description)Description

Similarly to other objects in Vulkan, the acceleration structure creation merely creates an object with a specific “shape” as specified by the information in [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html) and `compactedSize` in `pCreateInfo`.

Once memory has been bound to the acceleration structure using [vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindAccelerationStructureMemoryNV.html), that memory is populated by calls to [vkCmdBuildAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildAccelerationStructureNV.html) and [vkCmdCopyAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureNV.html).

Acceleration structure creation uses the count and type information from the geometries, but does not use the data references in the structures.

Valid Usage (Implicit)

- [](#VUID-vkCreateAccelerationStructureNV-device-parameter)VUID-vkCreateAccelerationStructureNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateAccelerationStructureNV-pCreateInfo-parameter)VUID-vkCreateAccelerationStructureNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoNV.html) structure
- [](#VUID-vkCreateAccelerationStructureNV-pAllocator-parameter)VUID-vkCreateAccelerationStructureNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateAccelerationStructureNV-pAccelerationStructure-parameter)VUID-vkCreateAccelerationStructureNV-pAccelerationStructure-parameter  
  `pAccelerationStructure` **must** be a valid pointer to a [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handle
- [](#VUID-vkCreateAccelerationStructureNV-device-queuecount)VUID-vkCreateAccelerationStructureNV-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoNV.html), [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateAccelerationStructureNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0