# vkCreateExternalComputeQueueNV(3) Manual Page

## Name

vkCreateExternalComputeQueueNV - Create an external compute queue for use by a compatible external API.



## [](#_c_specification)C Specification

To create an external compute queue for use by compatible external APIs call:

```c++
// Provided by VK_NV_external_compute_queue
VkResult vkCreateExternalComputeQueueNV(
    VkDevice                                    device,
    const VkExternalComputeQueueCreateInfoNV*   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkExternalComputeQueueNV*                   pExternalQueue);
```

## [](#_parameters)Parameters

- `device` is the VkDevice that the external queue will be a part of.
- `pCreateInfo` is a pointer to a [VkExternalComputeQueueCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueCreateInfoNV.html) structure specifying configuration info for creating the external queue.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pExternalQueue` is a pointer to a [VkExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueNV.html) object that will be filled with the handle for the created external queue.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateExternalComputeQueueNV-device-parameter)VUID-vkCreateExternalComputeQueueNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateExternalComputeQueueNV-pCreateInfo-parameter)VUID-vkCreateExternalComputeQueueNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkExternalComputeQueueCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueCreateInfoNV.html) structure
- [](#VUID-vkCreateExternalComputeQueueNV-pAllocator-parameter)VUID-vkCreateExternalComputeQueueNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateExternalComputeQueueNV-pExternalQueue-parameter)VUID-vkCreateExternalComputeQueueNV-pExternalQueue-parameter  
  `pExternalQueue` **must** be a valid pointer to a [VkExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueNV.html) handle
- [](#VUID-vkCreateExternalComputeQueueNV-device-queuecount)VUID-vkCreateExternalComputeQueueNV-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_TOO_MANY_OBJECTS`

## [](#_see_also)See Also

[VK\_NV\_external\_compute\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_compute_queue.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExternalComputeQueueCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueCreateInfoNV.html), [VkExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateExternalComputeQueueNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0