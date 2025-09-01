# vkCreateDataGraphPipelineSessionARM(3) Manual Page

## Name

vkCreateDataGraphPipelineSessionARM - Create a data graph pipeline session



## [](#_c_specification)C Specification

To create a data graph pipeline session, call

```c++
// Provided by VK_ARM_data_graph
VkResult vkCreateDataGraphPipelineSessionARM(
    VkDevice                                    device,
    const VkDataGraphPipelineSessionCreateInfoARM* pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDataGraphPipelineSessionARM*              pSession);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the data graph pipeline session.
- `pCreateInfo` is a pointer to a [VkDataGraphPipelineSessionCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionCreateInfoARM.html) structure.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pSession` is a pointer to a [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html) handle in which the resulting data graph pipeline session object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateDataGraphPipelineSessionARM-device-parameter)VUID-vkCreateDataGraphPipelineSessionARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateDataGraphPipelineSessionARM-pCreateInfo-parameter)VUID-vkCreateDataGraphPipelineSessionARM-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkDataGraphPipelineSessionCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionCreateInfoARM.html) structure
- [](#VUID-vkCreateDataGraphPipelineSessionARM-pAllocator-parameter)VUID-vkCreateDataGraphPipelineSessionARM-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateDataGraphPipelineSessionARM-pSession-parameter)VUID-vkCreateDataGraphPipelineSessionARM-pSession-parameter  
  `pSession` **must** be a valid pointer to a [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html) handle
- [](#VUID-vkCreateDataGraphPipelineSessionARM-device-queuecount)VUID-vkCreateDataGraphPipelineSessionARM-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html), [VkDataGraphPipelineSessionCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionCreateInfoARM.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateDataGraphPipelineSessionARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0