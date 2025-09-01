# vkDestroyDataGraphPipelineSessionARM(3) Manual Page

## Name

vkDestroyDataGraphPipelineSessionARM - Destroy a data graph pipeline session object



## [](#_c_specification)C Specification

To destroy a data graph pipeline session, call:

```c++
// Provided by VK_ARM_data_graph
void vkDestroyDataGraphPipelineSessionARM(
    VkDevice                                    device,
    VkDataGraphPipelineSessionARM               session,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the data graph pipeline session.
- `session` is the handle of the data graph pipeline session to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyDataGraphPipelineSessionARM-session-09793)VUID-vkDestroyDataGraphPipelineSessionARM-session-09793  
  All submitted commands that refer to `session` **must** have completed execution
- [](#VUID-vkDestroyDataGraphPipelineSessionARM-session-09794)VUID-vkDestroyDataGraphPipelineSessionARM-session-09794  
  If `VkAllocationCallbacks` were provided when `session` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyDataGraphPipelineSessionARM-session-09795)VUID-vkDestroyDataGraphPipelineSessionARM-session-09795  
  If no `VkAllocationCallbacks` were provided when `session` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyDataGraphPipelineSessionARM-device-parameter)VUID-vkDestroyDataGraphPipelineSessionARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyDataGraphPipelineSessionARM-session-parameter)VUID-vkDestroyDataGraphPipelineSessionARM-session-parameter  
  `session` **must** be a valid [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html) handle
- [](#VUID-vkDestroyDataGraphPipelineSessionARM-pAllocator-parameter)VUID-vkDestroyDataGraphPipelineSessionARM-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyDataGraphPipelineSessionARM-session-parent)VUID-vkDestroyDataGraphPipelineSessionARM-session-parent  
  `session` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `session` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyDataGraphPipelineSessionARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0