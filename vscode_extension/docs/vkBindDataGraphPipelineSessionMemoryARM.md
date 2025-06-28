# vkBindDataGraphPipelineSessionMemoryARM(3) Manual Page

## Name

vkBindDataGraphPipelineSessionMemoryARM - Bind device memory to a data graph pipeline session object



## [](#_c_specification)C Specification

To attach memory to a data graph pipeline session object, call:

```c++
// Provided by VK_ARM_data_graph
VkResult vkBindDataGraphPipelineSessionMemoryARM(
    VkDevice                                    device,
    uint32_t                                    bindInfoCount,
    const VkBindDataGraphPipelineSessionMemoryInfoARM* pBindInfos);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the data graph pipeline session and memory.
- `bindInfoCount` is the length of the `pBindInfos` array.
- `pBindInfos` is a pointer to an array of [VkBindDataGraphPipelineSessionMemoryInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDataGraphPipelineSessionMemoryInfoARM.html) structures describing graph pipeline sessions and memory to bind.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkBindDataGraphPipelineSessionMemoryARM-device-parameter)VUID-vkBindDataGraphPipelineSessionMemoryARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkBindDataGraphPipelineSessionMemoryARM-pBindInfos-parameter)VUID-vkBindDataGraphPipelineSessionMemoryARM-pBindInfos-parameter  
  `pBindInfos` **must** be a valid pointer to an array of `bindInfoCount` valid [VkBindDataGraphPipelineSessionMemoryInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDataGraphPipelineSessionMemoryInfoARM.html) structures
- [](#VUID-vkBindDataGraphPipelineSessionMemoryARM-bindInfoCount-arraylength)VUID-vkBindDataGraphPipelineSessionMemoryARM-bindInfoCount-arraylength  
  `bindInfoCount` **must** be greater than `0`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkBindDataGraphPipelineSessionMemoryInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDataGraphPipelineSessionMemoryInfoARM.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkBindDataGraphPipelineSessionMemoryARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0