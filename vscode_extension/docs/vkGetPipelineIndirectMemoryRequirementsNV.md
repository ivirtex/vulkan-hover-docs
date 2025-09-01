# vkGetPipelineIndirectMemoryRequirementsNV(3) Manual Page

## Name

vkGetPipelineIndirectMemoryRequirementsNV - Get the memory requirements for the compute indirect pipeline



## [](#_c_specification)C Specification

To determine the memory requirements for a compute pipeline’s metadata, call:

```c++
// Provided by VK_NV_device_generated_commands_compute
void vkGetPipelineIndirectMemoryRequirementsNV(
    VkDevice                                    device,
    const VkComputePipelineCreateInfo*          pCreateInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the buffer.
- `pCreateInfo` is a [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html) structure specifying the creation parameters of the compute pipeline whose memory requirements are being queried.
- `pMemoryRequirements` is a pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure in which the requested pipeline’s memory requirements are returned.

## [](#_description)Description

If `pCreateInfo->pNext` chain includes a pointer to a [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineIndirectBufferInfoNV.html) structure, then the contents of that structure are ignored.

Valid Usage

- [](#VUID-vkGetPipelineIndirectMemoryRequirementsNV-deviceGeneratedComputePipelines-09082)VUID-vkGetPipelineIndirectMemoryRequirementsNV-deviceGeneratedComputePipelines-09082  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedComputePipelines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedComputePipelines) feature **must** be enabled
- [](#VUID-vkGetPipelineIndirectMemoryRequirementsNV-pCreateInfo-09083)VUID-vkGetPipelineIndirectMemoryRequirementsNV-pCreateInfo-09083  
  `pCreateInfo->flags` **must** include `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`

Valid Usage (Implicit)

- [](#VUID-vkGetPipelineIndirectMemoryRequirementsNV-device-parameter)VUID-vkGetPipelineIndirectMemoryRequirementsNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetPipelineIndirectMemoryRequirementsNV-pCreateInfo-parameter)VUID-vkGetPipelineIndirectMemoryRequirementsNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html) structure
- [](#VUID-vkGetPipelineIndirectMemoryRequirementsNV-pMemoryRequirements-parameter)VUID-vkGetPipelineIndirectMemoryRequirementsNV-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands\_compute](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands_compute.html), [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPipelineIndirectMemoryRequirementsNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0