# vkDestroyPipelineLayout(3) Manual Page

## Name

vkDestroyPipelineLayout - Destroy a pipeline layout object



## [](#_c_specification)C Specification

To destroy a pipeline layout, call:

```c++
// Provided by VK_VERSION_1_0
void vkDestroyPipelineLayout(
    VkDevice                                    device,
    VkPipelineLayout                            pipelineLayout,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the pipeline layout.
- `pipelineLayout` is the pipeline layout to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyPipelineLayout-pipelineLayout-00299)VUID-vkDestroyPipelineLayout-pipelineLayout-00299  
  If `VkAllocationCallbacks` were provided when `pipelineLayout` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyPipelineLayout-pipelineLayout-00300)VUID-vkDestroyPipelineLayout-pipelineLayout-00300  
  If no `VkAllocationCallbacks` were provided when `pipelineLayout` was created, `pAllocator` **must** be `NULL`
- [](#VUID-vkDestroyPipelineLayout-pipelineLayout-02004)VUID-vkDestroyPipelineLayout-pipelineLayout-02004  
  `pipelineLayout` **must** not have been passed to any `vkCmd*` command for any command buffers that are still in the [recording state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle) when `vkDestroyPipelineLayout` is called

Valid Usage (Implicit)

- [](#VUID-vkDestroyPipelineLayout-device-parameter)VUID-vkDestroyPipelineLayout-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyPipelineLayout-pipelineLayout-parameter)VUID-vkDestroyPipelineLayout-pipelineLayout-parameter  
  If `pipelineLayout` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pipelineLayout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-vkDestroyPipelineLayout-pAllocator-parameter)VUID-vkDestroyPipelineLayout-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyPipelineLayout-pipelineLayout-parent)VUID-vkDestroyPipelineLayout-pipelineLayout-parent  
  If `pipelineLayout` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `pipelineLayout` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyPipelineLayout).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0