# VkComputePipelineIndirectBufferInfoNV(3) Manual Page

## Name

VkComputePipelineIndirectBufferInfoNV - Structure describing the device address where pipeline's metadata will be saved



## [](#_c_specification)C Specification

The `VkComputePipelineIndirectBufferInfoNV` structure is defined as:

```c++
// Provided by VK_NV_device_generated_commands_compute
typedef struct VkComputePipelineIndirectBufferInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceAddress    deviceAddress;
    VkDeviceSize       size;
    VkDeviceAddress    pipelineDeviceAddressCaptureReplay;
} VkComputePipelineIndirectBufferInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `deviceAddress` is the address where the pipeline’s metadata will be stored.
- `size` is the size of pipeline’s metadata that was queried using [vkGetPipelineIndirectMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineIndirectMemoryRequirementsNV.html).
- `pipelineDeviceAddressCaptureReplay` is the device address where pipeline’s metadata was originally saved and can now be used to re-populate `deviceAddress` for replay.

## [](#_description)Description

If `pipelineDeviceAddressCaptureReplay` is zero, no specific address is requested. If `pipelineDeviceAddressCaptureReplay` is not zero, then it **must** be an address retrieved from an identically created pipeline on the same implementation. The pipeline metadata **must** also be placed on an identically created buffer and at the same offset using the [vkCmdUpdatePipelineIndirectBufferNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdUpdatePipelineIndirectBufferNV.html) command.

Valid Usage

- [](#VUID-VkComputePipelineIndirectBufferInfoNV-deviceGeneratedComputePipelines-09009)VUID-VkComputePipelineIndirectBufferInfoNV-deviceGeneratedComputePipelines-09009  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedComputePipelines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedComputePipelines) feature **must** be enabled
- [](#VUID-VkComputePipelineIndirectBufferInfoNV-flags-09010)VUID-VkComputePipelineIndirectBufferInfoNV-flags-09010  
  The pipeline creation flags in [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html)::`flags` **must** include `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`
- [](#VUID-VkComputePipelineIndirectBufferInfoNV-deviceAddress-09011)VUID-VkComputePipelineIndirectBufferInfoNV-deviceAddress-09011  
  `deviceAddress` **must** be aligned to the [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html)::`memoryRequirements.alignment`, as returned by [vkGetPipelineIndirectMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineIndirectMemoryRequirementsNV.html)
- [](#VUID-VkComputePipelineIndirectBufferInfoNV-deviceAddress-09012)VUID-VkComputePipelineIndirectBufferInfoNV-deviceAddress-09012  
  `deviceAddress` **must** have been allocated from a buffer that was created with usage `VK_BUFFER_USAGE_TRANSFER_DST_BIT` and `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT`
- [](#VUID-VkComputePipelineIndirectBufferInfoNV-size-09013)VUID-VkComputePipelineIndirectBufferInfoNV-size-09013  
  `size` **must** be greater than or equal to the [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html)::`memoryRequirements.size`, as returned by [vkGetPipelineIndirectMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineIndirectMemoryRequirementsNV.html)
- [](#VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09014)VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09014  
  If `pipelineDeviceAddressCaptureReplay` is non-zero then the [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedComputeCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedComputePipelines) feature **must** be enabled
- [](#VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09015)VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09015  
  If `pipelineDeviceAddressCaptureReplay` is non-zero then that address **must** have been allocated with flag `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` set
- [](#VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09016)VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09016  
  If `pipelineDeviceAddressCaptureReplay` is non-zero, the `pipeline` **must** have been recreated for replay
- [](#VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09017)VUID-VkComputePipelineIndirectBufferInfoNV-pipelineDeviceAddressCaptureReplay-09017  
  `pipelineDeviceAddressCaptureReplay` **must** satisfy the `alignment` and `size` requirements similar to `deviceAddress`

Valid Usage (Implicit)

- [](#VUID-VkComputePipelineIndirectBufferInfoNV-sType-sType)VUID-VkComputePipelineIndirectBufferInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_INDIRECT_BUFFER_INFO_NV`

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands\_compute](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands_compute.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkComputePipelineIndirectBufferInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0