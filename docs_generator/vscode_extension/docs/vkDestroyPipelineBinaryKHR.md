# vkDestroyPipelineBinaryKHR(3) Manual Page

## Name

vkDestroyPipelineBinaryKHR - Destroy a pipeline binary



## [](#_c_specification)C Specification

To destroy a [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html), call:

```c++
// Provided by VK_KHR_pipeline_binary
void vkDestroyPipelineBinaryKHR(
    VkDevice                                    device,
    VkPipelineBinaryKHR                         pipelineBinary,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the pipeline binary object.
- `pipelineBinary` is the handle of the pipeline binary object to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyPipelineBinaryKHR-pipelineBinary-09614)VUID-vkDestroyPipelineBinaryKHR-pipelineBinary-09614  
  If `VkAllocationCallbacks` were provided when `pipelineBinary` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyPipelineBinaryKHR-pipelineBinary-09615)VUID-vkDestroyPipelineBinaryKHR-pipelineBinary-09615  
  If no [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) were provided when `pipelineBinary` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyPipelineBinaryKHR-device-parameter)VUID-vkDestroyPipelineBinaryKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyPipelineBinaryKHR-pipelineBinary-parameter)VUID-vkDestroyPipelineBinaryKHR-pipelineBinary-parameter  
  If `pipelineBinary` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pipelineBinary` **must** be a valid [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html) handle
- [](#VUID-vkDestroyPipelineBinaryKHR-pAllocator-parameter)VUID-vkDestroyPipelineBinaryKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyPipelineBinaryKHR-pipelineBinary-parent)VUID-vkDestroyPipelineBinaryKHR-pipelineBinary-parent  
  If `pipelineBinary` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `pipelineBinary` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyPipelineBinaryKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0