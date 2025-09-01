# vkDestroyExternalComputeQueueNV(3) Manual Page

## Name

vkDestroyExternalComputeQueueNV - Destroys an external queue.



## [](#_c_specification)C Specification

To destroy a previously created external compute queue call:

```c++
// Provided by VK_NV_external_compute_queue
void vkDestroyExternalComputeQueueNV(
    VkDevice                                    device,
    VkExternalComputeQueueNV                    externalQueue,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the external queue.
- `externalQueue` is the [VkExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueNV.html) to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkDestroyExternalComputeQueueNV-device-parameter)VUID-vkDestroyExternalComputeQueueNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyExternalComputeQueueNV-externalQueue-parameter)VUID-vkDestroyExternalComputeQueueNV-externalQueue-parameter  
  `externalQueue` **must** be a valid [VkExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueNV.html) handle
- [](#VUID-vkDestroyExternalComputeQueueNV-pAllocator-parameter)VUID-vkDestroyExternalComputeQueueNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyExternalComputeQueueNV-externalQueue-parent)VUID-vkDestroyExternalComputeQueueNV-externalQueue-parent  
  `externalQueue` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_NV\_external\_compute\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_compute_queue.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyExternalComputeQueueNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0