# vkDestroyPipelineCache(3) Manual Page

## Name

vkDestroyPipelineCache - Destroy a pipeline cache object



## [](#_c_specification)C Specification

To destroy a pipeline cache, call:

```c++
// Provided by VK_VERSION_1_0
void vkDestroyPipelineCache(
    VkDevice                                    device,
    VkPipelineCache                             pipelineCache,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the pipeline cache object.
- `pipelineCache` is the handle of the pipeline cache to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyPipelineCache-pipelineCache-00771)VUID-vkDestroyPipelineCache-pipelineCache-00771  
  If `VkAllocationCallbacks` were provided when `pipelineCache` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyPipelineCache-pipelineCache-00772)VUID-vkDestroyPipelineCache-pipelineCache-00772  
  If no `VkAllocationCallbacks` were provided when `pipelineCache` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyPipelineCache-device-parameter)VUID-vkDestroyPipelineCache-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyPipelineCache-pipelineCache-parameter)VUID-vkDestroyPipelineCache-pipelineCache-parameter  
  If `pipelineCache` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pipelineCache` **must** be a valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) handle
- [](#VUID-vkDestroyPipelineCache-pAllocator-parameter)VUID-vkDestroyPipelineCache-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyPipelineCache-pipelineCache-parent)VUID-vkDestroyPipelineCache-pipelineCache-parent  
  If `pipelineCache` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `pipelineCache` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyPipelineCache)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0